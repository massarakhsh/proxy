package main

import (
	"github.com/massarakhsh/servnet/base"
	"github.com/massarakhsh/servnet/controller/gate"
	"github.com/massarakhsh/servnet/ruler"
	"github.com/massarakhsh/servnet/task/web"
	"fmt"
	"os"
	"time"

	"github.com/massarakhsh/lik"
)

func main() {
	lik.SetLevelInf()
	lik.SayError("System started")
	if !getArgs() {
		return
	}
	if !base.OpenDB(ruler.HostServ, ruler.HostBase, ruler.HostUser, ruler.HostPass) {
		return
	}
	ruler.RootCreator = gate.BuildGate
	go web.StartHttp()
	for !ruler.IsStoping() {
		time.Sleep(time.Second)
	}
	time.Sleep(time.Second * 3)
}

func getArgs() bool {
	args, ok := lik.GetArgs(os.Args[1:])
	if val := args.GetString("port"); val != "" {
		ruler.HostPort = lik.StrToInt(val)
	}
	if val := args.GetString("serv"); val != "" {
		ruler.HostServ = val
	}
	if val := args.GetString("base"); val != "" {
		ruler.HostBase = val
	}
	if val := args.GetString("user"); val != "" {
		ruler.HostUser = val
	}
	if val := args.GetString("pass"); val != "" {
		ruler.HostPass = val
	}
	if val := args.GetString("debug"); val != "" {
		ruler.DebugLevel = lik.StrToInt(val)
	}
	if len(ruler.HostBase) <= 0 {
		fmt.Println("HostBase name must be present")
		ok = false
	}
	if !ok {
		fmt.Println("Usage: proxy [-key val | --key=val]...")
		fmt.Println("port    - port value (80)")
		fmt.Println("serv    - Database server")
		fmt.Println("base    - Database name")
		fmt.Println("user    - Database user")
		fmt.Println("pass    - Database pass")
	}
	return ok
}

