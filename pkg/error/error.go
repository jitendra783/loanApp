package error

import "loanApp/pkg/logger"

var ErrorInfo map[string]*Error

func ErrorInit() {
	// define all error
	ErrorInfo = make(map[string]*Error)

	ErrorInfo[NoDataFound] = &Error{ErrorName: NoDataFound, Description: "no data fund system failure", ErrorCode: 1006}
	ErrorInfo[BadRequest] = &Error{ErrorName: BadRequest, Description: "missing value or inputs", ErrorCode: 1010}
	ErrorInfo[InternalServerError] = &Error{ErrorName: InternalServerError, Description: "internal server error", ErrorCode: 1021}
	ErrorInfo[GetDBError] = &Error{ErrorName: "DBError", Description: "Failed to get data", ErrorCode: 1007}
	ErrorInfo[AddDBError] = &Error{ErrorName: "DBError", Description: "Failed to update/insert data", ErrorCode: 1008}
	ErrorInfo[DelDBError] = &Error{ErrorName: "DBError", Description: "Failed to Delete the data", ErrorCode: 1009}
	logger.Log().Info("ErrorInit successful")

}
func (e *Error) GetErrorDetails(errMsg string) Error {
	if e == nil {
		return *ErrorInfo[DefaultError]
	}
	err := *e
	if errMsg != "" {
		err.Description = err.Description + " " + errMsg
	}
	return err
}
func (e *Error) Error() string {
	if e != nil {
		return " undefind error"
	}
	return e.ErrorName
}

func (e *Error) GetErrorSlice(str string) (output []Error) {
	output = append(output, e.GetErrorDetails(str))
	return
}
