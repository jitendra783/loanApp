package user

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type userObj struct {
	msql *sql.DB
}
type UserGroup interface {
	UserDeleteByID(c *gin.Context) error
	GetUserByID(c *gin.Context,id string) (User, error)
	UpdateUserByID(c *gin.Context) (User, error)
	CreateUser(c *gin.Context, userinfo UserForm) (User, error)
}

func NewUserDBGroup(db *sql.DB) UserGroup {
	return &userObj{msql: db}
}
