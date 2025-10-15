package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"os"
	"strings"
)

type Err struct {
	str string
}

func (e Err) Error() string {
	return e.str
}
func Binary(array []int, target int, lowIndex int, hightIndex int) (int, error) {
	if lowIndex > hightIndex || len(array) == 0 {
		return -1, &Err{str: "not dound"}
	}
	mid := int(lowIndex + (hightIndex-lowIndex)/2)
	if array[mid] > target {
		return Binary(array, target, lowIndex, mid-1)
	} else {
		return array[mid], nil
	}
}

func SequenceGrayCode(n uint) []uint {
	result := make([]uint, 0)
	var i uint
	for i = 0; i < 1<<n; i++ {
		result = append(result, i^(i>>1))
	}
	return result
}

func main() {
	var str strings.Builder
	f, _ := os.OpenFile("a.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	f.WriteString()

	sdd := "我"

	fmt.Printf("%+b\n", []byte(sdd))

	sddr := "1"

	fmt.Printf("%08b", []byte(sddr))
}

func Generate(minLength int, maxLength int) string {
	//var chars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+,.?/:;{}[]`~")

	length, err := rand.Int(rand.Reader, big.NewInt(int64(maxLength-minLength)))
	fmt.Println(length)
	if err != nil {
		panic(err) // handle this gracefully
	}
	length.Add(length, big.NewInt(int64(minLength)))

	intLength := int(length.Int64())

	fmt.Println(intLength)

	//newPassword := make([]byte, intLength)
	randomData := make([]byte, intLength+intLength/4)
	io.ReadFull(rand.Reader, randomData)
	fmt.Println(randomData)
	fmt.Println(string(randomData))
	// charLen := byte(len(chars))
	// maxrb := byte(256 - (256 % len(chars)))
	// i := 0
	// for {
	// 	if _, err := io.ReadFull(rand.Reader, randomData); err != nil {
	// 		panic(err)
	// 	}
	// 	for _, c := range randomData {
	// 		if c >= maxrb {
	// 			continue
	// 		}
	// 		newPassword[i] = chars[c%charLen]
	// 		i++
	// 		if i == intLength {
	// 			return string(newPassword)
	// 		}
	// 	}
	// }
	return "ww"
}
