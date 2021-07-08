package http_server

func (d *LogHttpServer) initRouter() {
	in := d.internal.Group("/v1")
	{
		in.GET("/version", d.h.Version)
	}
}
