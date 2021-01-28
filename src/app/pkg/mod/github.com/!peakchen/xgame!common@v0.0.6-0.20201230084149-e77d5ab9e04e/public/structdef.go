package public

import "github.com/Peakchen/xgameCommon/ado"

/*

 */
type IDBCache interface {
	Identify() string
	MainModel() string
	SubModel() string
}

type TCommonRedisCache struct {
	ado.IDBModule
}

func (this *TCommonRedisCache) Identify() string {
	return ""
}

func (this *TCommonRedisCache) MainModel() string {
	return ""
}

func (this *TCommonRedisCache) SubModel() string {
	return ""
}

func (this *TCommonRedisCache) Version() string {
	return "1.0"
}

type UpdateDBCacheCallBack func(string, string, []byte) bool
