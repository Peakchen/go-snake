package utils

// add by stefan

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

// String2Bytes convert string to []byte
func String2Bytes(val string) []byte {
	return []byte(val)
}

// String2Int convert string to int
func String2Int(val string) (int, error) {
	return strconv.Atoi(val)
}

// Int2String convert int to string
func Int2String(val int) string {
	return strconv.Itoa(val)
}

// String2Int64 convert string to int64
func String2Int64(val string) (int64, error) {
	return strconv.ParseInt(val, 10, 64)
}

// Int642String convert int64 to string
func Int642String(val int64) string {
	return strconv.FormatInt(val, 10)
}

// String2UInt64 convert string to uint64
func String2UInt64(val string) (uint64, error) {
	return strconv.ParseUint(val, 10, 64)
}

// UInt642String convert uint64 to string
func UInt642String(val uint64) string {
	return strconv.FormatUint(val, 10)
}

// NSToTime convert ns to time.Time
func NSToTime(ns int64) (time.Time, error) {
	if ns <= 0 {
		return time.Time{}, errors.New("ns is err")
	}
	bigNS := big.NewInt(ns)
	return time.Unix(ns/1e9, int64(bigNS.Mod(bigNS, big.NewInt(1e9)).Uint64())), nil
}

func String2Slicebyte(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func SliceByte2String(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{
		Data: bh.Data,
		Len:  bh.Len,
	}
	return *(*string)(unsafe.Pointer(&sh))
}

func SliceBytesLength(data []byte) int {
	dst := SliceByte2String(data)
	return len(dst)
}

/*
	string array data cover to string.
*/
func StrArray2Str(src []string) string {
	return strings.Replace(strings.Trim(fmt.Sprint(src), "[]"), " ", ",", -1)
}

func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
