package system

import (
	"fmt"

	"github.com/GabrielLuizSF/go_url/pkg/system/browser"
	"github.com/GabrielLuizSF/go_url/pkg/system/commands"
	"github.com/GabrielLuizSF/go_url/pkg/urls"
)
//this function should only be run on linux :)
func LinuxOS(){
	command := MyCommand();
	switch command{
		case "ls":
			commands.Execute("ls");
		case "i":
			commands.Execute("whoami");
    case "hello":
      fmt.Println("Hello ", myOSUsername("[ERROR]: Cannot get OS username\n[RESPONSE]: "))
		case "firefox":
			browser.FIREFOX();
		case "github":
			fmt.Println("[CTRL CLICK] =>","[",urls.Github(),"]");
    case "whatsapp":
      fmt.Println("[CTRL CLICK] =>","[",urls.Whatsapp(),"]");
    case "google":
      fmt.Println("[CTRL CLICK] =>","[",urls.Google(),"]");
    case "youtube":
      fmt.Println("[CTRL CLICK] =>","[",urls.Youtube(),"]");
		default:
			fmt.Println("UNDEFINIED COMMAND");
	}

	
}
