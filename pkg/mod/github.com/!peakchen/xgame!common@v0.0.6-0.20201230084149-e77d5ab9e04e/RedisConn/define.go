// add by stefan

package RedisConn

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

// pool Idl
const (
	IDle_Invalie = iota
	IDle_one
	IDle_two
	IDle_three
	IDle_four
	IDle_five
)

// idl timeout value
const (
	IDleTimeOut_invalid = iota
	//second
	IDleTimeOut_five_sec = 5 * time.Second
	IDleTimeOut_ten_sec  = 10 * time.Second
	//minute
	IDleTimeOut_one_min   = 60 * time.Second
	IDleTimeOut_two_min   = 120 * time.Second
	IDleTimeOut_three_min = 180 * time.Second
	IDleTimeOut_four_min  = 240 * time.Second
)

const (
	// common used MilliSecond
	MSec_one         = time.Millisecond
	MSec_ten         = 10 * time.Millisecond
	MSec_one_Hundred = 100 * time.Millisecond

	// common used Second
	Sec_five    = 5 * time.Second
	Sec_ten     = 10 * time.Second
	Sec_fifteen = 15 * time.Second
	Sec_twenty  = 20 * time.Second
	Sec_thirty  = 30 * time.Second
	Sec_fourty  = 40 * time.Second
	Sec_fifty   = 50 * time.Second
	Sec_sixty   = 60 * time.Second
)

type REDIS_INT32 int32

const (
	REDIS_SET_DEADLINE REDIS_INT32 = 600 //s

)

type TRedisScript struct {
	name   string
	script *redis.Script
}

// script name define...
const (
	ERedScript_Update string = "update"
)

//role key transfer redis hash key multi
const (
	ERedHasKeyTransferMultiNum = 1000
)
