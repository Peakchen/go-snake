// add by stefan

package RedisConn

import (
	"strings"

	"github.com/gomodule/redigo/redis"
)

func updateScript() (us *redis.Script) {
	us = redis.NewScript(2, `
	local key1 = KEYS[1] -- hash val
	local key2 = KEYS[2] -- key or hash field
	local ag1 = ARGV[1] -- field's value
	local ag2 = ARGV[2] -- value
	local ag3 = ARGV[3]	-- "ex" flag
	local ag4 = ARGV[4] -- exist time
	redis.call('HSET', key1, key2, ag1)
	if ag3 == "ex" then
		redis.call('SETEX', key2, ag4, ag2) 
	else
		redis.call('SET', key2, ag2)
	end
	`)
	return
}

/*
	key is 21 length string
	min key: "111111111111111111111" string to int32 val:1029
	max key: "fffffffffffffffffffff" string to int32 val:2142
	sub number : 1675
	key transfer interge %1000 with 1-1000
*/
func RoleKey2Haskey(key string) (hashk int) {
	for _, ki := range key {
		hashk += int(ki)
	}
	hashk = hashk % ERedHasKeyTransferMultiNum
	if hashk == 0 {
		hashk = 1
	}
	return
}

func MakeRedisModel(Identify, MainModel, SubModel string) string {
	return MainModel + "." + SubModel + "." + Identify
}

func ParseRedisKey(redkey string) (Identify, MainModel, SubModel string) {
	rediskeys := strings.Split(redkey, ".")
	MainModel = rediskeys[0]
	SubModel = rediskeys[1]
	Identify = rediskeys[2]
	return
}
