package utils

// add by stefan

import (
	"math"
	"math/rand"

	"github.com/Peakchen/xgameCommon/aktime"
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

//----------------------------------------int-------------------------------------------
// 随机 [low, high]
func RandomIntBetween(low int, high int) int {
	if low == high {
		return low
	}
	if low > high || high == 0 {
		return 0
	}

	v := rand.Intn(high-low+1) + low
	if v > math.MaxInt32 {
		v = math.MaxInt32
	} else if v < math.MinInt32 {
		v = low
	}
	return v
}

// c# random.next [low, high)
func RandomNextInt(low, high int) int {
	return RandomIntBetween(low, high-1)
}

// c# random.next [0, high)
func RandomInt(high int) int {
	return RandomNextInt(0, high)
}

//----------------------------------------int32-------------------------------------------
// 随机 [low, high]
func RandomInt32Between(low int32, high int32) int32 {
	if low == high {
		return low
	}
	if low > high || high == 0 {
		return 0
	}

	v := rand.Int31n(high-low+1) + low
	if v > int32(math.MaxInt32) {
		v = int32(math.MaxInt32)
	} else if v < int32(math.MinInt32) {
		v = low
	}
	return v
}

// c# random.next [low, high)
func RandomNextInt32(low, high int32) int32 {
	return RandomInt32Between(low, high-1)
}

// c# random.next [0, high)
func RandomInt32(high int32) int32 {
	return RandomNextInt32(0, high)
}

//----------------------------------------uint32-------------------------------------------
// 随机 [low, high]
func RandomUInt32Between(low uint32, high uint32) uint32 {
	if low == high {
		return low
	}
	if low > high || high == 0 {
		return 0
	}

	v := uint32(rand.Int31n(int32(high-low+1))) + low
	if v > uint32(math.MaxInt32) {
		v = uint32(math.MaxInt32)
	}
	return v
}

// c# random.next [low, high)
func RandomNextUInt32(low, high uint32) uint32 {
	return RandomUInt32Between(low, high-1)
}

// c# random.next [0, high)
func RandomUInt32(high uint32) uint32 {
	return RandomNextUInt32(0, high)
}

//----------------------------------------int64-------------------------------------------
// 随机 [low, high]
func RandomInt64Between(low int64, high int64) int64 {
	if low == high {
		return low
	}
	if low > high || high == 0 {
		return 0
	}

	v := rand.Int63n(high-low+1) + low
	if v > int64(math.MaxInt64) {
		v = int64(math.MaxInt64)
	} else if v < int64(math.MinInt64) {
		v = low
	}
	return v
}

// c# random.next [low, high)
func RandomInt64Next2(low, high int64) int64 {
	return RandomInt64Between(low, high-1)
}

// c# random.next [0, high) int64
func RandomInt64(high int64) int64 {
	return RandomInt64Next2(0, high)
}

// 随机 [low, high]
func RandomUInt64Between(low uint64, high uint64) uint64 {
	if low == high {
		return low
	}
	if low > high || high == 0 {
		return 0
	}

	v := uint64(rand.Int63n(int64(high-low+1))) + low
	if v > uint64(math.MaxInt64) {
		v = uint64(math.MaxInt64)
	}
	return v
}

// c# random.next [low, high)
func RandomUInt64Next2(low, high uint64) uint64 {
	return RandomUInt64Between(low, high-1)
}

// c# random.next [0, high) int64
func RandomUInt64(high uint64) uint64 {
	return RandomUInt64Next2(0, high)
}

//----------------------------------String---------------------------------------------
func RandomLowerString(length int) string {
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		b := rand.Intn(26) + 97
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func RandomUpperString(length int) string {
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		b := rand.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
