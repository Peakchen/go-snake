package Cache

/*
	second level cache
	purpose: templorary cache data for quick operation
	add by stefan
*/

import "container/list"

type TCacheMgr struct {
	cache *TCache
}

var (
	_cmr = &TCacheMgr{}
)

func Init() {
	_cmr = &TCacheMgr{
		cache: &TCache{
			td: ConstCacheOverTime,
			cl: list.New(),
		},
	}

	_cmr.run()
}

func (this *TCacheMgr) run() {
	_cmr.cache.Run()
}

func GetTempData(identify string) (data interface{}) {
	data = _cmr.cache.Get(identify)
	return
}

func SetTempData(identify string, data interface{}) {
	_cmr.cache.Set(identify, data)
}
