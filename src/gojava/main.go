package main

import (
	"fmt"
	"gojava/classfile"
	"gojava/classpath"
	"strings"
)

func main() {
	debugJVM()
	//	cmd := parseCmd()
	//	if cmd.versionFlag {
	//		fmt.Printf("version 0.0.1")
	//	} else if cmd.helpFlag || cmd.class == "" {
	//		printUsage()
	//	} else {
	//		startJVM(cmd)
	//	}
}

func debugJVM() {
	cp := classpath.Parse("", "D:/nolink/workspace/soa-test/target/classes")
	className := "com/ctrip/finance/easypay/GauseTest"
	cf := loadClass(className, cp)
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class: com.ctrip.finance.easypay.GauseTest\n")
	}
}

func startJVM(cmd *Cmd) {
	//fmt.Printf("classpath:%s class:%s args: %v\n", cmd.CpOption, cmd.Class, cmd.Args)
	cp := classpath.Parse(cmd.xJreOption, cmd.cpOption)
	//fmt.Printf("classpath:%s class:%s args: %v\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	//	classData, _, err := cp.ReadClass(className)
	//	if err != nil {
	//		fmt.Printf("Could not load main class %s\n.", cmd.class)
	//		return
	//	}
	//	fmt.Printf("class data:%v\n", classData)
	cf := loadClass(className, cp)
	//	fmt.Println(cmd.class)
	//	printClassInfo(cf)
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class: %s\n", cmd.class)
	}
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	fmt.Printf("class data:%v\n", classData)
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf(" %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf(" %s\n", m.Name())
	}
}
