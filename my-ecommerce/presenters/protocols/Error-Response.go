package protocols

type ErrorResponse struct {
	Stack interface{} `json:"stack"`
	Message string `json:"message"`
}
