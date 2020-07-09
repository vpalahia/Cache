package main

import (
	_ "net/http/pprof"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/vpalahia/Cache/api"
	"github.com/vpalahia/Cache/cache"
)

func main() {
	cache.InitializeCache()

	r := gin.Default()
	pprof.Register(r)
	r.GET("/cache/fetch", api.Fetch)
	r.POST("/cache/insert", api.Insert)
	r.Run(":9091")
}
