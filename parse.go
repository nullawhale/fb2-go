package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	path := os.Args[1]

	//var fb2 FB2

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	decoder := xml.NewDecoder(bytes.NewReader(data))
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			return
		}
		if err != nil {
			panic(err)
		}
		if token == nil {
			break
		}
		switch el := token.(type) {
		case xml.CharData:
			if !trimText(string(el)) {
				printP(string(el))
			}
		}
	}
}

func trimText(s string) bool {
	res := strings.Trim(s, "\n")
	res = strings.Trim(s, "\t")
	res = strings.TrimSpace(s)

	if res == "" {
		return true
	} else {
		return false
	}
}

func printP(s string) {
	words := strings.Fields(s)
	fmt.Println(words)
}
