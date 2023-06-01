package dto

type ErrorResponse struct {
	ResponseMeta
	Data   any `json:"data"`
	Errors any `json:"errors,omitempty"`
}

func DefaultErrorResponseWithMessage(msg string) ErrorResponse {
	return ErrorResponse{
		ResponseMeta: ResponseMeta{
			Success:      false,
			MessageTitle: "Oops, something went wrong.",
			Message:      msg,
			ResponseTime: "",
		},
		Data: nil,
	}
}
