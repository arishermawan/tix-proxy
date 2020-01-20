package view

type Response struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	Errors  []ErrorResponse    `json:"errors"`
}

type ErrorResponse struct {
	Code     string `json:"code"`
	Message  string `json:"message"`
	Title    string `json:"message_title"`
	Severity string `json:"message_severity"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}
