package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func parseSize(size string) float64 {
	trimSize := strings.Trim(size, " ")
	sizeArr := strings.Split(trimSize, " ")
	sizeInt, err := strconv.Atoi(sizeArr[0])
	if err != nil {
		log.Fatalf("Failed to parse to int, %s", err)
	}
	return float64(sizeInt) / 1024.0 / 1024.0
}

func main() {
	// Reading /proc/meminfo to gain info about memory usage
	meminfo, err := os.Open("/proc/meminfo")
	if err != nil {
		log.Fatalf("Failed to open /proc/meminfo,%s", err)
	}
	defer meminfo.Close()

	// Flag stuff, ex. ./i3_processor -m=full
	mode := flag.String("m", "full", "Specify mode: full, partial, or percentage")
	delay := flag.Int("d", 5, "Specify delay in seconds")
	flag.Parse()

	for {

		meminfoScan := bufio.NewScanner(meminfo)
		var memFree, MemTotal float64
		// Scan for needed label, ie "MemTotal" in /proc/meminfo. then break loop when all required value are found
		count := 0
		for meminfoScan.Scan() {
			line := meminfoScan.Text()
			lineTokens := strings.Split(line, ":")

			switch lineTokens[0] {
			case "MemTotal":
				MemTotal = parseSize(lineTokens[1])
				count++
			case "MemFree":
				memFree = parseSize(lineTokens[1])
				count++
			case "Buffers":
				memFree += parseSize(lineTokens[1])
				count++
			case "Cached":
				memFree += parseSize(lineTokens[1])
				count++
			}
			if count == 4 {
				break
			}

		}

		// Flag stuff
		memUsed := MemTotal - memFree
		switch *mode {
		case "full":
			fmt.Printf("%.2fGB/%.2fGB\n", memUsed, MemTotal)
			fmt.Printf("%.2f%% Memory used\n", memUsed/MemTotal*100)
		case "partial":
			fmt.Printf("%.2fGB/%.2fGB\n", memUsed, MemTotal)
		case "percentage":
			fmt.Printf("%.2f%% Memory used", memUsed/MemTotal*100)
		default:
			fmt.Println("Invalid mode. Please use full, partial, or percentage.")
		}

		err = meminfoScan.Err()
		if err != nil {
			log.Fatalf("Error during scanning /proc/meminfo, %s", err)
		}

		// Reset current offest to beginning of file since scanner keep internal state and would carry on between loop
		meminfo.Seek(0, 0)
		// Delay of loop
		time.Sleep(time.Duration(*delay) * time.Second)
	}
}
