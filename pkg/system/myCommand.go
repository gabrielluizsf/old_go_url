package system

import "fmt"

func MyCommand()string{
	var command string;
		fmt.Printf("=> ");
		fmt.Scanf("%s",&command);
		return command;
}
