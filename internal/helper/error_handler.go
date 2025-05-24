package errorhandler

type ErrorHandler struct {
	Error   bool   `json:"error" default:"true"`
	Message string `json:"message"`
}
