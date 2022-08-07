package pkg

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	net2 "github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type ERInfoApp struct {
	SystemOS             string
	Hostname             string
	KernelVersion        string
	Platform             string
	PlatformFamily       string
	PlatformVersion      string
	KennelVersion        string
	IPAddr               string
	VirtualMemAll        float64
	VirtualMemUse        float64
	VirtualMemUsePercent float64
	SwapMemAll           float64
	SwapMemUse           float64
	SwapMemUsePercent    float64
	Cpu                  string
	CpuPercent           float64
	OnlineUser           []host.UserStat
	Disk                 []Disk
	BootTime             string
	Processes            []*process.Process
	NetInterfaces        []net2.InterfaceStat
	NetConnections       []net2.ConnectionStat
}

type Disk struct {
	Device      string  `json:"device"`
	Mountpoint  string  `json:"mountpoint"`
	Fstype      string  `json:"fstype"`
	Total       float64 `json:"total"`
	Used        float64 `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

var ERInfoApi ERInfoApp

func (e *ERInfoApp) Do() {
	e.GetVirtualMemory()
	e.GetSwapMemory()
	e.GetPlatform()
	e.GetProcesses()
	e.GetHostname()
	e.GetOutboundIP()
	e.GetNetInterface()
	e.GetOnlineUser()
	e.GetCpu()
	e.GetDisk()
	e.GetBootTime()
	e.GetNetConnections()
	e.test()

	e.View()
}

/*
	这里应该持续统计一分钟，可以 go ERApi.ERInfoApp.View() 放后台执行
*/
func (e *ERInfoApp) View() {
	fmt.Println("----------------------------------------------------")
	fmt.Printf("操作系统: %v\n主机名  : %v\nIP 地址 : %v\n\n", e.SystemOS, e.Hostname, e.IPAddr)
	fmt.Printf("平台信息: %v %v %v\n内核版本: %v\n\n", e.Platform, e.PlatformFamily, e.PlatformVersion, e.KernelVersion)
	fmt.Printf("物理总内存  : %.2fGB\n已用    : %.2fGB\n使用率  : %.2f%%\n\n", e.VirtualMemAll, e.VirtualMemUse, e.VirtualMemUsePercent)
	fmt.Printf("虚拟总内存  : %.2fGB\n已用    : %.2fGB\n使用率  : %.2f%%\n\n", e.SwapMemAll, e.SwapMemUse, e.SwapMemUsePercent)
	if e.SystemOS != "windows" {
		onlineUserNum := len(e.OnlineUser)
		fmt.Println("在线用户: ", onlineUserNum)
		for i := 0; i < onlineUserNum; i++ {
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
	}
	fmt.Printf("\nCPU: %v\n使用率: %.2f%%\n\n", e.Cpu, e.CpuPercent)
	fmt.Printf("监测到磁盘: %v块\n", len(e.Disk))
	for i := 0; i < len(e.Disk); i++ {
		diskInfo := e.Disk[i]
		fmt.Printf("磁盘 %d: %v\n", i+1, diskInfo.Mountpoint)
		fmt.Printf("磁盘类型: %v\n", diskInfo.Fstype)
		fmt.Printf("磁盘大小: %.2fGB\n", diskInfo.Total)
		fmt.Printf("已用    : %.2fGB\n", diskInfo.Used)
		fmt.Printf("使用率  : %.2f%%\n", diskInfo.UsedPercent)
	}
	fmt.Printf("\n系统启动时间: %v\n\n", e.BootTime)
	for l := 0; l < len(e.NetInterfaces); l++ {
		v := e.NetInterfaces[l]
		n := len(v.Addrs)
		if n > 0 {
			for i := 0; i < n; i++ {
				if !strings.Contains(v.Addrs[i].Addr, ":") && !strings.Contains(v.Addrs[i].Addr, "127.0.0.1") {
					fmt.Printf("Name: %v Addr: %v\n", v.Name, v.Addrs[i].Addr)
				}
			}
		}
	}
	fmt.Printf("\n第0个进程信息: %v\n", e.Processes[0].String())
	fmt.Printf("第0条网络连接信息: %v\n", e.NetConnections[0])

}

func (e *ERInfoApp) GetVirtualMemory() {
	v, err := mem.VirtualMemory()
	if err != nil {
		log.Println(err)
		ERApi.ERErrorsApp.Do(err)
		e.VirtualMemAll = 0
		e.VirtualMemUse = 0
		e.VirtualMemUsePercent = 0
		return
	}
	e.VirtualMemAll = float64(v.Total) / 1024 / 1024 / 1024
	e.VirtualMemUse = float64(v.Used) / 1024 / 1024 / 1024
	e.VirtualMemUsePercent = v.UsedPercent
}

func (e *ERInfoApp) GetSwapMemory() {
	v, err := mem.SwapMemory()
	if err != nil {
		log.Println(err)
		ERApi.ERErrorsApp.Do(err)
		e.SwapMemAll = 0
		e.SwapMemUse = 0
		e.SwapMemUsePercent = 0
		return
	}
	e.SwapMemAll = float64(v.Total) / 1024 / 1024 / 1024
	e.SwapMemUse = float64(v.Used) / 1024 / 1024 / 1024
	e.SwapMemUsePercent = v.UsedPercent
}

func (e *ERInfoApp) GetPlatform() {
	kennelVersion, err := host.KernelVersion()
	if err != nil {
		log.Println(err)
		ERApi.ERErrorsApp.Do(err)
		e.KernelVersion = "Unknown"
		return
	}
	e.KernelVersion = kennelVersion
	platform, family, version, err := host.PlatformInformation()
	if err != nil {
		log.Println(err)
		ERApi.ERErrorsApp.Do(err)
		e.Platform = "Unknown"
		e.PlatformFamily = "Unknown"
		e.PlatformVersion = "Unknown"
		return
	}
	e.Platform = platform
	e.PlatformFamily = family
	e.PlatformVersion = version
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

func (e *ERInfoApp) GetNetInterface() {
	counters, err := net2.Interfaces()
	if err != nil {
		log.Println(err)
		ERApi.ERErrorsApp.Do(err)
		e.NetInterfaces = nil
		return
	}
	//for l := 0; l < len(counters); l++ {
	//	v := counters[l]
	//	n := len(v.Addrs)
	//	if n > 0 {
	//		for i := 0; i < n; i++ {
	//			if !strings.Contains(v.Addrs[i].Addr, ":") && !strings.Contains(v.Addrs[i].Addr, "127.0.0.1") {
	//				fmt.Printf("Name: %v Addr: %v\n", v.Name, v.Addrs[i].Addr)
	//			}
	//		}
	//	}
	//}
	e.NetInterfaces = counters
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
	totalPercent, _ := cpu.Percent(3*time.Second, false)
	e.Cpu = fmt.Sprintf("%d 核心 %d 能效", physicalCnt, logicalCnt)
	e.CpuPercent = totalPercent[0]
}

func (e *ERInfoApp) GetDisk() {
	infos, err := disk.Partitions(false)
	if err != nil {
		log.Println(err)
		ERApi.ERErrorsApp.Do(err)
		return
	}
	for i := 0; i < len(infos); i++ {
		info, err := disk.Usage(infos[i].Mountpoint)
		if err != nil {
			log.Println(err)
			ERApi.ERErrorsApp.Do(err)
			continue
		}
		e.Disk = append(e.Disk, Disk{Device: infos[i].Device, Mountpoint: infos[i].Mountpoint, Fstype: infos[i].Fstype, Total: float64(info.Total) / 1024 / 1024 / 1024, Used: float64(info.Used) / 1024 / 1024 / 1024, UsedPercent: info.UsedPercent})
	}
}

func (e *ERInfoApp) GetBootTime() {
	timestamp, err := host.BootTime()
	if err != nil {
		log.Println(err)
		ERApi.ERErrorsApp.Do(err)
		e.BootTime = "Unknown"
		return
	}
	e.BootTime = time.Unix(int64(timestamp), 0).Local().Format("2006-01-02 15:04:05")
}

func (e *ERInfoApp) GetProcesses() {
	// TODO 优化网络连接信息
	////获取到所有进程的详细信息
	//p1, _ := process.Pids()  //获取当前所有进程的pid
	//fmt.Println("p1:",p1)
	//p2,_ := process.GetWin32Proc(1120)  //对应pid的进程信息
	//fmt.Println("p2:",p2)
	////fmt.Println(p2[0].ParentProcessID)  //获取父进程的pid

	processes, err := process.Processes()
	if err != nil {
		log.Println(err)
		ERApi.ERErrorsApp.Do(err)
		e.Processes = nil
		return
	}
	e.Processes = processes
}

func (e *ERInfoApp) GetNetConnections() {
	v, err := net2.Connections("all")
	if err != nil {
		log.Println(err)
		ERApi.ERErrorsApp.Do(err)
		e.NetConnections = nil
		return
	}
	e.NetConnections = v
}

func (e *ERInfoApp) test() {

}
