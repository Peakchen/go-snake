package mapConfig

import (
	"github.com/Peakchen/go-tmx/tmx"
	"os"
	"github.com/Peakchen/xgameCommon/akLog"
)

var layerGIDMap = map[int][]tmx.GID{}

func LoadMapTmxConfig(file string){

	r, err := os.Open(file)
	if err != nil {
		return
	}

	m, err := tmx.Read(r)
	if err != nil {
		return
	}

	for i, v := range m.Layers {

		layerGIDs, err := m.DecodeLayer(&v)
		if err != nil {
			akLog.Error("can not decode tmx ")
			return
		}
		
		layerGIDMap[i] = layerGIDs
	}

	
}

func GetDefaultMap()[]tmx.GID{
	return GetTmxMap(0)
}

func GetTmxMap(i int)[]tmx.GID{
	return layerGIDMap[i]
}