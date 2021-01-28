package utils

import (
	"github.com/rs/xid"
)

//github.com/rs/xid:           bva855ahtor3pi2tuq0g
func GetOnlyString_v6() string {
	return xid.New().String()
}
