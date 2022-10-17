//go:build linux
// +build linux

package meme

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type MemData struct {
	Buffers      float64
	Cached       float64
	MemAvailable float64
	MemFree      float64
	MemTotal     float64
	Shared       float64
	Used         float64

	SwapFree  float64
	SwapTotal float64
	SwapUsed  float64
}

func toFloat(raw string) float64 {
	if raw == "" {
		return 0
	}
	res, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 0
	}
	return res
}

func parseLine(raw string) (string, float64) {
	// remove "kB" at the end and remove spaces
	text := strings.ReplaceAll(raw[:len(raw)-2], " ", "")
	// split by the semicolon
	keyValue := strings.Split(text, ":")
	return keyValue[0], toFloat(keyValue[1])
}

func GetMemInfo() (MemData, error) {
	res := MemData{}

	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return res, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		key, value := parseLine(scanner.Text())
		switch key {
		case "MemTotal":
			res.MemTotal = value

		case "Buffers":
			res.Buffers = value
		case "Shmem":
			res.Shared = value
		case "Cached":
			res.Cached += value
		case "SReclaimable":
			res.Cached += value
		case "MemFree":
			res.MemFree = value
		case "MemAvailable":
			res.MemAvailable = value

		case "SwapTotal":
			res.SwapTotal = value
		case "SwapFree":
			res.SwapFree = value
		}
	}

	res.Used = res.MemTotal - res.MemFree - res.Buffers - res.Cached
	res.SwapUsed = res.SwapTotal - res.SwapFree

	return res, nil
}
