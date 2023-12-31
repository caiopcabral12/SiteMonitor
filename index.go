package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const timeUpload = 5

func Intro() {

	fmt.Println("Hey! What's your name?")

	var name string
	fmt.Scan(&name)

	fmt.Println("Hello Mister", name, ", Welcome to SiteXereck!")
	fmt.Println(" - Choose an option:")
	fmt.Println("")

}

func Comands() {

	fmt.Println("       1 - Initiate program")
	fmt.Println("       2 - Logs")
	fmt.Println("       3 - Exit")
}

func ExecComand() int {
	var comand int
	fmt.Scan(&comand)
	fmt.Println("")
	fmt.Println("You choose comand ", comand)
	fmt.Println("")

	return comand
}

func main() {

	Intro()
	for {
		Comands()
		comand := ExecComand()

		switch comand {
		case 1:
			openSite()
		case 2:
			fmt.Println("Sites logs bitch")
			writeLog("")
		case 3:
			fmt.Println("See you soon!  :)")
			os.Exit(0)
		default:
			fmt.Printf("Unknown command! Choose a number between 1 and 3")
		}
	}
}

func TestSite(lines string) {
	resp, err := http.Get(lines)

	if err != nil {
		fmt.Println("Error:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("The website:", lines, "is working perfectly!")
		registerLog(lines, true)
	} else {
		fmt.Println("The website:", lines, "is Offline  :( - ERROR", resp.StatusCode)
		registerLog(lines, false)
	}
}

func TestAllSites(lines []string) {
	for i, site := range lines {

		if site != "Test them all!" {
			fmt.Printf("Testing site %d: %s\n", i+1, site)
			TestSite(site)
		}

	}
	fmt.Println("All sites tested.")
}

func openSite() []string {

	fmt.Println("Sites available:")
	fmt.Println("")

	var lines []string

	data, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	reader := bufio.NewReader(data)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		lines = append(lines, line)

		if err == io.EOF {
			break
		}
	}

	data.Close()

	for i, site := range lines {
		fmt.Println("- Option", i+1, ":", site)
	}

	fmt.Println("")
	fmt.Println("What site do you want?")

	var option int
	fmt.Scan(&option)

	if option < 1 || option > len(lines) {
		fmt.Println("Bad request! Choose a number between 1 and", len(lines))
		fmt.Println("")

	}

	for {
		if lines[option-1] == "Test them all!" {
			TestAllSites(lines)
		} else {
			TestSite(lines[option-1])
		}
		time.Sleep(timeUpload * time.Second)
		fmt.Println("")
	}
}

func registerLog(lines string, status bool) {
	logs, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	date := time.Now()

	if err != nil {
		fmt.Println("Error:", err)
	}

	if status == true {
		logs.WriteString("- " + date.Format(string("02/01/2006 03:04:05")) + " - Website online: " + lines + "\n")
	} else {
		logs.WriteString("- " + date.Format(string("02/01/2006 03:04:05")) + " - Website offline: " + lines + "\n")
	}

	logs.Close()
}

func writeLog(logs string) {
	log, err := os.ReadFile("logs.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(log))
}
