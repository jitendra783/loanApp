package notification

import (
	"loanApp/pkg/db"

	"github.com/gin-gonic/gin"
)

type notifyObj struct{
	db db.DBLayer
}
func NotificationService  (db db.DBLayer) NotificationGroup {
	return &notifyObj{db:db}
}
type NotificationGroup interface {
	Emi_due(c *gin.Context)
	Getnotification(c *gin.Context)
	Pre_approved_offer(c *gin.Context)
}