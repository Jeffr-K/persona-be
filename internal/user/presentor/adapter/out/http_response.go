package out

type BusinessResponse struct {
	message string
	data    interface{}
}

func NewBusinessResponse(message string, data interface{}) *BusinessResponse {
	return &BusinessResponse{
		message: message,
		data:    data,
	}
}
