package errorhandler

type ErrorHandler struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}
