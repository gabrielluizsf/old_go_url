package warning

import "fmt"

func PRINT_DEFAULT_ERRORS(ERROR error,MESSAGE string){
	if ERROR != nil{
	fmt.Println("[MESSAGE]:",MESSAGE,"\n[ERROR]:",ERROR);
	}else{
    fmt.Println("[Command successfully executed]");
  }
}
