package main

import (
	"example.com/exemplo/endpoints"
	"github.com/gin-gonic/gin"
)

func main() {
	param := "Hello, World!"

	app := gin.Default()
	router := endpoints.NewEndpoints(param)

	app.GET("/ping", router.Ping)
	app.Run()
}
