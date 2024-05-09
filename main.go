package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	saveFile := flag.Bool("save", false, "")
	saveas := flag.String("as", "", "")
	flag.Parse()
	url := flag.Arg(0)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}

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
		fmt.Println("file succesfully saved")
	} else {
		fmt.Println(string(body))
	}
}
