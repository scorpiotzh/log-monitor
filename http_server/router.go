package http_server

func (d *LogHttpServer) initRouter() {
	in := d.internal.Group("/v1")
	{
		in.GET("/version", d.h.Version)
		in.POST("/search", d.h.SearchLogApiInfo)
		in.POST("/push/log", d.h.PushLog)
	}
}
