package cryptopals

import (
	"encoding/hex"
	"encoding/base64"
	"log"
)

func HexToBase64(s string) string {
	dec, err := hex.DecodeString(s)
	if (err != nil) {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString([]byte (dec))
}
