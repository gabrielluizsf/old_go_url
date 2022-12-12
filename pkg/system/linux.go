package system

import (
	"fmt"

	"github.com/GabrielLuizSF/go_url/pkg/system/browser"
	"github.com/GabrielLuizSF/go_url/pkg/system/commands"
	"github.com/GabrielLuizSF/go_url/pkg/urls"
)

func LinuxOS(){

	command := MyCommand();
	switch command{
		case "ls":
			commands.Execute("ls");
		case "i":
			commands.Execute("whoami");
		case "firefox":
			browser.FIREFOX();
		case "github":
			fmt.Println(urls.Github());
		default:
			fmt.Println("UNDEFINIED COMMAND");
	}

	
}