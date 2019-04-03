package user

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/goweb/model"
)

func Save(c *gin.Context)  {
	fmt.Printf("save user in controller")

	type UserData struct {
		Name string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var ud UserData
	if err := c.ShouldBindWith(&ud, binding.JSON); err != nil {
		fmt.Println(err)
		return
	}
	var u model.User


}
