package user

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/goweb/model"
	"github.com/goweb/db"
	"net/http"
	"github.com/goweb/common"
)

func Save(c *gin.Context)  {
	name := c.Query("name")
	pwd := c.Query("password")
	email := c.Query("email")
	fmt.Printf("save user %s in controller\n",name)


	var u model.User


	if err := db.DB.Where(" name = ?", name).Find(&u).Error; err == nil {
		fmt.Printf("sorry,user name %s exist.\n", name)
		common.SendErrorJson("user " + name + " exist",c)
		return
	}
	u.Name = name
	u.Password = pwd

	if err := db.DB.Create(&u).Error; err != nil {
		fmt.Errorf("insert user error",err)
		return
	}
	go func() {
		sendEmail(email, "New user active","please active your account")
	}()

	c.JSON(http.StatusOK,gin.H{
		"result":"OK",
	})
}

func sendEmail(to string, title string, content string)  {


	common.SendMail(to,title,content)
}

func GetUserList(c *gin.Context)  {


	var userList  []model.User

	if err := db.DB.Find(&userList).Error; err != nil {
		fmt.Println("query user list error")
	}

	c.JSON(http.StatusOK,gin.H{
		"result":userList,
	})
}

func GetUserById( c *gin.Context)  {
	id := c.Param("id")
	var u model.User
	if err := db.DB.Where(" id = ?", id).First(&u).Error; err != nil {
		fmt.Println("query user by id error",err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result":u,
		"count":1,
	})
	}

func DeleteUserById(c *gin.Context)  {
	id := c.Param("id")
	var u model.User

	if err := db.DB.Where(" id = ?", id).Delete(&u).Error; err != nil {
		fmt.Println("delete user by id error")
	}
	c.JSON(http.StatusOK,gin.H{
		"count":1,
	})
}