package db

import (
	"database/sql"
	"loanApp/pkg/db/loan"
	"loanApp/pkg/db/user"
)

type dbObj struct {
	user.UserGroup
	loan.LoanGroup
}
type DBLayer interface {
	user.UserGroup
	loan.LoanGroup
}

func NewDbObj(mysqlDB *sql.DB) DBLayer {
	temp := &dbObj{user.NewUserDBGroup(mysqlDB),
		loan.NewLoanDBGroup(mysqlDB),
	}
	return temp
}
