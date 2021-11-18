package presenter

import "nirikshan-backend/pkg/utils"

//ResponseStruct defines the structure for error responses
type ResponseStruct struct {
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

//CreateSuccessResponse is a helper function that creates a ResponseStruct
func CreateSuccessResponse(data interface{}, message string) *ResponseStruct {
	return &ResponseStruct{
		Status:  true,
		Data:    data,
		Message: message,
	}
}

//CreateErrorResponse is a helper function that creates a ResponseStruct
func CreateErrorResponse(appError utils.AppError) *ResponseStruct {
	return &ResponseStruct{
		Status:  false,
		Data:    appError.Error(),
		Message: utils.ErrorDescriptions[appError],
	}
}
