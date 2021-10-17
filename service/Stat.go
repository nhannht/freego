package service

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Stat struct {
	totalTime  float64
	idleTime   float64
	cpuPercent float64
}

func (s *Stat) TotalTime() float64 {
	return s.totalTime
}

func (s *Stat) SetTotalTime(totalTime float64) {
	s.totalTime = totalTime
}

func (s *Stat) IdleTime() float64 {
	return s.idleTime
}

func (s *Stat) SetIdleTime(idleTime float64) {
	s.idleTime = idleTime
}

func (s *Stat) CpuPercent() float64 {
	return s.cpuPercent
}

func (s *Stat) SetCpuPercent(cpuPercent float64) {
	s.cpuPercent = cpuPercent
}

func (s *Stat) UpdateCpu() {
	//fmt.Println("CPU usage % at 1 second intervals:\n")
	prevIdleTime, prevTotalTime := readStatFile()

	time.Sleep(time.Millisecond * 100)
	// Second time

	idleTime, totalTime := readStatFile()

	deltaIdleTime := idleTime - prevIdleTime
	deltaTotalTime := totalTime - prevTotalTime
	cpuUsage := 1.0 - float64(deltaIdleTime)/float64(deltaTotalTime)
	s.SetCpuPercent(cpuUsage)
	//fmt.Printf("%d : %6.3f\n", i, cpuUsage)
}

func readStatFile() (uint64, uint64) {
	file, err := os.Open("/proc/stat")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	firstLine := scanner.Text()[5:] // get rid of cpu plus 2 spaces
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	split := strings.Fields(firstLine)
	currentIdleTime, _ := strconv.ParseUint(split[3], 10, 64)
	currentTotalTime := uint64(0)
	for _, s := range split {
		u, _ := strconv.ParseUint(s, 10, 64)
		currentTotalTime += u
	}
	return currentIdleTime, currentTotalTime
}
