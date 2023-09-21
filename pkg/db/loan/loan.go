package loan

import (
	"loanApp/pkg/logger"

	"github.com/gin-gonic/gin"
)

func (u *loanObj) LoanDeleteByID(c *gin.Context) error {
	logger.Log().Debug("start")
	defer logger.Log().Debug("end")
	return nil

}
func (u *loanObj) loandetailByID(c *gin.Context) (LoanDetails, error) {
	logger.Log().Debug("start")
	defer logger.Log().Debug("end")
	// id := 1
	var loan LoanDetails
	// if err := u.msql.Table("loan").Find(&loan, "loan_id = ? ", id); err.Error != nil {
	// 	logger.Log().Error("Failed to find the loan", zap.Error(err.Error))
	// 	return loan, err.Error
	// } else if err.RowsAffected == 0 {
	// 	logger.Log().Error("loan doesn't exists", zap.Error(err.Error))
	// 	return loan, e.ErrorInfo["NoDataFound"]
	// }
	return loan, nil

}
func (u *loanObj) LoanDetailBy(c *gin.Context) (LoanDetails, error) {
	logger.Log().Debug("start")
	defer logger.Log().Debug("end")
	//id := 1
	var loan LoanDetails
	// if err := u.msql.Table("loan").Find(&loan, "loan_id = ? ", id); err.Error != nil {
	// 	logger.Log().Error("Failed to find the loan", zap.Error(err.Error))
	// 	return loan, err.Error
	// } else if err.RowsAffected == 0 {
	// 	logger.Log().Error("loan doesn't exists", zap.Error(err.Error))
	// 	return loan, e.ErrorInfo["NoDataFound"]
	// }
	return loan, nil

}

func (u *loanObj) GetLoanBYID(c *gin.Context) (LoanDetails, error) {
	logger.Log().Debug("start")
	defer logger.Log().Debug("end")
	// id := 1
	var loan LoanDetails
	// if err := u.msql.Table("loan").Find(&loan, "loan_id = ? ", id); err.Error != nil {
	// 	logger.Log().Error("Failed to find the loan", zap.Error(err.Error))
	// 	return loan, err.Error
	// } else if err.RowsAffected == 0 {
	// 	logger.Log().Error("loan doesn't exists", zap.Error(err.Error))
	// 	return loan, e.ErrorInfo["NoDataFound"]
	// }
	return loan, nil

}
