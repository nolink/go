
package main

import (
	"fmt"
)


func main() {
	cmd := ParseCmd()
	if cmd.VersionFlag{
		fmt.Printf("version 0.0.1")
	}else if cmd.HelpFlag || cmd.Class == ""{
		PrintUsage()
	}else{
		StartJVM(cmd)
	}
}


func StartJVM(cmd *Cmd){
	fmt.Printf("classpath:%s class:%s args: %v\n", cmd.CpOption, cmd.Class, cmd.Args)
}