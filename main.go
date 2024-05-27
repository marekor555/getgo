package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/pterm/pterm"
)

func main() {
	saveFile := flag.Bool("save", false, "save file")      // --save user wants to save the output
	saveAs := flag.String("as", "", "set custom filename") // --as user wants to change default name (for eg. with .html websites)
	help := flag.Bool("help", false, "display help")
	flag.Parse()
	url := flag.Arg(0)

	if *help {
		fmt.Println("example commands:")
		fmt.Println("save: \n  getgo --save \"example.com/downloadme.txt\"")
		fmt.Println("save with filename: \n  getgo --save --as hello.html \"duckduckgo.com\"")
		fmt.Println("dump: \n  getgo \"example.com/dumpme.txt\" ")
		return
	}

	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace)
	spinner, _ := pterm.DefaultSpinner.Start("getting response...")

	response, error := http.Get(url)
	if error != nil {
		response, error = http.Get("https://" + url) // if normal http.Get fails, user probably forgot https://
		if error != nil {
			spinner.Fail("")
			logger.Fatal(error.Error())
		}
	}

	sizeBytes, error := strconv.Atoi(response.Header.Get("Content-Length"))        // get filesize in bytes
	spinner.UpdateText(fmt.Sprintf("reading response of %vMB", sizeBytes/1000000)) // update spinner text to show the file size
	responseBody, error := io.ReadAll(response.Body)
	if error != nil {
		spinner.Fail("")
		logger.Fatal(error.Error())
	}

	var filename string

	if *saveAs != "" {
		filename = *saveAs
	} else {
		// if no name specified, get the filename with it's extension
		urlSplit := strings.Split(url, "/")
		if strings.Split(response.Header.Get("Content-Type"), ";")[0] != "text/html" { // if it is not text/html, it probably is already specified
			filename = urlSplit[len(urlSplit)-1]
		} else {
			filenameSplit := strings.Split(urlSplit[len(urlSplit)-1], ".")
			filename = filenameSplit[0] + ".html"
		}
	}

	workingDirectory, error := os.Getwd() // get working directory
	if error != nil {
		spinner.Fail("")
		logger.Fatal(error.Error())
	}

	if *saveFile { //save the file if specified
		if _, error := os.Stat(filepath.Join(workingDirectory, filename)); error == nil { // if file exists, user needs to be informed
			spinner.UpdateText("File exists, overwriting in 5s...")
			time.Sleep(time.Second * 5)
		}
		error = os.WriteFile(filename, responseBody, 0644)
		if error != nil {
			spinner.Fail("")
			logger.Fatal(error.Error())
		}
		spinner.Success("Succesfully saved file")
	} else { // if nothing is specified to do with the file, dump the output
		spinner.Success("Dumping data into term")
		fmt.Println(string(responseBody))
	}
}
