package db

import (
	"database/sql"
	"loanApp/pkg/db/loan"
	"loanApp/pkg/db/onboard"
	"loanApp/pkg/db/user"

	"gorm.io/gorm"
)

type dbObj struct {
	user.UserGroup
	loan.LoanGroup
	onboard.OnboardGroup
}
type DBLayer interface {
	user.UserGroup
	loan.LoanGroup
	onboard.OnboardGroup
}

func NewDbObj(mysqlDB *sql.DB, psql *gorm.DB) DBLayer {
	temp := &dbObj{user.NewUserDBGroup(mysqlDB),
		loan.NewLoanDBGroup(mysqlDB),
		onboard.NewOnboardGroup(psql),
	}
	return temp
}
