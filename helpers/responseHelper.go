package helpers

type HTTPSuccessResponse struct {
	Data interface{} `json:"data"`
}

type CodeMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type HTTPErrorResponse struct {
	Error CodeMessage `json:"error"`
}

func SuccessResponse(data interface{}) HTTPSuccessResponse {
	return HTTPSuccessResponse{Data: data}
}

func ErrorResponse(code int, message string) HTTPErrorResponse {
	r := CodeMessage{
		Code:    code,
		Message: message,
	}
	return HTTPErrorResponse{Error: r}
}

func MResponse(code int, message string) CodeMessage {
	return CodeMessage{
		Code:    code,
		Message: message,
	}
}

