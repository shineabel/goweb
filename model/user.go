package model

type User struct {
	Id uint `gorm:"primary_key;auto_increment" json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	Email string `json:"email"`
}

func (u *User) TableName() string {
	return "user_info"
}