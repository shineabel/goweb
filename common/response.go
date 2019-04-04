package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendErrorJson(message string, args... interface{})  {

	if len(args) == 0 {
		panic("miss *gin.context")
	}

	var c *gin.Context

	var en int
	if len(args) == 1 {
		ctx , ok := args[0].(*gin.Context)
		if !ok {
			panic("miss *gin.context")
		}

		c = ctx
	} else if len(args) == 2{
		errorNo, ok := args[0].(int)
		if !ok {
			panic("error no error")
		}
		en = errorNo

		ctx , ok := args[1].(*gin.Context)
		if !ok {
			panic("miss *gin.context")
		}

		c = ctx
	}
	c.JSON(http.StatusOK, gin.H{
		"code":en,
		"msg":message,
		"data":gin.H{},
	})
	c.Abort()
}
