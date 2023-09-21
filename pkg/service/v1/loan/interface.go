package loan

import (
	"loanApp/pkg/db"

	"github.com/gin-gonic/gin"
)

type loanObj struct {
	db db.DBLayer
}
type LoanGroup interface {
	LoanRegister(c *gin.Context)
	GetLoanID(c *gin.Context)
	LoanDetail(c *gin.Context)
	LoanForclosureByID(c *gin.Context)
	LoanUpdate(c *gin.Context)
}

func LoanService(db db.DBLayer) LoanGroup {
	return &loanObj{}
}
