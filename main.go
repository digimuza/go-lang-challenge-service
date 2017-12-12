package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
)

func main() {
	iterations := 1000
	var challanege []byte

	//Generated random 30 byte seed
	seed, _ := GenerateRandomBytes(30)
	challanege = append(challanege, seed...)

	//Generate 2 byte soliution
	soliution, _ := GenerateRandomBytes(2)

	fmt.Println("Created soliution ", converToString(soliution))
	challanege = append(challanege, soliution...)

	h := createHash(challanege, iterations)
	fmt.Println(converToString(h))

	sol := SolveChallange(iterations, h, seed)

	for s := range sol {
		fmt.Println(hex.EncodeToString(s))
	}

}

// GenerateRandomBytes - function generate random bytes
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func createHash(b []byte, iterations int) []byte {
	for i := 0; i < iterations; i++ {
		b = GenerateHashFromBytes(b)
	}

	return b
}

//GenerateHashFromBytes function generate hash from bytes
func GenerateHashFromBytes(b []byte) []byte {
	h := sha256.New()
	h.Write(b)

	return h.Sum(nil)
}

func converToString(b []byte) string {
	return hex.EncodeToString(b)
}

func outputByte(b []byte) {
	fmt.Println(fmt.Sprintf("%x", b))
}

// SolveChallange Challange solving algorythm
func SolveChallange(iterations int, hash []byte, seed []byte) chan []byte {

	out := make(chan []byte)

	max := uint16(math.Pow(256, 0x2) - 1)
	next := uint16(0)
	go func() {
		for true {

			var challanege []byte
			challanege = append(challanege, seed...)

			buf := new(bytes.Buffer)
			binary.Write(buf, binary.BigEndian, next)

			soliution := buf.Bytes()
			challanege = append(challanege, soliution...)

			trying := createHash(challanege, iterations)

			if converToString(trying) == converToString(hash) {
				fmt.Println(fmt.Sprintln("Soliution", converToString(soliution)))
				out <- hash
			}

			if next == max {
				break
			}
			next++
		}
		close(out)
	}()

	return out
}
