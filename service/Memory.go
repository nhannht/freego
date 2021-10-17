package service

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Memory struct {
	memFree      float64
	memTotal     float64
	memAvailable float64
	percent      float64
}

func (m *Memory) MemAvailable() float64 {
	return m.memAvailable
}

func (m *Memory) SetMemAvailable(memAvailable float64) {
	m.memAvailable = memAvailable
}

func (m *Memory) Percent() float64 {
	return m.percent
}

func (m *Memory) SetPercent(percent float64) {
	m.percent = percent
}

func (m *Memory) MemFree() float64 {
	return m.memFree
}

func (m *Memory) SetMemFree(memFree float64) {
	m.memFree = memFree
}

func (m *Memory) MemTotal() float64 {
	return m.memTotal
}

func (m *Memory) SetMemTotal(memTotal float64) {
	m.memTotal = memTotal
}

func (m *Memory) UpdateMemory() {
	_, err := os.Stat("/proc/meminfo")
	if err != nil {
		log.Fatalln("File not exist")
	}
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		log.Fatalln("Some thing wrong when read file /proc/meminfo file")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		//log.Println(line[1])
		if strings.Contains(line[0], "MemTotal") {
			total, err := strconv.ParseFloat(line[1], 64)
			m.SetMemTotal(total)
			if err != nil {
				log.Println(err)
			}
		}
		if strings.Contains(line[0], "MemAvailable") {
			available, _ := strconv.ParseFloat(line[1], 64)
			m.SetMemAvailable(available)
		}
		percentUsed := m.MemAvailable() / m.MemTotal()
		m.SetPercent(1 - percentUsed)
	}
}
