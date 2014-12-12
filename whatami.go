package main

import (
	"fmt"
	"log"
	"os"
	linuxproc "github.com/c9s/goprocinfo/linux"
	humanize "github.com/dustin/go-humanize"
)

func printHostName() {
	hostName, err := os.Hostname()
	if err != nil {
		log.Fatal("Kernel, adımı söylemiyor!")
	}
	fmt.Printf("┏━%s\n", hostName)
}

func printCPUInfo() {
	cpuInfo, err := linuxproc.ReadCPUInfo("/proc/cpuinfo")
	if err != nil {
		log.Fatal("İşlemci bilgisine erişilemedi!")
	}
	fmt.Printf("┣━━ %d tane işlemci\n", len(cpuInfo.Processors))
	for _, cpu := range cpuInfo.Processors {
		fmt.Printf("┣━━━━ %10s\n", cpu.ModelName)
	}
}

func printMemoryInfo() {
	memInfo, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		log.Fatal("Bellek bilgisine erişilemedi!")
	}
	fmt.Printf("┣━━ %s toplam bellek\n", humanize.IBytes(memInfo["MemTotal"] * 1024) )
}

func printDiskInfo(){
	diskInfo, err := linuxproc.ReadDisk("/")
	if err != nil {
		log.Fatal("Disk bilgisine erişilemedi!")
	}
	fmt.Printf("┗━━ %s toplam disk alanı\n", humanize.IBytes(diskInfo.All) )

}

func main() {
	printHostName()
	printCPUInfo()
	printMemoryInfo()
	printDiskInfo()
}