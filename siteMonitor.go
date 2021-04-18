package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitorLoops int = 3
const sleepTime = 3

func main() {
	name := "Guilherme"
	version := 1.1

	sayIntro(name, version)

	for {
		sayMenu()

		command := requestCommand()
	
		switch command {
			case 1:
				monit()
			case 2: 
				fmt.Println("------------------------------------------")
				fmt.Println("Selecionou 2")
			case 0:
				fmt.Println("------------------------------------------")
				fmt.Println("Thanks! 👋")
				fmt.Println("------------------------------------------")
				os.Exit(0)
			default:
				fmt.Println("------------------------------------------")
				fmt.Println("What? 🧐")
				os.Exit(-1)
		}
	}
}

func sayIntro(name string, version float64) {
	fmt.Println("------------------------------------------")
	fmt.Println("Hello Sr.", name)
	fmt.Println("Progam version:", version)
}

func sayMenu() {
	fmt.Println("------------------------------------------")
	fmt.Println("1 - Monitor sites")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Quit")
}

func requestCommand() int {
	var command int

	fmt.Scan(&command)
	return command
}

func monit() {
	sites := requestSites()
	fmt.Println("------------------------------------------")
	for i := 0; i < monitorLoops; i++ {
		fmt.Println("Testing...", 3 - i, "times")
		for _, site := range sites {
			testSite(site)
		}
		fmt.Println("----------")
		time.Sleep(sleepTime * time.Second)
	}
}

func requestSites() []string {
	fmt.Println("------------------------------------------")
	fmt.Println("Enter sites separated by commas with no space, like: 'https://example.com;https://mysite.com.br'")
	var sites string

	fmt.Scan(&sites)
	return strings.Split(sites, ";")
}

func testSite(site string) {
	var message string
	res, _ := http.Get(site)
	if res.StatusCode == 200 {
		message = "On air 👏"
	} else {
		message = "Is down 😱"
	}
	fmt.Println("> Testing site:", site, "->", message, "witch statusCode:", res.StatusCode)
}