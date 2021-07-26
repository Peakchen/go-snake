package logic

//load table unified entry

import (
	"go-snake/world/mapConfig"
)

func LoadTab(){

	mapConfig.LoadMapTmxConfig("../table/map.tmx")
	

}