package models

type OrderStatusID int

const (
	UnknownOrderStatusID OrderStatusID = iota
	New
	AwaitingPayment
	Failed
	Payed
	Cancelled
)

func GetName(id OrderStatusID) string {
	switch id {
	case UnknownOrderStatusID:
		return "Unknown order status id"
	case New:
		return "New"
	case AwaitingPayment:
		return "Awaiting payment"
	case Failed:
		return "Failed"
	case Payed:
		return "Payed"
	case Cancelled:
		return "Cancelled"
	}
	return ""
}
