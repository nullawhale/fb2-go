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

var N = 64
var T = 4

func main() {
	path := os.Args[1]

	name := ""
	isSection := false
	isTitle := false
	isPoem := false
	var bodyAttrs = make(map[string]string)
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
		switch el := token.(type) {
		case xml.StartElement:
			name = el.Name.Local
			if name == "section" {
				isSection = true
			} else if name == "title" {
				isTitle = true
			} else if name == "poem" {
				fmt.Println()
				isPoem = true
			} else if name == "body" {
				for _, attr := range el.Attr {
					bodyAttrs[attr.Name.Local] = attr.Value
				}
			}
		case xml.CharData:
			if isTitle && name == "p" {
				s := string(el)
				if strings.TrimSpace(s) != "" {
					if bodyAttrs["name"] == "notes" {
						printTitle(s, false)
					} else {
						printTitle(s, true)
					}
				}
			} else if isSection && name == "p" {
				s := string(el)
				if strings.TrimSpace(s) != "" {
					printP(strings.Fields(s))
				}
			} else if isPoem && name == "emphasis" {
				s := string(el)
				if strings.TrimSpace(s) != "" {
					printTitle(s, false)
				}
			} else if isSection {
				switch name {
				case "strong", "emphasis", "a":
					s := string(el)
					if strings.TrimSpace(s) != "" {
						printP(strings.Fields(s))
					}
				}
			}
		case xml.EndElement:
			name = el.Name.Local

			if name == "section" {
				isSection = false
			} else if name == "title" {
				isTitle = false
			} else if name == "poem" {
				fmt.Println()
				isPoem = false
			} else if name == "body" {
				bodyAttrs = make(map[string]string)
			}
		}
	}
}

func printTitle(words string, newline bool) {
	_N := N - length(words)
	words = strings.Repeat(" ", _N/2) + words + strings.Repeat(" ", _N/2)

	if newline {
		fmt.Println(strings.Repeat(".", N))
	}
	fmt.Println(words)
	if newline {
		fmt.Println(strings.Repeat(".", N))
	}
}

func printP(words []string) {
	row := words[0]
	var ret []string
	sect := true
	for i := 1; i < len(words); i++ {
		if length(words[i])+length(row) < N {
			row += " " + words[i]
		} else {
			tmpWords := strings.Split(strings.TrimSpace(row), " ")
			if sect {
				tmpWords[0] = strings.Repeat(" ", T) + tmpWords[0]
			}
			var newRow string
			_N := N
			for _, w := range tmpWords {
				_N -= length(w)
			}
			for i := len(tmpWords) - 1; i > 0; i-- {
				space := strings.Repeat(" ", _N/i)
				newRow = space + tmpWords[i] + newRow
				_N -= _N / i
			}
			newRow = tmpWords[0] + newRow
			row = newRow
			ret = append(ret, row)
			row = words[i]
			sect = false
		}
	}
	row = row + strings.Repeat(" ", N-length(row))
	ret = append(ret, row)

	for _, w := range ret {
		fmt.Println(w)
	}
	//fmt.Println()
}

func length(s string) int {
	return len([]rune(s))
}
