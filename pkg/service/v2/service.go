package v2

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *serviceObj) Status(c *gin.Context) {

	c.String(http.StatusOK, "Working!")
}
