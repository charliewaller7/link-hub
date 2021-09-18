package main

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"main/data"
	"main/src"
	"net/http"
	"os"
)

var router *gin.Engine

func main() {
	// TODO - Path as cli param
	path := "config.yaml"

	config, err := src.LoadConfig(path)
	if err != nil {
		panic(fmt.Sprintf("Fatal|Unable to load config %q|%v", path, err))
	}

	runRouter(config)
}

func runRouter(config data.Config) {
	if p := os.Getenv("PORT"); p == "" {
		_ = os.Setenv("PORT", "10101")
	}

	router = gin.Default()
	router.LoadHTMLGlob("static/templates/*")
	router.Static("/static", "./static")
	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK, "index.html", gin.H(structs.Map(config)),
		)
	})

	router.Run()
}