package response

import "github.com/Leonardo-Antonio/api.process-vtex/entity"

var Success = func(
	message string, data interface{},
) *entity.Response {
	return &entity.Response{
		MessageType: "SUCCESS",
		Message:     message,
		Err:         false,
		Data:        data,
	}
}

var Error = func(
	message string, data interface{},
) *entity.Response {
	return &entity.Response{
		MessageType: "ERROR",
		Message:     message,
		Err:         true,
		Data:        data,
	}
}
