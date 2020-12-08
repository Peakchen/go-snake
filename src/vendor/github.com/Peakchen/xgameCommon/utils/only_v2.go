package utils

import (
	"fmt"
	"strconv"

	"github.com/Peakchen/xgameCommon/aktime"
	"github.com/sony/sonyflake"
)

func getMachineID_linux() (uint16, error) {
	return 0, nil
}

func checkMachineID_linux(machineID uint16) bool {
	return machineID != 0
}

func getMachineID_win() (uint16, error) {
	val, err := strconv.Atoi(GetPhysicalID())
	if err != nil {
		return 0, err
	}
	return uint16(val), nil
}

func checkMachineID_win(machineID uint16) bool {
	return machineID != 0
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
