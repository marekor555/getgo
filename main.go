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
	saveFile := flag.Bool("save", false, "")
	saveas := flag.String("as", "", "")
	flag.Parse()
	url := flag.Arg(0)

	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace)
	spinner, _ := pterm.DefaultSpinner.Start("getting response...")

	resp, err := http.Get(url)
	if err != nil {
		spinner.Fail("")
		logger.Fatal(err.Error())
	}

	sizeBytes, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	spinner.UpdateText(fmt.Sprintf("reading response of %vMB", sizeBytes/1000000))
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		spinner.Fail("")
		logger.Fatal(err.Error())
	}

	urlSplit := strings.Split(url, "/")
	filename := urlSplit[len(urlSplit)-1]
	if *saveas != "" {
		filename = *saveas
	}

	workingDir, err := os.Getwd()
	if err != nil {
		spinner.Fail("")
		logger.Fatal(err.Error())
	}
	if _, err := os.Stat(filepath.Join(workingDir, filename)); err == nil {
		spinner.UpdateText("File exists, overwriting in 5s...")
		time.Sleep(time.Second * 5)
	}

	if *saveFile {
		err = os.WriteFile(filename, body, 0644)
		if err != nil {
			spinner.Fail("")
			logger.Fatal(err.Error())
		}
		spinner.Success("Succesfully saved file")
	} else {
		fmt.Println(string(body))
		spinner.Success("Dumped data into term")
	}
	println()
}
