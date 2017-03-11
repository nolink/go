
package main

import "flag"
import "fmt"
import "os"

type Cmd struct{
	HelpFlag bool
	VersionFlag bool
	CpOption string
	XJreOption string
	Class string
	Args []string
}

func ParseCmd() *Cmd{
	cmd := &Cmd{}
	flag.Usage = PrintUsage
	flag.BoolVar(&cmd.HelpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.HelpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.VersionFlag, "version", false, "pring version info")
	flag.StringVar(&cmd.CpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.CpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XJreOption, "XJre", "", "classpath")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.Class = args[0]
		cmd.Args = args[1:]
	}
	return cmd
}

func PrintUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}