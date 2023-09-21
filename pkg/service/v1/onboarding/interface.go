package onboarding

import (
	"loanApp/pkg/db"

	"github.com/gin-gonic/gin"
)

type onbaord struct {
	db db.DBLayer
}

func OnboardService(db db.DBLayer) OnboardGroup {
	return onbaord{db: db}
}

type OnboardGroup interface {
	SendOtp() error
	VerifyOtp() error
	MpinGeneration(c *gin.Context) error
	MpinVarification(c *gin.Context) error
}
