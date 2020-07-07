package utils

type Responses struct {
	Status  int
	Message string
	Data    interface{}
}

type OtherResponses struct {
	Status  int
	Message string
}

func Response(status int, message string, data interface{}) Responses {
	var response Responses
	response.Status = status
	response.Message = message
	response.Data = data
	return response
}

func OtherResponse(status int, message string) OtherResponses {
	var response OtherResponses
	response.Status = status
	response.Message = message
	return response
}

