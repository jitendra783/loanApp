package onboard
import (
	e "loanApp/pkg/error"
	"loanApp/pkg/logger"
	"loanApp/pkg/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func(m* mpinObj) SendOtp(c *gin.Context, id int) (string, error) {
	return "", nil
}

func (m *mpinObj) VerifyOtp(c *gin.Context, otp int )(bool , error){
	return true ,nil
}

func (m *mpinObj) MpinCheck(c *gin.Context, mpin string) (bool, error) {
	logger.Log(c).Debug("start")
	defer logger.Log().Debug("end")
	//id := c.GetString(config.USERID)
	id := utils.GenerateRandomString()
	//query formation type 1 
	query := "select mpin from schemaname.table_name where id = ? and mpin = ? "
	err := m.psql.Find(query, id, 11111)

	// query formation type to using sprintf
	// query formation using gorm functions 
	var users []User
	//sample for test id 1 and mpin 111111
	mpin1 := 1111111
	idtest := 1 
	err = m.psql.Table("users").Select("mpin").Where("mpin > ? and id = ? ", mpin1, idtest).Find(&users)
	if err != nil {
		log.Fatal("Failed to fetch the query:", err)
		return false, e.ErrorInfo[e.GetDBError]
	}
	return true, err.Error
}

func ( m *mpinObj) MpinUpdate(c *gin.Context ,id int) (User, error) {
	logger.Log(c).Debug("start")
	defer logger.Log(c).Debug("end")
	//mpin := "111111" 
	var user User
	// query := m.psql.Table("schemaname.tablename").Where("id = ?", 1).Update("mpin", mpin).Error

	// // Execute the query
	// defer rows.Close()
	// rows, err := m.psql.Query(query, id)
	// if err != nil {
	// 	log.Fatal("Failed to execute the query:", err)
	// }
	
	return user, nil
}
