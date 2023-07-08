package loan

import (
	"loanApp/pkg/logger"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (l *loanObj) LoanRegister(c *gin.Context) {
	log.Println("loanid", 123)
	//id := 1
	response, err := l.db.GetLoanBYID(c)
	if err != nil {
		logger.Log().Error("")
	}
	logger.Log().Debug("response ", zap.Any("res", response))
}
func (l *loanObj) GetLoanID(c *gin.Context) {
	log.Println("loanid", 123)
	//id := 1
	response, err := l.db.LoanDetailBy(c)
	if err != nil {
		logger.Log().Error("")
	}
	logger.Log().Debug("response ", zap.Any("res", response))
}
func (l *loanObj) LoanDetail(c *gin.Context) {
	log.Println("loanid", 123)
	//id := 1
	response, err := l.db.LoanDetailBy(c)
	if err != nil {
		logger.Log().Error("")
	}
	logger.Log().Debug("response ", zap.Any("res", response))
}
func (l *loanObj) LoanForclosureByID(c *gin.Context) {
	log.Println("loanid", 123)
	//id := 1
	response, err := l.db.LoanDetailBy(c)
	if err != nil {
		logger.Log().Error("")
	}
	logger.Log().Debug("response ", zap.Any("res", response))
}
func (l *loanObj) LoanUpdate(c *gin.Context) {
}
