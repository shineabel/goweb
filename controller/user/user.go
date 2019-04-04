package user

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/goweb/model"
	"github.com/goweb/db"
	"net/http"
)

func Save(c *gin.Context)  {
	name := c.Query("name")
	pwd := c.Query("password")
	fmt.Printf("save user %s in controller\n",name)


	var u model.User


	if err := db.DB.Where(" name = ?", name).Find(&u).Error; err == nil {
		fmt.Printf("sorry,user name %s exist.\n", name)
		return
	}
	u.Name = name
	u.Password = pwd

	if err := db.DB.Create(&u).Error; err != nil {
		fmt.Errorf("insert user error",err)
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"result":"OK",
	})
}
