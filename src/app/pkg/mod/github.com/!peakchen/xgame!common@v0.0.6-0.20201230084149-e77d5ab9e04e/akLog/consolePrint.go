package akLog

// add by stefan

import (
	"fmt"

	"github.com/Peakchen/xgameCommon/aktime"
	"github.com/Peakchen/xgameCommon/public"
)

func FmtPrintf(src string, params ...interface{}) {
	var dst string
	if len(params) == 0 {
		dst = fmt.Sprintf(aktime.Now().Local().Format(public.CstTimeFmt) + " " + src)
	} else {
		dst = fmt.Sprintf(aktime.Now().Local().Format(public.CstTimeFmt)+" "+src, params...)
	}

	fmt.Println(dst)
}

func FmtPrintln(params ...interface{}) {
	content := make([]interface{}, 0, len(params)+1)
	content = append(content, aktime.Now().Format(public.CstTimeFmt))
	if len(params) > 0 {
		content = append(content, params...)
	}
	fmt.Println(content...)
}

func RetError(context string, params ...interface{}) error {
	return fmt.Errorf(context, params...)
}
