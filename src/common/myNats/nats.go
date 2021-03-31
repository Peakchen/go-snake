package myNats

import (
	"github.com/Peakchen/xgameCommon/utils"
)

func Register(addr string, codec utils.CodecType){

	SetCodec(utils.GetCodecByType(codec))

	CreateMyNats(addr)

}