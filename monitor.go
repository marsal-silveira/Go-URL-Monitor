package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	// "io/ioutil"
	"bufio"
	"net/http"
	"os"
	"strconv"
	"time"
)

const _version = 1.1
const _monitoring = 2
const _delay = 1

const _urlsFileName = "urls.txt"
const _logFileName = "log.txt"

func main() {

	showHeader()

	for {
		showOptions()

		var option = readOption()

		switch option {
		case 0:
			exit(false)
		case 1:
			startMonitoring()
		case 2:
			readLogs()
		default:
			fmt.Println("")
			fmt.Println("Opção inválida!\nFavor informe um valor entre 0 e 2")
			// exit(true)
		}
	}
}

func showHeader() {

	fmt.Println("Olá")
	fmt.Println("Este programa está na versão:", _version)
}

func showOptions() {

	fmt.Println("")
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func readOption() int {

	var option int
	fmt.Scan(&option)

	return option
}

func exit(error bool) {

	fmt.Println("")
	fmt.Println("Saindo do Programa!")

	var option = 0
	if error {

		option = -1
	}
	os.Exit(option)
}

func startMonitoring() {

	fmt.Println("")
	fmt.Println("Iniciando Monitoramento...")
	var urls = readURLFromFile()

	for i := 0; i < _monitoring; i++ {

		// for i := 0; i < len(urls); i++ {
		for index, url := range urls {

			fmt.Println("")
			fmt.Println("Testando url:", index, "->", url)
			testURL(url)
		}

		time.Sleep(_delay * time.Second)
	}
}

func testURL(url string) {

	var response, err = http.Get(url)
	// fmt.Println(response)

	if err != nil {

		fmt.Println("Error on testing url:", url, "Error:", err)
		return
	}

	if response.StatusCode == 200 {

		fmt.Println("Url:", url, "is up!")
		writeLog(url, true)
	} else {

		writeLog(url, false)
		fmt.Println("Url:", url, "is down. Status Code:", response.StatusCode)
	}
}

func readURLFromFile() []string {

	var urls []string

	file, err := os.Open(_urlsFileName)
	// var file, err = ioutil.ReadFile(_fileName)
	if err != nil {

		fmt.Println("Error on opening file.", err)
		return nil
	}
	// fmt.Println(string(file))

	var reader = bufio.NewReader(file)
	for {

		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		urls = append(urls, line)
		// fmt.Println("Line:", line)

		if err == io.EOF {
			break
		}
	}

	file.Close()
	return urls
}

func writeLog(url string, status bool) {

	// file, err := os.Open(_logFileName)
	file, err := os.OpenFile(_logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {

		fmt.Println("Error on writing in log file.", err)
		return
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05 - ") + url + " - online: " + strconv.FormatBool(status) + "\n")

	// fmt.Println(file)
	file.Close()
}

func readLogs() {

	fmt.Println("")
	fmt.Println("Exibindo Logs...")

	file, err := ioutil.ReadFile(_logFileName)

	if err != nil {

		fmt.Println("Error reading log file.", err)
		return
	}

	fmt.Println(string(file))
}
