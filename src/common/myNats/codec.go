package myNats

import (
	"github.com/Peakchen/xgameCommon/utils"
)

var (
	natsCodec utils.ICodec		
)

func SetCodec(codec utils.ICodec){
	natsCodec = codec
}

func Codec()utils.ICodec{
	return natsCodec
}