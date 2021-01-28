// +build windows

package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/StackExchange/wmi"
	"golang.org/x/sys/windows"
	"net"
	"sort"
	"strings"
	"unsafe"
)

type cpuInfo struct {
	CPU        int32  `json:"cpu"`
	VendorID   string `json:"vendorId"`
	PhysicalID string `json:"physicalId"`
}

type win32_Processor struct {
	Manufacturer string
	ProcessorID  *string
}

func GetPhysicalID() string {
	var ids []string
	if guid, err := getMachineGuid(); err != nil {
		panic(err.Error())
	} else {
		ids = append(ids, guid)
	}
	if cpuinfo, err := getCPUInfo(); err != nil && len(cpuinfo) > 0 {
		panic(err.Error())
	} else {
		ids = append(ids, cpuinfo[0].VendorID+cpuinfo[0].PhysicalID)
	}
	if mac, err := getMACAddress(); err != nil {
		panic(err.Error())
	} else {
		ids = append(ids, mac)
	}
	sort.Strings(ids)
	idsstr := strings.Join(ids, "|/|")
	return GetMd5String(idsstr, true, true)
}

func getMACAddress() (string, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		panic(err.Error())
	}
	mac, macerr := "", errors.New("无法获取到正确的MAC地址")
	for i := 0; i < len(netInterfaces); i++ {
		//fmt.Println(netInterfaces[i])
		if (netInterfaces[i].Flags&net.FlagUp) != 0 && (netInterfaces[i].Flags&net.FlagLoopback) == 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				ipnet, ok := address.(*net.IPNet)
				//fmt.Println(ipnet.IP)
				if ok && ipnet.IP.IsGlobalUnicast() {
					// 如果IP是全局单拨地址，则返回MAC地址
					mac = netInterfaces[i].HardwareAddr.String()
					return mac, nil
				}
			}
		}
	}
	return mac, macerr
}

func getCPUInfo() ([]cpuInfo, error) {
	var ret []cpuInfo
	var dst []win32_Processor
	q := wmi.CreateQuery(&dst, "")
	fmt.Println(q)
	if err := wmiQuery(q, &dst); err != nil {
		return ret, err
	}

	var procID string
	for i, l := range dst {
		procID = ""
		if l.ProcessorID != nil {
			procID = *l.ProcessorID
		}

		cpu := cpuInfo{
			CPU:        int32(i),
			VendorID:   l.Manufacturer,
			PhysicalID: procID,
		}
		ret = append(ret, cpu)
	}

	return ret, nil
}

// WMIQueryWithContext - wraps wmi.Query with a timed-out context to avoid hanging
func wmiQuery(query string, dst interface{}, connectServerArgs ...interface{}) error {
	ctx := context.Background()
	if _, ok := ctx.Deadline(); !ok {
		ctxTimeout, cancel := context.WithTimeout(ctx, 3000000000) //超时时间3s
		defer cancel()
		ctx = ctxTimeout
	}

	errChan := make(chan error, 1)
	go func() {
		errChan <- wmi.Query(query, dst, connectServerArgs...)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		return err
	}
}

func getMachineGuid() (string, error) {
	// there has been reports of issues on 32bit using golang.org/x/sys/windows/registry, see https://github.com/shirou/gopsutil/pull/312#issuecomment-277422612
	// for rationale of using windows.RegOpenKeyEx/RegQueryValueEx instead of registry.OpenKey/GetStringValue
	var h windows.Handle
	err := windows.RegOpenKeyEx(windows.HKEY_LOCAL_MACHINE, windows.StringToUTF16Ptr(`SOFTWARE\Microsoft\Cryptography`), 0, windows.KEY_READ|windows.KEY_WOW64_64KEY, &h)
	if err != nil {
		return "", err
	}
	defer windows.RegCloseKey(h)

	const windowsRegBufLen = 74 // len(`{`) + len(`abcdefgh-1234-456789012-123345456671` * 2) + len(`}`) // 2 == bytes/UTF16
	const uuidLen = 36

	var regBuf [windowsRegBufLen]uint16
	bufLen := uint32(windowsRegBufLen)
	var valType uint32
	err = windows.RegQueryValueEx(h, windows.StringToUTF16Ptr(`MachineGuid`), nil, &valType, (*byte)(unsafe.Pointer(&regBuf[0])), &bufLen)
	if err != nil {
		return "", err
	}

	hostID := windows.UTF16ToString(regBuf[:])
	hostIDLen := len(hostID)
	if hostIDLen != uuidLen {
		return "", fmt.Errorf("HostID incorrect: %q\n", hostID)
	}

	return hostID, nil
}
