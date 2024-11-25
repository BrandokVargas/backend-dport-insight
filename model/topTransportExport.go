package model

type TopTransport struct {
	Transporte string `json:"transporte"`
	Total      int64  `json:"total"`
}

type TopsTransports []TopTransport
