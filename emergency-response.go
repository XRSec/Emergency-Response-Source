package main

import (
	"flag"
	"fmt"
	. "github.com/XRSec/Emergency-Response-Source/pkg"
	"log"
	"runtime"
)

var (
	help, version                            bool
	buildTime, commitId, versionData, author string
)

func init() {
	log.SetPrefix("[Emergency-Response] ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	flag.BoolVar(&help, "help", false, "Display help information")
	flag.BoolVar(&help, "h", false, "Display help information")
	flag.BoolVar(&version, "version", false, "HTML-TO-MARKDOWN version")
	flag.BoolVar(&version, "v", false, "HTML-TO-MARKDOWN version")

	ERApi.ERInfoApp.SystemOS = runtime.GOOS
	switch ERApi.ERInfoApp.SystemOS {
	case "windows":
		log.Printf("操作系统: %v 暂不支持", ERApi.ERInfoApp.SystemOS)
		//os.Exit(0)
	case "darwin":
		log.Printf("操作系统: %v 暂不支持", ERApi.ERInfoApp.SystemOS)
		//os.Exit(0)
	case "linux":
		log.Printf("操作系统: %v, 欢迎使用!", ERApi.ERInfoApp.SystemOS)
	default:
		log.Printf("操作系统: %v 暂不支持", ERApi.ERInfoApp.SystemOS)
		//os.Exit(0)
	}
}

func main() {
	flag.Parse()
	if version {
		fmt.Printf("\n ╷──────────────────────────────────────────────────────────────────────────────╷ \n")
		fmt.Printf(" │                                                                              │\n")
		fmt.Printf(" │  Emergency-Response                                                          │\n")
		fmt.Printf(" │  Version: %6v\t | BuildTime: %18v                       │\n", versionData, buildTime)
		fmt.Printf(" │  Author: %7v\t | CommitId: %41v  │\n", author, commitId)
		fmt.Printf(" │                                                                              │\n")
		fmt.Printf(" ╵──────────────────────────────────────────────────────────────────────────────╵ \n\n")
		return
	}
	if help {
		flag.Usage()
		return
	}
	ERApi.ERInfoApp.Do()
	//if err := cmd.RootCmd.Execute(); err != nil {
	//	fmt.Fprintln(os.Stderr, err)
	//	os.Exit(1)
	//}
}
