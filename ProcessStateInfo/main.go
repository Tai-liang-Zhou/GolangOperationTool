package main

import (
	"fmt"
	"sort"
	"time"

	// ps "github.com/mitchellh/go-ps"
	ps "github.com/shirou/gopsutil/process"
)

type ProcessInfo struct {
	ProcessName string
	CPUPercent  float64
	MEMPercent  float32
	MENMB       uint64
	CREATETime  string
}

func main() {

	pids, err := ps.Pids()
	if err != nil {
		fmt.Println(err)
	}

	processInfos := []ProcessInfo{}

	for _, value := range pids {

		ok, err := ps.PidExists(value)
		if err != nil {
			fmt.Println(err)
		}
		if ok {

			processItem, err := ps.NewProcess(value)
			if err != nil {
				fmt.Println(err)
			}
			processName, err := processItem.Name()
			if err != nil {
				fmt.Println(err)
			}

			cpuPercent, err := processItem.CPUPercent()
			if err != nil {
				fmt.Println(err)
			}
			memPercent, err := processItem.MemoryPercent()
			if err != nil {
				fmt.Println(err)
			}

			memInfo, err := processItem.MemoryInfo()
			if err != nil {
				fmt.Println(err)
			}

			createtime, err := processItem.CreateTime()
			if err != nil {
				fmt.Println(err)
			}

			resTime := time.Unix(createtime/1000, 0)

			tempInfo := ProcessInfo{
				ProcessName: processName,
				CPUPercent:  cpuPercent,
				MEMPercent:  memPercent,
				MENMB:       (memInfo.RSS) / 1024 / 1024,
				CREATETime:  time.Since(resTime).String(),
			}
			processInfos = append(processInfos, tempInfo)
		}
	}

	sort.Slice(processInfos, func(i, j int) bool {
		return float64(processInfos[i].MEMPercent) > float64(processInfos[j].MEMPercent)
	})

	for _, value := range processInfos[:10] {
		fmt.Printf("CPU: %.3f %%, MEM: %.3f %%(%d MB), Process: %s, CreateTime : %s\n\r", value.CPUPercent, value.MEMPercent, value.MENMB, value.ProcessName, value.CREATETime)
	}

}
