package onboard

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type mpinObj struct {
	psql *gorm.DB
}
type OnboardGroup interface {
	MpinCheck(c *gin.Context, mpin string) (bool, error)
	MpinUpdate(c *gin.Context, id int) (User, error)
}

func NewOnboardGroup(db *gorm.DB) OnboardGroup {
	return &mpinObj{psql: db}
}
