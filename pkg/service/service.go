package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *service) Status(c *gin.Context){
	c.String(http.StatusOK, "working")
}