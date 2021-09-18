package main

import (
	"fmt"
	"main/data"
	"main/src"
	"net/http"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
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
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK, "index.html", gin.H(structs.Map(config)),
		)
	})

	router.Run()
}