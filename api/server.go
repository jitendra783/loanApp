package api

import (
	"context"
	"fmt"
	"loanApp/pkg/config"
	"loanApp/pkg/db"
	e "loanApp/pkg/error"
	"loanApp/pkg/logger"
	serv "loanApp/pkg/service"
	"log"
	"net/http"
	"strconv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
)

var srv *http.Server
var ctx context.Context
var databases []*gorm.DB

func Start() error {
	ctx = context.Background()

	config := config.GetConfig()
	logLevel, err := strconv.Atoi(config.GetString("log.Level"))
	if err != nil {
		log.Fatal("Invalid log config: ", err)
	}
	logger.LoggerInit(config.GetString("log.path"), zapcore.Level(logLevel))
	e.ErrorInit()

	mysqlConn, mysqlConnectionSql, err := db.MysqlConnect()
	if err != nil {
		logger.Log().Error("Failed to connect mysql database", zap.Error(err))
		return err
	}
	databases = make([]*gorm.DB, 0)
	databases = append(databases, mysqlConn)

	dbObj := db.NewDbObj(mysqlConnectionSql)

	serviceObj := serv.NewServiceGroupObject(dbObj)

	// if err := serviceObj.GetV1Service().Status(); err != nil {
	// 	return err
	// }
	startRouter(serviceObj)
	return nil
}

func startRouter(obj serv.ServiceGroupLayer) {
	srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", config.GetConfig().GetInt("server.port")),
		Handler: Router(obj, logger.Log()), //getRouter set the api specs for version-1 routes
	}
	// run api router
	logger.Log().Info("starting router")

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Log().Fatal("Error starting server", zap.Error(err))
	}
	logger.Log().Info("server working")
}
