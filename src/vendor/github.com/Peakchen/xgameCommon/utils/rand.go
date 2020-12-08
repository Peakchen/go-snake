package utils

// add by stefan

import (
	"github.com/Peakchen/xgameCommon/aktime"
	"math/rand"
)

/*
	rand api
*/

func init() {
	t := aktime.Now().Unix()
	s := rand.NewSource(t)
	rand.New(s).Seed(t)
}

/*
	rand int32 number
	space: [0,n)
	min: 1
	max: n
*/
func RandInt32FromZero(n int32) (result int32) {
	result = rand.Int31n(n)
	return
}

/*
	rand int32 number
	space: [1,n]
	min: 1
	max: n
*/
func RandInt32(n int32) (result int32) {
	n = n + 1
	result = rand.Int31n(n)
	if result == 0 {
		result = 1
	}
	return
}

/*
	rand int64 number
	space: [1,n]
	min: 1
	max: n
*/
func RandInt64(n int64) (result int64) {
	n = n + 1
	result = rand.Int63n(n)
	if result == 0 {
		result = 1
	}
	return
}

/*
	create rand unsign 64 bit by single number
*/
func RandByUInt64(n int64) (result int64) {
	s := rand.NewSource(n)
	result = rand.New(s).Int63()
	return
}

/*
	create rand unsign 32 bit by single number
*/
func RandByUInt32(n int64) (result uint32) {
	s := rand.NewSource(n)
	result = rand.New(s).Uint32()
	return
}
