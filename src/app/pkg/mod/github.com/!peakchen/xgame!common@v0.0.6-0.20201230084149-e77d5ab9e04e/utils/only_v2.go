package utils

import (
	"fmt"

	"github.com/Peakchen/xgameCommon/aktime"
	"github.com/sony/sonyflake"
)

func getMachineID_linux() (uint16, error) {
	return 0, nil
}

func checkMachineID_linux(machineID uint16) bool {
	return false
}

func getMachineID_win() (uint16, error) {
	return 0, nil
}

func checkMachineID_win(machineID uint16) bool {
	return false
}

func NewOnly_v2() (id uint64, err error) {
	var (
		getFunc   func() (uint16, error)
		checkFunc func(uint16) bool
	)
	if IsWindows() {
		getFunc = getMachineID_win
		checkFunc = checkMachineID_win
	} else if IsLinux() {
		getFunc = getMachineID_linux
		checkFunc = checkMachineID_linux
	} else {
		//...
	}

	if getFunc == nil || checkFunc == nil {
		err = fmt.Errorf("new only v2 id fail.")
		panic(err)
		return
	}

	settings := sonyflake.Settings{
		StartTime:      aktime.Now(),
		MachineID:      getFunc,
		CheckMachineID: checkFunc,
	}

	sf := sonyflake.NewSonyflake(settings)
	id, err = sf.NextID()
	if err != nil {
		panic(err)
	}
	return
}

func NewInt64_v2() uint64 {
	id, _ := NewOnly_v2()
	return id
}

func NewString_v2() string {
	id, _ := NewOnly_v2()
	return Int642String(int64(id))
}
