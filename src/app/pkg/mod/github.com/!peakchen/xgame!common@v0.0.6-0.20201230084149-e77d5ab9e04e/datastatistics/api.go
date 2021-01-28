package datastatistics

import (
	"github.com/Peakchen/xgameCommon/akLog"
	"fmt"
	"net/http"
	"sort"
)

func IsExistStatisticsTitle(SrcTitles []string, dstTitle string) (err error) {
	titelsLen := len(SrcTitles)
	if sort.Search(titelsLen, func(i int) bool {
		return SrcTitles[i] == dstTitle
	}) == titelsLen {
		err = fmt.Errorf("can not search title, input title: %v.", dstTitle)
		return
	}
	err = nil
	return
}

func MonitorRoutine(ip string, port string) {
	err := http.ListenAndServe(ip+":"+port, nil)
	if err != nil {
		akLog.Error("ListenAndServe: ", err)
	}
}
