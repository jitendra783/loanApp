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
	GetUserDetailByID(c *gin.Context) (User, error)
	UpdateUserByID(c *gin.Context) (User, error)
}

func NewUserDBGroup(db *sql.DB) UserGroup {
	return &userObj{msql: db}
}
