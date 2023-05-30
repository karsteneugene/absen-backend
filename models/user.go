package models

type User struct {
	ID        uint   `json:"id" xml:"id" form:"id"`
	Username  string `json:"username" xml:"username" form:"username"`
	Password  string `json:"password" xml:"password" form:"password"`
	FirstName string `json:"first_name" xml:"first_name" form:"first_name"`
	LastName  string `json:"last_name" xml:"last_name" form:"last_name"`
}
