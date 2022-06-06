package data

import "fmt"

type StatisticCpu struct {
	UserPercent   float64 `json:"userPercent"`
	SystemPercent float64 `json:"systemPercent"`
	IdlePercent   float64 `json:"idlePercent"`
	CoreCount     int     `json:"coreCount"`
}

type StatisticMemory struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
	Cached      uint64  `json:"cached"`
	Free        uint64  `json:"free"`
	FreePercent float64 `json:"freePercent"`
}

type StatisticNetwork struct {
	Name          string `json:"name"`
	RxBytesTotal  uint64 `json:"rxBytesTotal"`
	RxBytesPerSec uint64 `json:"rxBytesPerSec"`
	TxBytesTotal  uint64 `json:"txBytesTotal"`
	TxBytesPerSec uint64 `json:"txBytesPerSec"`
}

type Statistic struct {
	Cpu      StatisticCpu
	Memory   StatisticMemory
	Networks []StatisticNetwork
}

func (s Statistic) Print() {
	fmt.Printf("cpu core count: %d\n", s.Cpu.CoreCount)
	fmt.Printf("cpu user: %f %%\n", s.Cpu.UserPercent)
	fmt.Printf("cpu system: %f %%\n", s.Cpu.SystemPercent)
	fmt.Printf("cpu idle: %f %%\n", s.Cpu.IdlePercent)

	fmt.Printf("memory total: %d bytes, %d MB\n", s.Memory.Total, s.Memory.Total/1024/1024)
	fmt.Printf("memory used: %d bytes, %d MB\n", s.Memory.Used, s.Memory.Used/1024/1024)
	fmt.Printf("memory used percent: %f %%\n", s.Memory.UsedPercent)
	fmt.Printf("memory cached: %d bytes, %d MB\n", s.Memory.Cached, s.Memory.Cached/1024/1024)
	fmt.Printf("memory free: %d bytes, %d MB\n", s.Memory.Free, s.Memory.Free/1024/1024)
	fmt.Printf("memory free percent: %f %%\n", s.Memory.FreePercent)

	for _, n := range s.Networks {
		fmt.Printf("network %s rx: %d bytes, %d MB\n", n.Name, n.RxBytesTotal, n.RxBytesTotal/1024/1024)
		fmt.Printf("network %s rx/s: %d bytes, %d MB\n", n.Name, n.RxBytesPerSec, n.RxBytesPerSec/1024/1024)
		fmt.Printf("network %s tx: %d bytes, %d MB\n", n.Name, n.TxBytesTotal, n.TxBytesTotal/1024/1024)
		fmt.Printf("network %s tx/s: %d bytes, %d MB\n", n.Name, n.TxBytesPerSec, n.TxBytesPerSec/1024/1024)
	}
}
