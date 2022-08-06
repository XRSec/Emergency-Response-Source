package pkg

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

type ERInfoApp struct {
	SystemOS      string
	Hostname      string
	KernelVersion string
	IPAddr        string
	MemAll        float64
	MemUse        float64
	MemUsePercent float64
	Cpu           string
	cpuPercent    float64
	OnlineUser    []host.UserStat
}

var ERInfoApi ERInfoApp

func (e *ERInfoApp) Do() {
	e.GetOS()
	e.GetKernelVersion()
	e.GetHostname()
	e.GetOutboundIP()
	e.GetOnlineUser()
	e.GetCpu()

	e.View()
}

/*
	这里应该持续统计一分钟，可以 go ERApi.ERInfoApp.View() 放后台执行
*/
func (e *ERInfoApp) View() {
	fmt.Println("----------------------------------------------------")
	fmt.Printf("操作系统: %v\n主机名  : %v\nIP 地址 : %v\n", e.SystemOS, e.Hostname, e.IPAddr)
	fmt.Printf("总内存  : %.2fGB\n已用    : %.2fGB\n使用率  : %.2f%%\n", e.MemAll, e.MemUse, e.MemUsePercent)
	if e.SystemOS != "Windows" {
		fmt.Println("在线用户:")
	}
	for i := 0; i < len(e.OnlineUser); i++ {
		user := e.OnlineUser[i]
		started := strconv.Itoa(user.Started)
		if len(started) == 10 {
			started = time.Unix(int64(user.Started), 0).Format("2006-01-02 15:04:05")
		}
		from := user.Host
		if from == "" {
			from = "localhost/null"
		}
		fmt.Printf("NUM: %d USER: %-5v TTY: %-6v FROM: %-15v LOGIN@: %v\n", i, user.User, user.Terminal, from, started)
	}
	fmt.Printf("CPU: %v\n使用率: %.2f%%", e.Cpu, e.cpuPercent)
}

func (e *ERInfoApp) GetOS() {
	v, err := mem.VirtualMemory()
	if err != nil {
		log.Println(err)
		ERApi.ERErrorsApp.Do(err)
		e.MemAll = 0
		e.MemUse = 0
		e.MemUsePercent = 0
		return
	}
	e.MemAll = float64(v.Total) / 1024 / 1024 / 1024
	e.MemUse = float64(v.Used) / 1024 / 1024 / 1024
	e.MemUsePercent = v.UsedPercent
}

func (e *ERInfoApp) GetKernelVersion() {
	version, err := host.KernelVersion()
	if err != nil {
		log.Println(err)
		ERApi.ERErrorsApp.Do(err)
		e.KernelVersion = "Unknown"
		return
	}
	e.KernelVersion = version
}

func (e *ERInfoApp) GetHostname() {
	str, err := os.Hostname()
	if err != nil {
		log.Println(err)
		ERApi.ERErrorsApp.Do(err)
		e.Hostname = "Unknown"
		return
	}
	e.Hostname = str
}

func (e *ERInfoApp) GetOutboundIP() {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		ERApi.ERErrorsApp.Do(err)
		e.IPAddr = "Unknown"
		return
	}
	defer func(conn net.Conn) {
		if err := conn.Close(); err != nil {
			ERApi.ERErrorsApp.Do(err)
		}
	}(conn)
	e.IPAddr = conn.LocalAddr().(*net.UDPAddr).IP.String()
}

func (e *ERInfoApp) GetOnlineUser() {
	users, err := host.Users()
	if err != nil {
		log.Println(err)
		ERApi.ERErrorsApp.Do(err)
		e.OnlineUser = nil
		return
	}
	e.OnlineUser = users
}

func (e *ERInfoApp) GetCpu() {
	physicalCnt, _ := cpu.Counts(false)
	logicalCnt, _ := cpu.Counts(true)
	totalPercent, _ := cpu.Percent(5*time.Second, false)
	e.Cpu = fmt.Sprintf("%d 核心 %d 能效", physicalCnt, logicalCnt)
	e.cpuPercent = totalPercent[0] / float64(physicalCnt)
}
