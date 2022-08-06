package main

import (
	. "github.com/XRSec/Emergency-Response-Source/pkg"
	"log"
	"runtime"
)

func init() {
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
	ERApi.ERInfoApp.Do()
	//if err := cmd.RootCmd.Execute(); err != nil {
	//	fmt.Fprintln(os.Stderr, err)
	//	os.Exit(1)
	//}
}
