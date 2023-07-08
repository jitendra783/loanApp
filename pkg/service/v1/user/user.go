package user

import (
	"loanApp/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (u *userObj) UserRegister(c *gin.Context) {
	logger.Log().Debug("start")
	defer logger.Log().Debug("end")

}
func (u *userObj) GetUserID(c *gin.Context) {
	logger.Log().Debug("start")
	defer logger.Log().Debug("end")
	user, err := u.db.GetUserDetailByID(c)
    if err !=nil{
    logger.Log().Error("error ", zap.Error(err))
	}
	c.JSON(http.StatusOK,gin.H{"massage":"Succcess", "response":user})

}
func (u *userObj) UserDetail(c *gin.Context) {
	logger.Log().Debug("start")
	defer logger.Log().Debug("end")
	id := c.Param("id")
	logger.Log().Debug("id which need to delete", zap.Any("id",id))
	user, err :=u.db.GetUserDetailByID(c)

	if err != nil{
		logger.Log().Error("")
	}
	logger.Log().Debug("user details", zap.Any("user",user))


}
func (u *userObj) UserDeleteByID(c *gin.Context) {
	logger.Log().Debug("start")
	defer logger.Log().Debug("end")
	id := c.Param("id")
	
	logger.Log().Debug("id which need to delete", zap.Any("id",id))
	err :=u.db.UserDeleteByID(c)
	if err != nil{
		logger.Log().Error("")
	}


}
func (u *userObj) UserUpdate(c *gin.Context) {
	logger.Log().Debug("start")
	defer logger.Log().Debug("end")

}
