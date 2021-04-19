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

const LOOPS_MONITOR int = 3
const LOOP_DELAY = 3

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
				printLogs()
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
	sites := readSites()
	fmt.Println("------------------------------------------")
	for i := 0; i < LOOPS_MONITOR; i++ {
		fmt.Println("Testing...", 3 - i, "times")
		for _, site := range sites {
			testSite(site)
		}
		fmt.Println("----------")
		time.Sleep(LOOP_DELAY * time.Second)
	}
}

func readSites() []string {
	fmt.Println("------------------------------------------")
	fmt.Println("Reading sites.txt File... 📁 📄 📄 📂")
	
	var sites []string
	file, err := os.Open("sites.txt")

	verifyErr(err)

	reader := bufio.NewReader(file)

	
	for {
		row, err := reader.ReadString('\n')
		
		row = strings.TrimSpace(row)
		
		sites = append(sites, row)
		if err == io.EOF {
			break
		}
	}
	
	file.Close()

	return sites
}

func testSite(site string) {
	var message string
	res, err := http.Get(site)

	verifyErr(err)

	if res.StatusCode == 200 {
		message = "On air 👏"
		logTest(site, true, "🎉🎉")
	} else {
		message = "Is down 😱"
		logTest(site, false, "😰😰")
	}
	fmt.Println("> Testing site:", site, "->", message, "witch statusCode:", res.StatusCode)
}

func logTest(site string, status bool, emoji string) {
	file, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	verifyErr(err)

	_, err = file.WriteString(time.Now().Format("2006-02-01 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + " " + emoji + " " + "\n")

	verifyErr(err)

	file.Close()
}

func verifyErr(err error) {
	if err != nil {
		fmt.Println("An Error Ocurred:", err)
		os.Exit(-1)
	}
}

func printLogs() {
	fmt.Println("------------------------------------------")
	fmt.Println("Reading file log.txt...")
	file, err := ioutil.ReadFile("log.txt")
	verifyErr(err)
	fmt.Println(string(file))
}