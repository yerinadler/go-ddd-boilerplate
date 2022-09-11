package api

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func Ok(data interface{}, message string) *Response {
	return &Response{
		Status:  "000",
		Message: message,
		Data:    data,
	}
}

func Error(errorCode string, errorMessage interface{}, message string) *Response {
	return &Response{
		Status:  errorCode,
		Message: message,
		Error:   errorMessage,
	}
}
