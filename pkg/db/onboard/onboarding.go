package user

import (
	e "loanApp/pkg/error"
	"loanApp/pkg/logger"
	"loanApp/pkg/utils"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (u *userObj) CreateUser(c *gin.Context, userinfo UserForm) (User, error) {
	logger.Log().Debug("start")
	defer logger.Log().Debug("end")
	//id := c.GetString(config.USERID)
	var user User
	id := utils.GenerateRandomString()
	query := "insert into mylaon.user (id ,first_name, last_name,email,mobile_no,pan_card,mpin ) value( ? ,?,NULL,NULL,?,NULL,NULL)"
	// Execute the query
	_, err := u.msql.Exec(query, id, userinfo.Name, userinfo.Mobile)
	logger.Log().Debug("query", zap.Any("q", query))
	if err != nil {
		log.Fatal("Failed to execute the query:", err)
		return user, err
	}
	//undefined error while checking in db
	// user, err = u.GetUserByID(c,id)
	// if err != nil{
	// 	log.Fatal("Failed to execute the query:", err)
	// 	return user, err
	// }
	user.Id = id
	user.Create_at = time.Now()
	user.First_Name = userinfo.Name
	user.Mobile_no = userinfo.Mobile

	return user, nil
}

// var user User
func (u *userObj) GetUserByID(c *gin.Context, id string) (User, error) {
	logger.Log().Debug("start")
	defer logger.Log().Debug("end")
	//id := c.GetString(config.USERID)
	var user User // Assuming User is the struct representing the user entity
	//id = "a"
	query := "SELECT id,mobile_no,first_name FROM mylaon.user WHERE id = ?"

	// Execute the query
	rows, err := u.msql.Query(query, id)
	logger.Log().Debug("query", zap.Any("q", query))
	if err != nil {
		log.Fatal("Failed to execute the query:", err)
	}
	defer rows.Close()
	if !rows.Next() {
		return user, e.ErrorInfo[e.NoDataFound]
	} else {
		err := rows.Scan(&user.Id,&user.Mobile_no,&user.First_Name) // Assuming UserID, Name, Email are the columns in the 'user' table
		if err != nil {
			logger.Log().Error("Failed to fetch the user", zap.Error(err))
			return user, err
		}
	}
	return user, nil
}
func (u *userObj) UpdateUserByID(c *gin.Context) (User, error) {
	logger.Log().Debug("start")
	defer logger.Log().Debug("end")
	id := 1
	query := "SELECT * FROM user WHERE user_id = ?"

	// Execute the query
	rows, err := u.msql.Query(query, id)
	if err != nil {
		log.Fatal("Failed to execute the query:", err)
	}
	defer rows.Close()

	// Process the result
	var user User // Assuming User is the struct representing the user entity
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.First_Name, &user.Last_name) // Assuming UserID, Name, Email are the columns in the 'user' table
		if err != nil {
			log.Fatal("Failed to scan row:", err)
		}
	} else {
		log.Println("User not found")
	}

	// if err := u.msql.Table("user").Find(&user, "user_id", id); err.Error != nil {
	// 	logger.Log().Error("Failed to find the user", zap.Error(err.Error))
	// 	return user, err.Error
	// } else if err.RowsAffected == 0 {
	// 	logger.Log().Error("User doesn't exists", zap.Error(err.Error))
	// 	return user, e.ErrorInfo["NoDataFound"]
	// }
	return user, nil
}
func (u *userObj) UserDeleteByID(c *gin.Context) error {
	logger.Log().Debug("start")
	defer logger.Log().Debug("end")
	// id := c.GetString(config.USERID)
	// if err := u.msql.Table("users").Delete(&User{}, id); err.Error != nil {
	// 	logger.Log().Error("Failed to delete the user", zap.Error(err.Error))
	// 	return err.Error
	// } else if err.RowsAffected == 0 {
	// 	logger.Log().Error("User doesn't exists", zap.Error(err.Error))
	// 	return e.ErrorInfo["NoDataFound"]
	// }
	return nil

}
