package cryptopals

import (
	"encoding/hex"
	"encoding/base64"
	"log"
	"os"
)

func HexToBase64(s string) string {
	dec, err := hex.DecodeString(s)
	if (err != nil) {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString([]byte (dec))
}

func FixedXOR(buff1 string, buff2 string) string {
	if len(buff1) == 0 || len(buff2) == 0 {
		os.Exit(1)
	}
	decBuff1, err := hex.DecodeString(buff1)
	if err != nil {
		log.Fatal(err)
	}
	decBuff2, err2 := hex.DecodeString(buff2)
	if err2 != nil {
		log.Fatal(err2)
	}
	
	a := len(decBuff1)
	b := make([]byte, a)
	for i := 0; i < a; i++ {
		b[i] = decBuff1[i] ^ decBuff2[i]
	}

	return hex.EncodeToString(b)
}
