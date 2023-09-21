package v2

import (
	"loanApp/pkg/db"
	"loanApp/pkg/service/v2/notification"

	"github.com/gin-gonic/gin"
)

type serviceObj struct {
	notification.NotificationGroup
}

type ServiceLayer interface {
	notification.NotificationGroup
	Status(*gin.Context)
}

func NewServiceObject(db db.DBLayer) ServiceLayer {
	return &serviceObj{
		notification.NotificationService(db),
	}
}