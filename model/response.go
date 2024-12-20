package model

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Responses []Response

type MessageResponse struct {
	Data     any       `json:"data"`
	Errors   Responses `json:"errors"`
	Messages Responses `json:"messages"`
}
