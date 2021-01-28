//+build ignore

package main

import (
	"io/ioutil"
	"math"
	"math/rand"
)

const (
	Byte = 1
	KB   = 1024 * Byte
	MB   = 1024 * KB
	GB   = 1024 * MB
)

func main() {
	generate(10*MB, wide, "wide_10mb.txt")
	generate(512*MB, wide, "wide_512mb.txt")
	generate(10*MB, high, "high_10mb.txt")
	generate(512*MB, high, "high_512mb.txt")
	generate(10*MB, square, "square_10mb.txt")
	generate(512*MB, square, "square_512mb.txt")
}

func generate(size int, populate func([]byte), path string) {
	data := make([]byte, size)
	populate(data)
	check(ioutil.WriteFile(path, data, 0666))
}

func wide(b []byte) {
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
}

func high(b []byte) {
	for i := range b {
		if i%2 == 1 {
			b[i] = '\n'
		} else {
			b[i] = alphabet[rand.Intn(len(alphabet))]
		}
	}
}

func square(b []byte) {
	w := int(math.Sqrt(float64(len(b))))
	for i := range b {
		if i%w == w-1 {
			b[i] = '\n'
		} else {
			b[i] = alphabet[rand.Intn(len(alphabet))]
		}
	}
}

var alphabet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func check(err error) {
	if err != nil {
		panic(err)
	}
}
