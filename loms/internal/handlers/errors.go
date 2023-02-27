package handlers

import "errors"

var (
	ErrEmptyOrderID = errors.New("empty orderID")
	ErrEmptyUser    = errors.New("empty user")
)
