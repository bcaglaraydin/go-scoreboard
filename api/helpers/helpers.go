package helpers

func ResponseError(status int, msg string) HTTPError {
	er := HTTPError{
		Code:    status,
		Message: msg,
	}
	return er
}

type HTTPError struct {
	Code    int    `json:"status"`
	Message string `json:"message" `
}
