package busi

import (
	"testing"
)

func TestGenOrderId(t *testing.T) {
	orderId := GenOrderId(16545664533364543)
	t.Log(orderId)
}
