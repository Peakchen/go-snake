package myNats

import (
	"github.com/Peakchen/xgameCommon/utils"
)

func Register(addr string){

	SetCodec(utils.GetCodecByType(utils.ENCodecType_Gob))

	CreateMyNats(addr)

}