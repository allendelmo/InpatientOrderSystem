package util

import "errors"

type OrderStatus int32

const (
	Pending            OrderStatus = 0
	ReadyToCollect     OrderStatus = 1
	Cancelled          OrderStatus = 2
	PartiallyDispensed OrderStatus = 3
)

func OrderStatusToString(status OrderStatus) (statusStr string, err error) {
	switch status {
	case Pending:
		statusStr = "PENDING"
	case ReadyToCollect:
		statusStr = "READY TO COLLECT"
	case Cancelled:
		statusStr = "CANCELLED"
	case PartiallyDispensed:
		statusStr = "PARTIALLY DISPENSED"
	default:
		err = errors.New("status id is not known")
	}

	return statusStr, err
}
