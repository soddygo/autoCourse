package course

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	in          string
	out         string
	args        []string
}

func PaseCmd() *Cmd {
	cmd := &Cmd{}
	flag.BoolVar(&cmd.helpFlag, "help", false, "帮助信息尚未编辑")
	flag.BoolVar(&cmd.helpFlag, "?", false, "帮助信息尚未编辑")
	flag.StringVar(&cmd.in, "in", "", "绝对路径的输入excel模板")
	flag.StringVar(&cmd.out, "out", "", "绝对路径的输出excel")
	flag.Parse()

	//args := flag.Args()
	//if len(args) > 0 {
	//	cmd.class = args[0]
	//	cmd.args = args[1:]
	//}

	return cmd
}

func (self *Cmd) GetVersion() bool {
	return self.versionFlag
}

func (self *Cmd) GetHelp() bool {
	return self.helpFlag
}
func (self *Cmd) GetIn() string {
	return self.in
}
func (self *Cmd) GetOut() string {
	return self.out
}


func PrintUsage() {
	fmt.Printf("Usage: %s [-options] in [args...]\n", os.Args[0])
	//flag.PrintDefaults()
}