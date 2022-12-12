package commands

import (
	"fmt"
    "os/exec"
	"github.com/GabrielLuizSF/go_url/warning"
)

func Execute(command string){
	out, err := exec.Command(command).Output();
		warning.PRINT_DEFAULT_ERRORS(err,"INVALID COMMAND");
		fmt.Println("[Command Successfully Executed]");
	output := string(out[:])
	fmt.Println(output)
}