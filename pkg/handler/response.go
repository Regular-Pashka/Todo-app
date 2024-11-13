package handler 

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type err struct {
	Message string `json"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message int) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, error{message})
}