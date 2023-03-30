package models

type ErrOrder struct {
	Order OutboxOrder
	Err   error
}
