package responsehandler

type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessWithData(data interface{}, message string) SuccessResponse {
	return SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func SuccessMessage(message string) SuccessResponse {
	return SuccessResponse{
		Success: true,
		Message: message,
	}
}
