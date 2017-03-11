
package main

import (
	"fmt"
	"strings"
	"gojava/classpath"
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
	//fmt.Printf("classpath:%s class:%s args: %v\n", cmd.CpOption, cmd.Class, cmd.Args)
	cp := classpath.Parse(cmd.XJreOption, cmd.CpOption)
	fmt.Printf("classpath:%s class:%s args: %v\n", cp, cmd.Class, cmd.Args)
	className := strings.Replace(cmd.Class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not load main class %s\n.", cmd.Class)
		return
	}
	fmt.Printf("class data:%v\n", classData)
}