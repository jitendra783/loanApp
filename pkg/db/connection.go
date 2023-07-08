package db

import (
	"context"
	"database/sql"
	e "loanApp/pkg/error"
	elog "loanApp/pkg/logger"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	// "gorm.io/gorm/schema"
	"fmt"
	"time"
)

type dbLogger struct{}

func MysqlConnect() (*gorm.DB, *sql.DB, error) {
	username := "root"
	password := "Jitu@2783"
	host := "localhost"
	port := "3306"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8&parseTime=True&loc=Local",
		username, password, host, port)
	// Connect to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		elog.Log().Error("failed to connect mysql connection", zap.Error(err), zap.String("connStr", dsn))
		return nil, db, e.ErrorInfo["MysqlDBConnError"]
	}
	db1, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// NamingStrategy: schema.NamingStrategy{
		// 	TablePrefix: "myloan.",
		// 	SingularTable: true,
		// },
		SkipDefaultTransaction: true, /// Removed defuault transaction used by GORM to faster the query and execution
		Logger:                 customLogger(),
	})
	// if err != nil {
	// 	elog.Log().Error("failed to connect mysql connection", zap.Error(err), zap.String("connStr", dsn))
	// 	return nil,db, e.ErrorInfo["MysqlDBConnError"]
	// }
	elog.Log().Info("Mysql Database Connected")
	return db1, db, nil
}

func customLogger() logger.Interface {
	return dbLogger{}
}

func (d dbLogger) Error(ctx context.Context, data string, others ...interface{}) {
	elog.Log(ctx).Info("database", zap.String("error", data), zap.Any("description", others))
}

func (d dbLogger) Info(ctx context.Context, data string, others ...interface{}) {
	elog.Log(ctx).Info("database", zap.String("msg", data), zap.Any("description", others))
}

func (d dbLogger) Warn(ctx context.Context, data string, others ...interface{}) {
	elog.Log(ctx).Info("database", zap.String("msg", data), zap.Any("description", others))
}

func (d dbLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	query, others := fc()
	if err != nil {
		elog.Log(ctx).Info("database", zap.String("query", query), zap.Any("rows-affected", others), zap.Error(err))
	} else {
		elog.Log(ctx).Info("database", zap.String("query", query), zap.Any("rows-affected", others))
	}
}

func (d dbLogger) LogMode(l logger.LogLevel) logger.Interface {
	return d
}
