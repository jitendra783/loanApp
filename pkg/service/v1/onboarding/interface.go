package onboarding

import (
	"loanApp/pkg/db"

	"github.com/gin-gonic/gin"
)

type onbaordObj struct {
	db db.DBLayer
}

func OnboardService(db db.DBLayer) OnboardGroup {
	return &onbaordObj{db: db}
}

type OnboardGroup interface {
	SendOtp(c *gin.Context) 
	VerifyOtp(c *gin.Context) 
	SetMpin(c *gin.Context) 
	VerifyMpin(c *gin.Context) 
}
