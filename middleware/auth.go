package middleware

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func Auth(c *gin.Context)  {

	fmt.Println(" auth op")
	c.Next()

}
