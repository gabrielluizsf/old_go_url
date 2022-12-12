package analyzer

import "fmt";

func ARCH_ANALYZER(){
	switch MyARCH(){
	case "amd64":
		fmt.Println(MyARCH());
	case "amd64p32":
		fmt.Println(MyARCH());
	default:
		fmt.Println("ARCH UNDEFINIED")

}
}