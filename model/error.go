package model

import "fmt"

type Error struct {
	Code    string
	Err     error
	Who     string
	Status  int
	Data    any
	Message string
	UserID  string
}

func NewError() Error {
	return Error{}
}

func (e *Error) Error() string {
	return fmt.Sprintf("Code: %s, Err: %v, Who: %s, Status: %d, Data: %v, UserID: %s",
		e.Code,
		e.Err,
		e.Who,
		e.Status,
		e.Data,
		e.UserID,
	)
}

func (e *Error) HasCode() bool {
	return e.Code != ""
}

func (e *Error) HasStatusHTTP() bool {
	return e.Status > 0
}

func (e *Error) HasData() bool {
	return e.Data != nil
}
