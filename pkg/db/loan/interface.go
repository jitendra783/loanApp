package loan

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type loanObj struct {
	msql *sql.DB
}

type LoanGroup interface {
	LoanDeleteByID(c *gin.Context) error
	loandetailByID(c *gin.Context) (LoanDetails, error)
	LoanDetailBy(c *gin.Context) (LoanDetails, error)
	GetLoanBYID(c *gin.Context) (LoanDetails, error)
}

func NewLoanDBGroup(db *sql.DB) LoanGroup {
	return &loanObj{msql: db}
}
