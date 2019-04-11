package main

import "fmt"

type User struct {
	Id int
	Name string
	Age int
}

var list []User

func AddUser(u User)  {

	list = append(list,u)
}

func RemoveUser(user User)  {
	for index, u := range list{
		if u.Id == user.Id {
			if len(list) == 1 {
				list = make([]User,0)
			} else if len(list) > 1 {
				if index == len(list) - 1 {
					list = list[:index]
				} else {
					list = append(list[:index -1 ],list[index + 1 :]...)
				}
			} else {
				fmt.Println("there is no user now")
			}
		}
	}
}

func UpdateUser(user User)  {
	for _, u := range list{
		if u.Id == user.Id {
			u.Name = user.Name
			u.Age = user.Age

			break
		}
	}
}

func Retrive()  {

	for index, u := range list{
		fmt.Println(index,u.Id,u.Name,u.Age)
	}
}
func main() {
	u1 := User{1,"ljw",33}
	u2 := User{2,"xxx",32}

	fmt.Println("before add:",list)

	//增
	AddUser(u1)
	AddUser(u2)
	fmt.Println("after add user:",list)

	u2 = User{2,"xianming",18}

	//删
	RemoveUser(u2)

	fmt.Println("remove user2:",list)



	u3 := User{3,"cat",30}
	AddUser(u3)

	//查
	Retrive()
}