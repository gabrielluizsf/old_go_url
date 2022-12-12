package analyzer

import (
	"fmt"
	"github.com/GabrielLuizSF/go_url/pkg/system"
)

func OSAnalyzer(){
	switch MyOS(){
		case "linux":
			fmt.Println(MyOS());
			system.LinuxOS();
		case "windows":
			fmt.Println(MyOS());
			system.WindowsOS();
		case "darwin":
			fmt.Println(MyOS());
			system.MACOS();
		default:
			fmt.Println("UNDEFINIED");
	}
	

	
}