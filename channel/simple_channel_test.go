package simple_channel

import "testing"

func TestChannel(t *testing.T) {
	SendReceive()
}

func TestRangeChannel(t *testing.T) {
	RangeChannel()
}
