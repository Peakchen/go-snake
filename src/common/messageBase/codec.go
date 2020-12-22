package messageBase

import "github.com/Peakchen/xgameCommon/utils"

var _codec utils.ICodec

func InitCodec(ic utils.ICodec) {
	_codec = ic
}

func Codec() utils.ICodec {
	return _codec
}
