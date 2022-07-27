package messages

type Message struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

var Success = Message{Code: "200", Message: "Success"}

var TelErrFailedToLogin = Message{Code: "500", Message: "Email or Password is wrong"}
var TelErrShouldBindError = Message{Code: "500", Message: "Function Should Bind Error"}
var TelErrUserNotFound = Message{Code: "500", Message: "Err User Not Found"}
var TelErrEmailAlreadyUsed = Message{Code: "500", Message: "Email Already Used"}
var TelErrRevocerRoute = Message{Code: "500", Message: "Tel Err Revocer Route"}
var TelErrItemNotFound = Message{Code: "500", Message: "Tel Err Item Not Found"}
var TelSavedSuccessfully = Message{Code: "500", Message: "Saved Successfully"}
var TelErrFailedToSave = Message{Code: "500", Message: "Failed To Save"}
var TelErrPageNotFound = Message{Code: "500", Message: "Page Not Found"}
