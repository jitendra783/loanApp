package loanapp

/*
variable 
	scope : function level variable
	declarion 
		var1 := any value (First letter lowercase)
		var var1 typeofVariable

	global 
		scope : entire group with name and in code with pkgname.varialbleName
		declarion :
			var Var1 variable type (First letter must be Uppercase)
			var Var1 = any value

map
	uses and where to use
	var1 := make(map[datatype]datatype){

	}
    var Var1 map[datatype]datatype{}
	var 

interface
	learn the use of interface and implementation


constants
	for each group create a constant file in
	which put all the constants and use from there
	also put all the global variables


logging
	log type and usables
	log which is predefined in go
	zap which is defined by uber mostly used

error
	Standard Error
	custom Error
		ErrorName string `json:"error"`
		Description string `json:"description"`
		ErrorCode   int    `json:"errorCode"
			

frameWork
	gin
	gorrila mux

router(Api version)-> service(version) -> db(tempate)
	Architacture :
		Api request
		    router >
				service (version) >
					binding of request
					db >
						functionalities to fetch data
					return from db
				return from service
		return Api response

Router :

Server:

httpReq
	Url
	Method
	Paylaod
	body
	header
		authorization
		
	client -Do


Binding (request-form)
	bind the requst type
Api response
		Data interface{} 	`json:"data"`
		Messege string 		`json:"messege"`
		Status   bool		`json:"status"`
		Errors	[]Error		`json:"error"`

database connectivity
	query optimization
	ping
	connection
		establish conn
		pass as arg

middleware
	Authorization
		Token & refresh Token
		Get and Read
		Set

testing
	unit test

grpc service
    what is it
	how to use
	where to use
	why use it

kafka stream
	what is it
	how to use
	where to use
	why use it

kubernate/ docker / yaml
	what is it
	how to use
	where to use
	why use it

caching
	redis
		what is it 
		how to use
		where to use
		why use it

*/

