package user

import (
	e "loanApp/pkg/error"
	dbuser "loanApp/pkg/db/user"
)

type User struct {
	User_ID string `json:"user_id"`
}
type ApiResponse struct {
	Data    []dbuser.User
	Status  bool      `json:"status"`
	Errors  []e.Error `json:"error"`
	Message string    `json:"message"`
}

type GetUserParam struct{
	Id string `uri:"id" binding:"required"`
}
