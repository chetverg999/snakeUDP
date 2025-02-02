package udp

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func Secretkey() int64 {
	safeNum, err := rand.Int(rand.Reader, big.NewInt(9999)) // Генерация числа от 0 до 100233
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return safeNum.Int64()
}
