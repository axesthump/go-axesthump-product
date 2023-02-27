package handlers

import "errors"

var (
	ErrEmptyUser = errors.New("empty user")
	ErrEmptySku  = errors.New("empty sku")
)
