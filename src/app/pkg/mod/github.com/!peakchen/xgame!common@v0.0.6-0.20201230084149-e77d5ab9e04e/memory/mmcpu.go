// add by stefan
package memory

import (
	"log"
	"os"
	"os/signal"
	"runtime/pprof"
	//"syscall"
)

const (
	LP_GOROUTINE = 1
)

var (
	m_mapRunInfo = map[int32]string{
		LP_GOROUTINE: "goroutine",
	}
)

func LoopWaitforSignal() {
	for {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)

	}
}

func LookupInfo(id int32) {
	var lstr, ok = m_mapRunInfo[id]
	if !ok {
		log.Printf("look up err id: %d.", id)
		return
	}

	pprof.Lookup(lstr).WriteTo(os.Stdout, 2)
}
