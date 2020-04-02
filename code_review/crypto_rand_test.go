package crypto_rand

import (
	"testing"
)

func TestKey(t *testing.T) {
	key := Key()
	t.Log(key)
}
