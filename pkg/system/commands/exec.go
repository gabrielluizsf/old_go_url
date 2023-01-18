package commands

import (
	"fmt"
    "os/exec"
	"github.com/GabrielLuizSF/go_url/warning"
)
//this function executes operating system commands
func Execute(command string){
	out, err := exec.Command(command).Output();
		warning.PRINT_DEFAULT_ERRORS(err,"INVALID COMMAND");
	output := string(out[:])
	fmt.Println(output)
}
