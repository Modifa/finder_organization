package main

import (
	router "github.com/Modifa/finder_organization.git/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	gin.DisableConsoleColor()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(gin.ErrorLogger())
	// r.Use(CORSMiddleware())
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	// OPTIONS method for ReactJS
	config.AllowMethods = []string{"POST", "GET", "OPTIONS", "PUT"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "x-access-token", "content-type", "Content-Length", "Authorization", "Cache-Control"}
	config.ExposeHeaders = []string{"Content-Length"}
	r.Use(cors.New(config))
	// r.Use(gzip.Gzip(gzip.DefaultCompression))

	router.Init(r)

	return r
}
