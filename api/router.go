package api

import (
	"loanApp/pkg/config"
	"loanApp/pkg/service"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Router(obj service.ServiceGroupLayer, logger *zap.Logger) *gin.Engine {
	router := gin.New()
	router.Use(customLogger(logger))
	router.Use(gin.Recovery())
	router.GET("/health", obj.GetV1Service().Status)
	userGroup := router.Group("user")
	{
		userGroup.POST("/register", obj.GetV1Service().UserRegister)
		userGroup.GET("/getdetail", obj.GetV1Service().GetUserID)
		userGroup.PUT("/update", obj.GetV1Service().UserUpdate)
		userGroup.DELETE("/delete", obj.GetV1Service().UserDeleteByID)
	}
	loanGroup := router.Group("loan")
	{
		loanGroup.GET("/loandetail", obj.GetV1Service().LoanDetail)
		loanGroup.GET("/loanid", obj.GetV1Service().GetLoanID)
	}

	notificationGroup := router.Group("Notification")
	{
		notificationGroup.GET("/notify", obj.GetV2Service().Getnotification)
		notificationGroup.GET("/emidue", obj.GetV2Service().Emi_due)
		notificationGroup.GET("/offer", obj.GetV2Service().Pre_approved_offer)
		notificationGroup.GET("/status", obj.GetV2Service().Status)
	}

	return router
}

func customLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		if c.FullPath() != "/health" {
			latency := time.Since(start).Milliseconds()
			userID := c.GetString(config.USERID)
			uID := c.GetString(config.REQUESTID)
			ucc := c.GetString(config.UCC)
			logger.Info(path,
				zap.String("requestID", uID),
				zap.String("ucc", ucc),
				zap.String("userId", userID),
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
				zap.Int64("latency", latency),
			)
		}
	}
}
