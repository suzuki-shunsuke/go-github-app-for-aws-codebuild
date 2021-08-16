package ghapputil

type Response struct {
	StatusCode int         `json:"statusCode"`
	Body       interface{} `json:"body"`
}

type ResponseBody struct {
	Message string `json:"message"`
}

func NewResponse(code int, message string) *Response {
	return &Response{
		StatusCode: code,
		Body: &ResponseBody{
			Message: message,
		},
	}
}
