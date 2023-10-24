package main

import (
	"fmt"
	"net/http"
	"os"
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
	openSite()
	Intro()
	for {
		Comands()
		comand := ExecComand()

		switch comand {
		case 1:
			Start()
		case 2:
			Logs()
		case 3:
			Quit()
		default:
			fmt.Printf("Unknown command! Choose a number between 1 and 3")
		}
	}
}

func Start() {
	fmt.Println("Sites available:")
	fmt.Println("")
	sites := []string{"https://random-status-code.herokuapp.com/", "https://google.com"}
	sites = append(sites, "Test them all!")

	for i, site := range sites {
		fmt.Println("- Site", i+1, ":", site)
	}
	fmt.Println("")
	fmt.Println("What site do you want?")

	var option int
	fmt.Scan(&option)

	if option < 1 || option > len(sites) {
		fmt.Println("Bad request! Choose a number between 1 and", len(sites))
		fmt.Println("")

		return
	}

	for {
		if sites[option-1] == "Test them all!" {
			TestAllSites(sites)
		} else {
			TestSite(sites[option-1])
		}
		time.Sleep(timeUpload * time.Second)
		fmt.Println("")
	}

}

func Logs() {
	fmt.Println("Sites logs bitch")
}

func Quit() {
	fmt.Println("See you soon!")
	os.Exit(0)
}

func TestSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Error:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("The website:", site, "is working perfectly!")
	} else {
		fmt.Println("The website:", site, "is Offline  :( - ERROR", resp.StatusCode)
	}
}

func TestAllSites(sites []string) {
	for i, site := range sites {

		if site != "Test them all!" {
			fmt.Printf("Testing site %d: %s\n", i+1, site)
			TestSite(site)
		}

	}
	fmt.Println("All sites tested.")
}

func openSite() {
	links, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Erorr:", err)
	} else {
		fmt.Println(links)
	}
}
