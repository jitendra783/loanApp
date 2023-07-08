package user

import (
	e "loanApp/pkg/error" 
)
type User struct{
	User_ID string `json:"user_id"`
	
}
type ApiResponse struct {
	Data []User
	Status  bool      `json:"status"`
	Errors  []e.Error `json:"error"`
	Message string    `json:"message"`
}
