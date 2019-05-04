package main

import (
	"gin-blog/config"
	"gin-blog/server/routers"
	"github.com/gin-gonic/gin"
	"net"
)

func init()  {
	config.NewContext()
}

func main()  {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	routers.InitializeRouter(r)

	addr := net.JoinHostPort(config.ProtAddr, config.ProtPort)
	r.Run(addr)
}