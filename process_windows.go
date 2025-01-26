//go:build windows

package main

import (
	"os"
	"syscall"
	"unsafe"
)

const processQueryInformation = 0x0400
const processVmRead = 0x0010

// 获取父进程名称 (Windows)
func getParentProcessName() (string, error) {
	ppid := os.Getppid()
	handle, err := syscall.OpenProcess(processQueryInformation|processVmRead, false, uint32(ppid))
	if err != nil {
		return "", err
	}
	defer syscall.CloseHandle(handle)

	var modName [syscall.MAX_PATH]uint16
	size := uint32(len(modName))
	err = GetModuleBaseName(handle, 0, &modName[0], size)
	if err != nil {
		return "", err
	}

	procName := syscall.UTF16ToString(modName[:])
	if len(procName) > 4 && procName[len(procName)-4:] == ".exe" {
		procName = procName[:len(procName)-4]
	}
	return procName, nil
}

// GetModuleBaseName is a wrapper around the Windows API function that retrieves the name of the specified module
func GetModuleBaseName(handle syscall.Handle, hModule syscall.Handle, baseName *uint16, size uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(
		procGetModuleBaseName.Addr(),
		4,
		uintptr(handle),
		uintptr(hModule),
		uintptr(unsafe.Pointer(baseName)),
		uintptr(size),
		0,
		0,
	)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

var (
	modpsapi              = syscall.NewLazyDLL("psapi.dll")
	procGetModuleBaseName = modpsapi.NewProc("GetModuleBaseNameW")
)
