package models

type SendStatus int

const (
	UnknownStatus SendStatus = iota
	Open
	InProgress
	Closed
)
