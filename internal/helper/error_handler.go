package responsehandler

type ErrorHandler struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func NewErrorHandler(message string) ErrorHandler {
	return ErrorHandler{
		Error:   true,
		Message: message,
	}
}
