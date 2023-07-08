package error

type Error struct {
	ErrorName string `json:"error"`
	Description string `json:"description"`
	ErrorCode   int    `json:"errorCode"`
}

const(
	NoDataFound            string = "NoDataFound"
	BadRequest             string = "BadRequest"
	GetDBError             string = "GetDBError"
	AddDBError             string = "AddDBError"
	DelDBError             string = "DelDBError"
	DefaultError           string = "DefaultError"
	InternalServerError    string = "InternalServerError"
)