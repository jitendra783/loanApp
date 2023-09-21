package v1

import (
	"loanApp/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *serviceObj) Status(c *gin.Context) {

	c.String(http.StatusOK, "Working!")
}
func (s *serviceObj) GetCurrentTime(c *gin.Context) {
	logger.Log(c).Debug("start")
	defer logger.Log(c).Debug("end")
    
}
