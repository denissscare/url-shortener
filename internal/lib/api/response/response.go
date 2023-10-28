package response

type Response struct {
	Status string `json:"status"` // error, ok
	Error  string `json:"error,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func OK() Response {
	return Response{
		Status: StatusOK,
	}
}
func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}
