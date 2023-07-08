package v1

import (
	"loanApp/pkg/db"
	"loanApp/pkg/service/v1/loan"
	"loanApp/pkg/service/v1/user"

	"github.com/gin-gonic/gin"
)
type serviceObj struct {
	user.UserGroup
	loan.LoanGroup
}

func NewServiceObject(db db.DBLayer ) ServiceLayer{
	return &serviceObj{
		user.UserService(db),
		loan.LoanService(db),
	}
}
type ServiceLayer interface{
    user.UserGroup
	loan.LoanGroup
	// Status work as health check for service
	Status(*gin.Context)
	GetCurrentTime(*gin.Context)
}

