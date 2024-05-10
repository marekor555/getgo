package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/pterm/pterm"
)

func main() {
	saveFile := flag.Bool("save", false, "")
	saveas := flag.String("as", "", "")
	flag.Parse()
	url := flag.Arg(0)

	spinner, _ := pterm.DefaultSpinner.Start("getting response...")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	sizeBytes, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	spinner.UpdateText(fmt.Sprintf("reading response of %vMB", sizeBytes/1000000))
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	urlSplit := strings.Split(url, "/")
	filename := urlSplit[len(urlSplit)-1]
	if *saveas != "" {
		filename = *saveas
	}

	if *saveFile {
		err = os.WriteFile(filename, body, 0644)
		if err != nil {
			log.Fatal(err.Error())
		}
		spinner.Success("Succesfully saved file")
	} else {
		fmt.Println(string(body))
		spinner.Success("Dumped data into term")
	}
	println()
}
