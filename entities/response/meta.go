package response

import (
	"akrab-bangkit2022-api/entities/teleconsultation/messages"
)


type Meta struct {
	messages.Message
	RequestID string `json:"request_id"`
}

type respMeta struct {
	Success                messages.Message
	TelErrFailedToLogin    messages.Message
	TelErrShouldBindError  messages.Message
	TelErrUserNotFound     messages.Message
	TelErrEmailAlreadyUsed messages.Message
	TelErrRevocerRoute     messages.Message
	TelErrPageNotFound     messages.Message
	TelErrItemNotFound     messages.Message
	TelSavedSuccessfully   messages.Message
	TelErrFailedToSave     messages.Message
}

var RespMeta = respMeta{
	Success:                messages.Success,
	TelErrFailedToLogin:   	messages.TelErrFailedToLogin,
	TelErrShouldBindError:  messages.TelErrShouldBindError,
	TelErrUserNotFound:     messages.TelErrUserNotFound,
	TelErrEmailAlreadyUsed: messages.TelErrEmailAlreadyUsed,
	TelErrRevocerRoute:     messages.TelErrRevocerRoute, 
	TelErrPageNotFound:     messages.TelErrPageNotFound, 
	TelErrItemNotFound:     messages.TelErrItemNotFound,
	TelSavedSuccessfully:   messages.TelSavedSuccessfully, 
	TelErrFailedToSave:     messages.TelErrFailedToSave,
	
}
