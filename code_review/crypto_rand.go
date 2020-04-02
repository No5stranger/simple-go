package crypto_rand

import (
	"crypto/rand"
	"encoding/hex"
	//"fmt"
)

func Key() string {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}
	//return fmt.Sprintf("%x", buf)
	return hex.EncodeToString(buf)
}
