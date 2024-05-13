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
	saveFile := flag.Bool("save", false, "") // --save user wants to save the output
	saveas := flag.String("as", "", "")      // --as user wants to change default name (for eg. with .html websites)
	flag.Parse()
	url := flag.Arg(0)

	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace)
	spinner, _ := pterm.DefaultSpinner.Start("getting response...")

	resp, err := http.Get(url)
	if err != nil {
		resp, err = http.Get("https://" + url) // if normal http.Get fails, user probably forgot https://
		if err != nil {
			spinner.Fail("")
			logger.Fatal(err.Error())
		}
	}

	sizeBytes, err := strconv.Atoi(resp.Header.Get("Content-Length"))              // get filesize in bytes
	spinner.UpdateText(fmt.Sprintf("reading response of %vMB", sizeBytes/1000000)) // update spinner text to show the file size
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		spinner.Fail("")
		logger.Fatal(err.Error())
	}

	urlSplit := strings.Split(url, "/")
	filename := urlSplit[len(urlSplit)-1]
	if *saveas != "" {
		filename = *saveas // TODO: verify if it is actuall filename (different method?)
	}

	workingDir, err := os.Getwd() // get working directory
	if err != nil {
		spinner.Fail("")
		logger.Fatal(err.Error())
	}

	if *saveFile { //save the file if specified
		if _, err := os.Stat(filepath.Join(workingDir, filename)); err == nil { // if file exists, user needs to be informed
			spinner.UpdateText("File exists, overwriting in 5s...")
			time.Sleep(time.Second * 5)
		}
		err = os.WriteFile(filename, body, 0644)
		if err != nil {
			spinner.Fail("")
			logger.Fatal(err.Error())
		}
		spinner.Success("Succesfully saved file")
	} else { // if nothing is specified to do with the file, dump the output
		fmt.Println(string(body))
		spinner.Success("Dumped data into term")
	}
}
