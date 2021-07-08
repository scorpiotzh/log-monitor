package utils

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// exit the program gracefully
func ListenSysInterrupt(handler func(sig os.Signal)) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)    // signal int, kill session windows
	signal.Notify(signalChan, syscall.SIGHUP)  // ctrl+C
	signal.Notify(signalChan, os.Kill)         // kill -9
	signal.Notify(signalChan, syscall.SIGTERM) // kill -15
	go func() {
		for {
			select {
			case s := <-signalChan:
				handler(s)
			}
		}
	}()
}

// update config constant without restart the whole server
func AddConfigFileWatcher(configFilePath string, handler func(optName fsnotify.Op)) error {
	arr := strings.Split(configFilePath, "/")
	fileDir, fileName := "", ""
	if size := len(arr); size > 0 {
		fileDir = strings.Join(arr[0:size-1], "/")
		fileName = arr[size-1]
	}
	watcher := NewChangeWatcher(fileDir)
	addWatchFunc := func(target string) error {
		return watcher.RegisterWatcher(target, func(fileName string, optName fsnotify.Op) {
			fmt.Println(fileDir, target)
			if fileName != target {
				return
			}
			go func() {
				fmt.Println(fmt.Sprintf("configuration file change detected: %s, time: %s", optName, time.Now().Format("2006-01-02 15:04:05")))
				handler(optName)
			}()
		})
	}
	if err := addWatchFunc(fileName); err != nil {
		return err
	}
	watcher.PrintlnWatchCount()
	return nil
}
