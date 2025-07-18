package endpoints

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Endpoints interface {
	Ping(c *gin.Context)
}

type endpoints struct {
	param string
}

func NewEndpoints(param string) Endpoints {
	return &endpoints{param}
}

func (provider *endpoints) Ping(c *gin.Context) {
	fmt.Sprintln(provider.param)

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
