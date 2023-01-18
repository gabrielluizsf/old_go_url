package analyzer

import (
	"runtime"
)
//this function returns the operating system name
func MyOS()string{
	return runtime.GOOS;
}
//this function returns the architecture of the system
func MyARCH()string{
	return runtime.GOARCH;
}
