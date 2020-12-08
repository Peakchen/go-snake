package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"strconv"
	"strings"
)

//create 32bit md5 string
func GetMd5String(s string, upper bool, half bool) string {
	h := md5.New()
	h.Write([]byte(s))
	result := hex.EncodeToString(h.Sum(nil))
	if upper == true {
		result = strings.ToUpper(result)
	}
	if half == true {
		result = result[8:24]
	}
	return result
}

//use rand number create Guid string
func NewOnly_v3() string {
	b := make([]byte, 48)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b), true, false)
}

func NewInt64_v3() (id int64) {
	sid := NewOnly_v3()
	if len(sid) == 0 {
		return
	}

	val, err := strconv.Atoi(sid)
	if err != nil {
		panic(err)
		return
	}

	id = int64(val)
	return
}

func NewString_v3() string {
	return NewOnly_v3()
}
