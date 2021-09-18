package main

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"main/data"
	"main/src"
	"net/http"
	"os"
	"flag"
)

var router *gin.Engine

func main() {
	path := flag.String("file", "config.yaml", "Config file to parse.")
	port := flag.String("port", "", "Port to run Link hub on.")
	flag.Parse()

	if *port != "" {
		_ = os.Setenv("PORT", *port)
	}
	if os.Getenv("PORT") == "" {
		_ = os.Setenv("PORT", "10101")
	}

	config, err := src.LoadConfig(*path)
	if err != nil {
		panic(fmt.Sprintf("Fatal|Unable to load config %q|%v", *path, err))
	}

	runRouter(config)
}

func runRouter(config data.Config) {
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