package interf

import "testing"

func TestMe(t *testing.T) {
	me := NewMe("thing more, thind different.")
	me.Write()
	if IsInterface(me) {
		t.Log("me is implements I")
	} else {
		t.Log("me is not implements I")
	}
}
