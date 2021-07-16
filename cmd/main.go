package main

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/scorpiotzh/mylog"
	"github.com/urfave/cli"
	"log-monitor/config"
	"log-monitor/elastic"
	"log-monitor/http_server"
	"log-monitor/timer"
	"log-monitor/utils"
	"os"
	"runtime"
	"runtime/debug"
	"sort"
	"sync"
	"time"
)

var (
	exit = make(chan struct{})
	log  = mylog.NewLogger("main", mylog.LevelDebug)
)

func main() {
	app := func() *cli.App {
		return cli.NewApp()
	}()
	app.Action = runServer
	app.HideVersion = true
	globalFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Usage: "config file abs path",
		},
	}
	app.Flags = append(app.Flags, globalFlags...)
	app.Commands = []cli.Command{}
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Before = func(ctx *cli.Context) error {
		debug.FreeOSMemory()
		minCore := runtime.NumCPU() // below go version 1.5,returns 1
		if minCore < 4 {
			minCore = 4
		}
		runtime.GOMAXPROCS(minCore)
		return nil
	}
	app.After = func(ctx *cli.Context) error {
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runServer(ctx *cli.Context) error {
	// 配置文件
	configFilePath := readConfigFilePath(ctx)
	log.Info("configFilePath:", configFilePath)
	if err := config.InitCfgFromFile(configFilePath, &config.Cfg); err != nil {
		return fmt.Errorf("InitCfgFromFile err:%s", err.Error())
	}
	// 热更新
	if err := utils.AddConfigFileWatcher(configFilePath, func(optName fsnotify.Op) {
		if err := config.InitCfgFromFile(configFilePath, &config.Cfg); err != nil {
			log.Error("InitCfgFromFile err:", err.Error())
		}
	}); err != nil {
		return fmt.Errorf("AddConfigFileWatcher err: %s", err.Error())
	}
	log.Info("config:", utils.Json(&config.Cfg))
	// 安全退出
	ctxServer, ctxCancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	utils.ListenSysInterrupt(func(sig os.Signal) {
		log.Warnf("signal [%s] to exit server...., time: %s", sig.String(), time.Now().String())
		ctxCancel()
		wg.Wait()
		exit <- struct{}{}
	})
	// 业务
	ela, err := elastic.Initialize(ctxServer, config.Cfg.ElasticServer.Url, config.Cfg.ElasticServer.Username, config.Cfg.ElasticServer.Password)
	if err != nil {
		return fmt.Errorf("elastic.Initialize err:%s", err.Error())
	}

	logTimer := timer.Initialize(ela)
	logTimer.RunDeleteLogByLogDate(config.Cfg.TimerServer.DeleteIndexList)
	logTimer.RunApiCheck(ctxServer, &wg)

	logHttp := http_server.Initialize(ela, config.Cfg.HttpServer.InAddress)
	logHttp.Run()

	<-exit
	log.Warn("success exit server. bye bye!")
	return nil
}

func readConfigFilePath(ctx *cli.Context) string {
	if configFileAbsPath := ctx.GlobalString("config"); configFileAbsPath != "" {
		return configFileAbsPath
	} else {
		defaultCfgFilePath := "conf/config.yaml"
		return defaultCfgFilePath
	}
}
