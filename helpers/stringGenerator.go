package helpers

import (
	"crypto/rand"
	"log"
	"math/big"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

const KeyLength = 10

func GetUUID() string {
	token := ""

	for i := 0; i < KeyLength; i++ {
		token += string(alphabet[cryptoRandSecure(int64(len(alphabet)))])
	}

	return token
}

func cryptoRandSecure(max int64) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		log.Println(err)
	}

	return nBig.Int64()
}
