package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type BatteryStatus struct {
	id     int
	status string
	power  string
}

func SplitStatus(stat string) []string {
	colIndex := strings.Index(stat, ":")
	commaIndex := strings.Index(stat, ",")
	return []string{stat[colIndex+1 : commaIndex], stat[commaIndex+1:]}
}

func main() {
	// Define a string flag "-d" with a default value of "all"
	dFlag := flag.String("d", "all", "Specify a value for the -d flag (all, power, status)")

	// Parse the command-line arguments
	flag.Parse()

	// Access the value of the flag
	dValue := *dFlag

	// Check for acpi in path
	_, err := exec.LookPath("acpi")
	if err != nil {
		fmt.Println("battery uses acpi to function")
		fmt.Println("acpi is not found in PATH")
	}
	// Running acpi
	cmd := exec.Command("acpi")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	// Parsing result
	parsedStat := SplitStatus(string(output))
	stat := BatteryStatus{id: 0, status: strings.Trim(parsedStat[0], " "), power: strings.Trim(parsedStat[1], " ")}

	// Output result
	switch dValue {
	case "all":
		// fmt.Printf("Battery %d: Status are %s and at %s", stat.id, stat.status, stat.power)
		fmt.Printf("%d: %s %s", stat.id, stat.status, stat.power)
	case "power":
		fmt.Print(stat.power)
		// fmt.Printf("Battery is at: %s", stat.power)
	case "status":
		fmt.Print(stat.status)
		// fmt.Printf("Battery status is: %s", stat.status)
	default:
		// Invalid value, provide a message or take appropriate action
		fmt.Println("Invalid value for -d. Please use 'all', 'power', or 'status'.")
	}
}
