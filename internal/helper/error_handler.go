package errorhandler

type ErrorHandler struct {
	Error        bool   `json:"error"`
	ErrorMessage string `json:"message"`
}
