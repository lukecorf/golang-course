package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitorNumber = 3
const delay = 5

func main() {

	introduction()
	for {
		showMenu()
		command := readConsole()

		// if command == 1 {
		// 	fmt.Println("Monitoring...")
		// } else if command == 2 {
		// 	fmt.Println("Logs:")
		// } else if command == 0 {
		// 	fmt.Println("Exiting Program.")
		// } else {
		// 	fmt.Println("Invalid Command")
		// }

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Logs:")
			printLogs()
		case 0:
			fmt.Println("Exiting Program.")
			os.Exit(0)
		default:
			fmt.Println("Invalid Command")
			os.Exit(-1)
		}
	}

}

func returnNameAndAgeAndVersion() (string, int, float32) {
	return "Luke", 120, 1.0
}

func introduction() {
	name, _, version := returnNameAndAgeAndVersion()
	fmt.Println("Hello Sr. ", name)
	fmt.Println("System Version: ", version)
}

func showMenu() {
	fmt.Println("1. Start Monitoring")
	fmt.Println("2. Show Logs")
	fmt.Println("0. Exit")
}

func readConsole() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("You choose the command: ", command)
	return command
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	// sites := []string{"https://random-status-code.herokuapp.com/", "https://www.google.com", "https://www.apple.com"}
	sites := readWebsiteFile()
	for i := 0; i < monitorNumber; i++ {
		for i, site := range sites {
			fmt.Println("Position:", i, "Value:", site)
			testWebsite(site)
		}
		fmt.Println("")
		time.Sleep(delay * time.Second)
	}

	fmt.Println("")
}

func testWebsite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Error:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Website:", site, "is ok!")
		writeLog(site, true)
	} else {
		fmt.Println("Website:", site, "with problems. Status Code", resp.StatusCode)
		writeLog(site, false)
	}
}

func readWebsiteFile() []string {
	var websites []string

	file, err := os.Open("websites.txt")

	if err != nil {
		fmt.Println("Error:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		websites = append(websites, line)
		if err == io.EOF {
			break
		}
	}
	file.Close()
	return websites
}

func writeLog(website string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " " + website + "- online: " + strconv.FormatBool(status) + "\n")

	file.Close()

}

func printLogs() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(string(file))
}
