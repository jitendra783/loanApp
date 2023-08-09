package user

import (
	dbuser "loanApp/pkg/db/user"
	e "loanApp/pkg/error"
	"loanApp/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

func (u *userObj) UserRegister(c *gin.Context) {
	logger.Log().Debug("start")
	defer logger.Log().Debug("end")
	var response ApiResponse
	var userinfo dbuser.UserForm
	if err := c.BindJSON(&userinfo); err != nil {
		response.Errors = append(response.Errors, e.ErrorInfo[e.BadRequest].GetErrorDetails(""))
		c.JSON(http.StatusBadRequest, response)
		return
	}
	resp, err := u.db.CreateUser(c, userinfo)
	if err != nil {
		logger.Log().Error("error ", zap.Error(err))
		response.Errors = append(response.Errors, e.ErrorInfo[e.AddDBError].GetErrorDetails(""))
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response.Data = append(response.Data, resp)
	response.Status = true
	response.Message = "user Created succefully"
	c.JSON(http.StatusOK, response)
}

func (u *userObj) GetUserByID(c *gin.Context) {
	logger.Log().Debug("start")
	defer logger.Log().Debug("end")
	var response ApiResponse
	var user GetUserParam
	if err := c.BindUri(&user); err != nil {
		response.Errors = append(response.Errors, e.ErrorInfo[e.BadRequest].GetErrorDetails(""))
		c.JSON(http.StatusBadRequest,response)
		return
	}
	resp, err := u.db.GetUserByID(c, user.Id)
	if err != nil {
		if err.Error() == e.NoDataFound {
			logger.Log().Error("no data found ", zap.Error(err))
			response.Errors = append(response.Errors, e.ErrorInfo[e.NoDataFound].GetErrorDetails(""))
			c.JSON(http.StatusOK,response)
			return
		} else {
			logger.Log().Error("user not available", zap.Error(err))
			response.Errors = append(response.Errors, e.ErrorInfo[e.GetDBError].GetErrorDetails(""))
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}
	response.Data = append(response.Data, resp)
	response.Message = "user Founded"
	response.Status = true
	c.JSON(http.StatusOK, response)

}
func (u *userObj) UserDetail(c *gin.Context) {
	logger.Log().Debug("start")
	defer logger.Log().Debug("end")
	var response ApiResponse
	var user GetUserParam
	if err := c.BindUri(&user); err != nil {
		response.Errors = append(response.Errors, e.ErrorInfo[e.BadRequest].GetErrorDetails(""))
	}

	resp, err := u.db.GetUserByID(c, user.Id)
	if err.Error() == e.NoDataFound {
		logger.Log().Error("no data found ", zap.Error(err))
		response.Errors = append(response.Errors, e.ErrorInfo[e.NoDataFound].GetErrorDetails(""))
		c.JSON(http.StatusBadRequest, response)
	} else {
		logger.Log().Error("user not available", zap.Error(err))
		response.Errors = append(response.Errors, e.ErrorInfo[e.GetDBError].GetErrorDetails(""))
		c.JSON(http.StatusInternalServerError, response)
	}
	response.Data = append(response.Data, resp)
	response.Message = "user Founded"
	response.Status = true
	c.JSON(http.StatusOK, response)
}
func (u *userObj) UserDeleteByID(c *gin.Context) {
	logger.Log().Debug("start")
	defer logger.Log().Debug("end")
	id := c.Param("id")

	logger.Log().Debug("id which need to delete", zap.Any("id", id))
	err := u.db.UserDeleteByID(c)
	if err != nil {
		logger.Log().Error("")
	}

}
func (u *userObj) UserUpdate(c *gin.Context) {
	logger.Log().Debug("start")
	defer logger.Log().Debug("end")

}
