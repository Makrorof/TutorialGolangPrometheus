package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func updateServerCPU() {
	stat, _ := cpu.Percent(time.Second, false)

	//cpuStat, _ := cpu.Times(true)
	//totalUsage := 0.0
	//
	//for _, ce := range cpuStat {
	//	fmt.Println(ce)
	//
	//	//totalUsage += ce.
	//}
	//
	//totalUsage = totalUsage / float64(len(cpuStat))

	cpuStats, err := cpu.Times(false)

	if err != nil {
		fmt.Println(err)
	} else if len(cpuStats) == 0 || len(stat) == 0 {
		return
	}

	cpuTime := cpuStats[0]

	serverCpuBasicSecondsTotalGauge.WithLabelValues("User").Set(cpuTime.User)
	serverCpuBasicSecondsTotalGauge.WithLabelValues("System").Set(cpuTime.System)
	serverCpuBasicSecondsTotalGauge.WithLabelValues("Idle").Set(cpuTime.Idle)
	serverCpuBasicSecondsTotalGauge.WithLabelValues("Nice").Set(cpuTime.Nice)
	serverCpuBasicSecondsTotalGauge.WithLabelValues("Iowait").Set(cpuTime.Iowait)
	serverCpuBasicSecondsTotalGauge.WithLabelValues("Irq").Set(cpuTime.Irq)
	serverCpuBasicSecondsTotalGauge.WithLabelValues("Softirq").Set(cpuTime.Softirq)
	serverCpuBasicSecondsTotalGauge.WithLabelValues("Steal").Set(cpuTime.Steal)
	serverCpuBasicSecondsTotalGauge.WithLabelValues("Guest").Set(cpuTime.Guest)
	serverCpuBasicSecondsTotalGauge.WithLabelValues("GuestNice").Set(cpuTime.GuestNice)

	serverCpuUsageGauge.Set(stat[0])
	//cpuBasicSecondsTotalGauge.WithLabelValues("CPU").Set(stat[0])

	//infoStat, _ := cpu.Info()
	//for _, ce := range infoStat {
	//	fmt.Println(ce)
	//}
}

func main() {
	setMetricsHandle()

	//CPU, ram gibi metricler eklenmesi gerekiyor.

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			updateServerCPU()
		}
	}()

	for i := 0; i < 100; i++ {
		go func() {
			count := 0
			for {
				count++

				if count >= 100000 {
					count = 0
					time.Sleep(time.Millisecond)
				}
			}
		}()
	}

	http.ListenAndServe(":9191", nil)
}
