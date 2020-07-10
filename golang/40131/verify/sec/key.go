package sec

import (
	"encoding/hex"
	"fmt"
)

var CryptoKey []byte

func InitWithKey(key []byte) {
	CryptoKey = key
	fmt.Println(hex.EncodeToString(key))
}