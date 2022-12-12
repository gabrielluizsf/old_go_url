package analyzer

import (
	"runtime"
)

func MyOS()string{
	return runtime.GOOS;
}
func MyARCH()string{
	return runtime.GOARCH;
}