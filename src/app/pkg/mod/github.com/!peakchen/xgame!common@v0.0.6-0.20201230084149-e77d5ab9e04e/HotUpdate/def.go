package HotUpdate

import (
	"os"
)

type TServerHotUpdateInfo struct {
	HUCallback func()
	Recvsignal os.Signal
}
