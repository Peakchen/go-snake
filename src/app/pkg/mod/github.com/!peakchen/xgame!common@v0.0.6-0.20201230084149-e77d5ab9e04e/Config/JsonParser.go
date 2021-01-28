package Config

import (
	"github.com/Peakchen/xgameCommon/akLog"
	"encoding/json"
	"io/ioutil"
)

/*
	json parser for config load.
	by stefan 20191108 v2.0
*/

var (
	_JsonParseTool *TJsonParseTool
)

type TJsonParseTool struct {
}

func (this *TJsonParseTool) Parse(jsonName string, data interface{}) (err error) {
	filedata, err := ioutil.ReadFile(jsonName)
	if err != nil {
		return
	}

	err = json.Unmarshal(filedata, data)
	return
}

func NewJsonParseTool() *TJsonParseTool {
	return &TJsonParseTool{}
}

func init() {
	_JsonParseTool = NewJsonParseTool()
}

/*
	purpose: parse json to related config module.
	param 1: config module.
	param 2: config read list.
	param 3: file name.
*/
func ParseJson2Cache(obj ICommonConfig, data interface{}, filename string) {
	if data == nil {
		akLog.Error("config data is nil, filename: ", filename)
		return
	}

	err := _JsonParseTool.Parse(filename, data)
	if err != nil {
		akLog.Error("[Parse json fail] err: ", err)
		return
	}

	cfg := &TConfig{
		data: data,
		obj:  obj,
	}

	errlist := cfg.Before()
	if errlist != nil && len(errlist) > 0 {
		for _, err := range errlist {
			akLog.Error("[config Before] err: ", err)
		}
		return
	}

	errlist = cfg.After()
	if errlist != nil && len(errlist) > 0 {
		for _, err := range errlist {
			akLog.Error("[config After] err: ", err)
		}
		return
	}

	return
}
