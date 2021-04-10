package akLog

// add by stefan

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"sync"
	"syscall"
	"time"

	"github.com/Peakchen/xgameCommon/aktime"
	"github.com/Peakchen/xgameCommon/public"
	"github.com/Peakchen/xgameCommon/tool"
	"github.com/Peakchen/xgameCommon/utils"
	"github.com/Shopify/sarama"
)

type LoadingContent struct {
	Content string
	logType string
}

type TAokoLog struct {
	filename      string
	filehandle    *os.File
	cancle        context.CancelFunc
	ctx           context.Context
	wg            sync.WaitGroup
	filesize      uint64
	logNum        uint64
	data          chan *LoadingContent
	FileNo        uint32
	consumeClient sarama.ConsumerGroup
}

const (
	EnAKLogFileMaxLimix = 500 * 1024 * 1024
	EnLogDataChanMax    = 1024
)

const (
	EnLogType_Info  string = "logInfo"
	EnLogType_Error string = "logError"
	EnLogType_Fail  string = "logFail"
	EnLogType_Debug string = "logDebug"
)

var (
	aokoLog    sync.Map //map[string]*TAokoLog
	brokerAddr []string
)

var exitchan = make(chan os.Signal, 1)

func init() {
	//aokoLog = map[string]*TAokoLog{}
}

func InitLogBroker(addr []string) {
	brokerAddr = addr
}

func (this *TAokoLog) createConsumer() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	// config.Consumer.Offsets.CommitInterval = 10
	// config.Consumer.Offsets.AutoCommit.Enable = true
	// config.Consumer.Offsets.AutoCommit.Interval = 2
	config.Version = sarama.V2_0_0_0
	client, err := sarama.NewConsumerGroup(brokerAddr, KAFKA_LOG_CONSUMER_GROUP, config)
	if err != nil {
		panic("create ConsumerGroup err: " + err.Error())
	}
	this.consumeClient = client
}

func checkNewLog(logtype string) (logobj *TAokoLog) {
	iter, ok := aokoLog.Load(logtype)
	if !ok {
		logobj = &TAokoLog{
			FileNo: 1,
		}
		if false == (brokerAddr == nil || len(brokerAddr) == 0) {
			logobj.createConsumer()
		}
		aokoLog.Store(logtype, logobj)
		initLogFile(logtype, logobj)
		go run(logobj)

		return logobj
	}
	return iter.(*TAokoLog)
}

func initLogFile(logtype string, log *TAokoLog) {
	var (
		RealFileName string
		PathDir      string = logtype
	)

	filename := utils.GetExeFileName()
	switch logtype {
	case EnLogType_Info:
		RealFileName = fmt.Sprintf("%v_Info_No%v_%v.log", filename, log.FileNo, aktime.Now().Local().Format(public.CstTimeDate))
	case EnLogType_Error:
		RealFileName = fmt.Sprintf("%v_Error_No%v_%v.log", filename, log.FileNo, aktime.Now().Local().Format(public.CstTimeDate))
	case EnLogType_Fail:
		RealFileName = fmt.Sprintf("%v_Fail_No%v_%v.log", filename, log.FileNo, aktime.Now().Local().Format(public.CstTimeDate))
	case EnLogType_Debug:
		RealFileName = fmt.Sprintf("%v_Debug_No%v_%v.log", filename, log.FileNo, aktime.Now().Local().Format(public.CstTimeDate))
	default:

	}

	exepath := utils.GetExeFilePath()
	filepath := exepath + "/" + PathDir
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		os.Mkdir(filepath, 0777)
		os.Chmod(filepath, 0777)
	}

	filehandler, err := os.OpenFile(filepath+"/"+RealFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic("open file fail, err: " + err.Error())
	}

	log.filehandle = filehandler
	log.filename = RealFileName
	log.data = make(chan *LoadingContent, EnLogDataChanMax)

}

func run(log *TAokoLog) {
	signal.Notify(exitchan, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGSEGV)
	log.ctx, log.cancle = context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go log.loop(&wg)
	if log.consumeClient != nil {
		wg.Add(1)
		go log.loop2(&wg)
	}
	wg.Wait()
	log.exit(&wg)
	time.Sleep(time.Duration(3) * time.Second)
}

func Error(args ...interface{}) {
	timeFormat := aktime.Now().Local().Format(public.CstTimeFmt)
	WriteLog(EnLogType_Error, "[Error]\t"+timeFormat, "", args)
}

func ErrorIDCard(identify string, args ...interface{}) {
	format := fmt.Sprintf("identify: %v.", identify)
	timeFormat := aktime.Now().Local().Format(public.CstTimeFmt)
	WriteLog(EnLogType_Error, "[IDCard]\t"+timeFormat, format, args)
}

func ErrorModule(data public.IDBCache, args ...interface{}) {
	format := fmt.Sprintf("main: %v, sub: %v, identify: %v, %v.", data.MainModel(), data.SubModel(), data.Identify(), args)
	timeFormat := aktime.Now().Local().Format(public.CstTimeFmt)
	WriteLog(EnLogType_Error, "[Module]\t"+timeFormat, format, args)
}

func Info(args ...interface{}) {
	timeFormat := aktime.Now().Local().Format(public.CstTimeFmt)
	WriteLog(EnLogType_Info, "[Info]\t"+timeFormat, "", args)
}

func Fail(args ...interface{}) {
	timeFormat := aktime.Now().Local().Format(public.CstTimeFmt)
	WriteLog(EnLogType_Fail, ("[Fail]\t" + timeFormat), "", args)
}

func Debug(format string, args ...interface{}) {
	WriteLog(EnLogType_Debug, "[Debug]\t", format, args)
}

func Panic() {
	log := checkNewLog(EnLogType_Fail)
	if log != nil {
		debug.PrintStack()
		buf := debug.Stack()
		log.filehandle.WriteString(string(buf[:]))
		log.endLog()
		//close(aokoLog.data)
	}
}

func WriteLog(logtype, title, format string, args []interface{}) {
	log := checkNewLog(logtype)
	if log == nil {
		Panic()
		return
	}

	var logStr string
	pc, _, line, ok := runtime.Caller(2)
	if ok {
		logStr += fmt.Sprintf(title+" "+"[%v:%v]", runtime.FuncForPC(pc).Name(), line)
	}
	logStr += fmt.Sprintf(" " + format)
	for i, data := range args {
		if i+1 <= len(args) {
			logStr += fmt.Sprintf("%v", data)
		}
		if i+1 < len(args) && i > 0 {
			logStr += ","
		}
	}
	logStr += "\n"
	if logtype == EnLogType_Fail {
		logStr += string(debug.Stack())
	}

	if log.filesize >= EnAKLogFileMaxLimix {
		FmtPrintf("log file: %v over max limix.", log.filename)
		log.FileNo++
		initLogFile(logtype, log)
		log.filesize = 0
	}

	log.filesize += uint64(len(logStr))
	log.logNum++

	log.data <- &LoadingContent{
		Content: logStr,
		logType: logtype,
	}

	switch logtype {
	case EnLogType_Info:
		logStr = fmt.Sprintf("\x1b[40m\x1b[%dm%s\x1b[0m", tool.LinuxForeground_YELLOW, logStr)
	case EnLogType_Error:
		logStr = fmt.Sprintf("\x1b[40m\x1b[%dm%s\x1b[0m", tool.LinuxForeground_Red, logStr)
	case EnLogType_Fail:
		logStr = fmt.Sprintf("\x1b[40m\x1b[%dm%s\x1b[0m", tool.LinuxForeground_WHITE, logStr)
	case EnLogType_Debug:
		logStr = fmt.Sprintf("\x1b[40m\x1b[%dm%s\x1b[0m", tool.LinuxBackground_GREEN, logStr)
	}

	fmt.Print(logStr)

	if log.logNum%EnLogDataChanMax == 0 {
		log.flush()
		log.data = make(chan *LoadingContent, EnLogDataChanMax)
	}
	aokoLog.Store(logtype, log)
}

func (this *TAokoLog) endLog() {
	if this.filehandle != nil {
		this.filehandle.Sync()
		this.filehandle.Close()
	}
}

func (this *TAokoLog) exit(wg *sync.WaitGroup) {
	fmt.Println("log exit: ", <-this.data, this.filesize, this.logNum)
	this.flush()
	this.endLog()
}

func (this *TAokoLog) loop(wg *sync.WaitGroup) {
	tick := time.NewTicker(time.Duration(10 * time.Second))
	defer func() {
		tick.Stop()
		wg.Done()
	}()

	for {
		select {
		case <-this.ctx.Done():
			return
		case log := <-this.data:
			this.writelog(log)
		case s := <-exitchan:
			fmt.Println("Got signal:", s)
			os.Exit(0)
			return
		case <-tick.C:
			this.flush()
		}
	}
}

func (this *TAokoLog) loop2(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	tick := time.NewTicker(time.Duration(5 * time.Second))
	for {
		select {
		case <-this.ctx.Done():
			tick.Stop()
			return
		case s := <-exitchan:
			fmt.Println("Got signal:", s)
			os.Exit(0)
			return
		case <-tick.C:
			if this.consumeClient != nil {
				err := this.consumeClient.Consume(this.ctx, []string{KAFKA_LOG_TOPIC}, this)
				if err != nil {
					fmt.Println("client.Consume error=", err.Error())
				}
			}
		}
	}
}

func (this *TAokoLog) writelog(src *LoadingContent) {
	_, err := this.filehandle.WriteString(src.Content)
	if err != nil {
		return
	}
	if src.logType == EnLogType_Fail {
		exitchan <- syscall.SIGKILL
	}
}

func (this *TAokoLog) flush() {
	this.writelog(<-this.data)
}

func (this *TAokoLog) Setup(s sarama.ConsumerGroupSession) error {
	return nil
}

func (this *TAokoLog) Cleanup(s sarama.ConsumerGroupSession) error {
	return nil
}

func (this *TAokoLog) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		key := string(message.Key)
		val := string(message.Value)
		this.writelog(&LoadingContent{
			Content: fmt.Sprintf("module: %v, info: %v.", key, val),
		})
		session.MarkMessage(message, "")
	}
	return nil
}
