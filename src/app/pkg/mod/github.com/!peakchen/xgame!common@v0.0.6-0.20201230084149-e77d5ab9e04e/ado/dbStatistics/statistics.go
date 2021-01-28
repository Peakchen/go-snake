package dbStatistics

/*
	purpose: statistics db write count for check db operate frequency.
	date: 20200114 14:30
*/

import (
	"bytes"
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/aktime"
	"github.com/Peakchen/xgameCommon/public"
	"github.com/Peakchen/xgameCommon/stacktrace"
	"github.com/Peakchen/xgameCommon/utils"
	"encoding/gob"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type TModelStatistics struct {
	strTime  string
	stacklog bytes.Buffer
}

type TDBStatistics struct {
	filehandle     *os.File
	chbuff         chan bytes.Buffer
	userOperModels sync.Map //key: identify, value: map[string][]*TModelStatistics
}

const (
	cstSaveLogTickers = 60 //s
	cstStatisticsLog  = "DBStatisticsLog"
)

var (
	_dbStatistics *TDBStatistics
)

func InitDBStatistics() {
	_dbStatistics = &TDBStatistics{
		filehandle: nil,
		chbuff:     make(chan bytes.Buffer),
	}
	_dbStatistics.Init()
	go _dbStatistics.loop()
}

func (this *TDBStatistics) Init() {
	exename := utils.GetExeFileName()
	_, err := os.Stat(cstStatisticsLog)
	if err != nil {
		akLog.Error("err: ", err)
		return
	}

	if os.IsNotExist(err) {
		err := os.Mkdir("DBStatisticsLog", 0644)
		if err != nil {
			akLog.Error("err: ", err)
			return
		}
	}

	fileName := fmt.Sprintf("./DBStatisticsLog/%v_DBStatistics_%v.log", exename, aktime.Now().Local().Format(public.CstTimeDate))
	filehandle, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		akLog.Error("can not create db statistics log file, err：", err)
		return
	}

	this.filehandle = filehandle
}

func (this *TDBStatistics) loop() {
	for {
		select {
		case data := <-this.chbuff:
			this.filehandle.WriteString(data.String())
		}
	}
}

func (this *TDBStatistics) exit() {
	this.filehandle.Sync()
	this.filehandle.Close()
}

func (this *TDBStatistics) Update(content string) {
	var buff bytes.Buffer
	if err := gob.NewEncoder(&buff).Encode(content); err != nil {
		akLog.Error("cover to buffer fail, err: ", err)
		return
	}

	this.chbuff <- buff
}

func (this *TDBStatistics) loadOrInitUser(identify string) (modeldata map[string][]*TModelStatistics) {
	modeldata = nil
	if len(identify) == 0 {
		return
	}
	value, _ := this.userOperModels.LoadOrStore(identify, make(map[string][]*TModelStatistics, 0))
	if value == nil {
		akLog.Error("cache model invalid.")
		return
	}

	var ok bool
	modeldata, ok = value.(map[string][]*TModelStatistics)
	if !ok {
		akLog.Error("cache model invalid data type.")
		return
	}
	return
}

func (this *TDBStatistics) deleteUser(identify string) {
	this.userOperModels.Delete(identify)
}

/*
	statistics msg logic runs with db operations.
*/
func DBOperStatistics(identify, model string) {
	if _dbStatistics == nil {
		InitDBStatistics()
	}
	modeldata := _dbStatistics.loadOrInitUser(identify)
	if modeldata == nil {
		return
	}

	var buff bytes.Buffer
	buff.WriteString(stacktrace.NormalStackLog())
	modelStatistics, ok := modeldata[model]
	if !ok {
		modelStatistics = []*TModelStatistics{}
	}

	modelStatistics = append(modelStatistics, &TModelStatistics{
		strTime:  aktime.Now().Local().Format(public.CstTimeFmt),
		stacklog: buff,
	})
	modeldata[model] = modelStatistics
}

/*
	statistics msg logic ends msg mainid and subid.
*/
func DBMsgStatistics(identify string, mainid, subid uint16) {
	if _dbStatistics == nil {
		InitDBStatistics()
	}
	modeldata := _dbStatistics.loadOrInitUser(identify)
	if modeldata == nil {
		return
	}

	var buff bytes.Buffer
	buff.WriteString("identify: " + identify + "\r\n")
	buff.WriteString("mainid: " + strconv.Itoa(int(mainid)) + "\r\n")
	buff.WriteString("subid: " + strconv.Itoa(int(subid)) + "\r\n")
	for model, statistics := range modeldata {
		buff.WriteString("model oper cnt: " + strconv.Itoa(len(statistics)) + "\r\n")
		buff.WriteString("model name: " + model + "\r\n")
		for _, log := range statistics {
			buff.WriteString("time: " + log.strTime + "\r\n")
			buff.WriteString("stack log: \r\n" + log.stacklog.String() + "\r\n")
		}
	}

	_dbStatistics.Update(buff.String())
	_dbStatistics.deleteUser(identify)
}

/*
	stop msg log statistics.
*/
func DBStatisticsStop() {
	_dbStatistics.exit()
}
