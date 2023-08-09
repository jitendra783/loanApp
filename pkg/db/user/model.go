package user

import "time"

type User struct {
	Id         string    `json:"user_id" gorm:"primaryKey"`
	First_Name string    `json:"firstName"`
	Last_name  string    `json:"lastName"`
	Mobile_no  int64    `json:"mobileNumber"`
	Email_id   string    `json:"emailId"`
	Create_at  time.Time `json:"created_at"`
}
type UserForm struct {
	Name   string `json:"name"`
	Dob    string `json:"dob"`
	Mobile int64 `json:"mobile"`
}
