package helper

import "strings"

// Create a new struct for the response data
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

//EmptyObj object is used when data doesnt want to be null on json
type EmptyObject struct {
}

// SuccessResponse returns a success response with the given data
func SuccessResponse(code int, message string, data interface{}) Response {
	return Response{
		Code:    code,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
}

// ErrorResponse returns an error response with the given data
func ErrorsResponse(code int, message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	return Response{
		Code:    code,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
}
