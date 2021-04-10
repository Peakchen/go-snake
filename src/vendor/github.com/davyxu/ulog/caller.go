package ulog

import (
	"runtime"
	"strings"
	"sync"
)

var (

	// qualified package name, cached at first use
	packageName string

	// Positions in the call stack when tracing to report the calling method
	minimumCallerDepth int

	// Used for caller information initialisation
	callerInitOnce sync.Once
)

// getPackageName reduces a fully qualified function name to the package name
// There really ought to be to be a better way...
func getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
}

const (
	maximumCallerDepth int = 25
	knownLogrusFrames  int = 4
)

// getCaller retrieves the name of the first non-logrus calling function
func getCaller() *runtime.Frame {

	// cache this package's fully-qualified name
	callerInitOnce.Do(func() {
		pcs := make([]uintptr, 2)
		_ = runtime.Callers(0, pcs)
		packageName = getPackageName(runtime.FuncForPC(pcs[1]).Name())

		// now that we have the cache, we can skip a minimum count of known-logrus functions
		// XXX this is dubious, the number of frames may vary
		minimumCallerDepth = knownLogrusFrames
	})

	// Restrict the lookback frames to avoid runaway lookups
	pcs := make([]uintptr, maximumCallerDepth)
	depth := runtime.Callers(minimumCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	for f, again := frames.Next(); again; f, again = frames.Next() {
		pkg := getPackageName(f.Function)

		// If the caller isn't part of this package, we're done
		if pkg != packageName {
			return &f
		}
	}

	// if we got here, we failed to find the caller's context
	return nil
}

// 将一个文件路径, 按照指定的级别缩短
// 例如: 输入C:/Go/src/testing/testing.go
// 当level=0时, 返回testing.go
// 当level=1时, 返回testing/testing.go
// 当level=2时, 返回src/testing/testing.go

func trimPath(filename string, level int) string {

	if level == -1 {
		return filename
	}

	var hitTimes int
	for i := len(filename) - 1; i >= 0; i-- {

		c := filename[i]
		if c == '/' {
			if hitTimes >= level {
				return filename[i+1:]
			}

			hitTimes++
		}
	}

	return filename
}
