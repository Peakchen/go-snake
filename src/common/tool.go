package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"time"
	//"sync"

	"github.com/Peakchen/xgameCommon/akLog"

)

func Dosafe(fn func(), exitfn func()) {
	defer ExceptionStack(exitfn)
	fn()
}

//var wg sync.WaitGroup
	
func DosafeRoutine(fn func(), exitfn func()) {

	//wg.Add(1)
	
	//go func(){

	    defer func(){

			err := recover()
			if err != nil {
				
				if exitfn != nil {
					exitfn()
				}

				errstr := fmt.Sprintf("\n%s runtime error: %v\n traceback:\n", separator, err)
				errstr += callerDebug()
				errstr += separator + "\n"
				akLog.Error(errstr, string(debug.Stack()))
			}

	        //wg.Done()

	    }()
	    
	    go fn()
	    
	//}()
	
	//wg.Wait()
}

func callerDebug() (str string) {
	str += "\n"
	i := 2
	for {
		pc, file, line, ok := runtime.Caller(i)
		if !ok || i > 10 {
			break
		}
		str += fmt.Sprintf("\t stack: %d %v [file: %s] [func: %s] [line: %d]\n", i-1, ok, file, runtime.FuncForPC(pc).Name(), line)
		i++
	}
	return str + "\n"
}

const (
	separator = "==========================="
)

func ExceptionStack(fn func()) {
	err := recover()
	if err != nil {
		if fn != nil {
			fn()
		}
		errstr := fmt.Sprintf("\n%s runtime error: %v\n traceback:\n", separator, err)
		errstr += callerDebug()
		errstr += separator + "\n"
		akLog.Error(errstr, string(debug.Stack()))
	}
}

func SafeExit() {
	akLog.Error("\nunknow exception, exit: \n", separator, callerDebug(), string(debug.Stack()), separator+"\n")
	time.Sleep(time.Second)
	os.Exit(0)
}

func HmacSha256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func ComputeHmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	//	fmt.Println(h.Sum(nil))
	sha := hex.EncodeToString(h.Sum(nil))
	//	fmt.Println(sha)

	//	hex.EncodeToString(h.Sum(nil))
	return base64.StdEncoding.EncodeToString([]byte(sha))
}

func ParseJson(src string, dst interface{}, desc string){

	if len(src) == 0 {
		panic(errors.New("json content is empty."))
	}

	err := json.Unmarshal([]byte(src), dst)
	if err != nil {
	
		panic(fmt.Errorf("json parse err: %v, desc: %v.", err, desc))

	}

}