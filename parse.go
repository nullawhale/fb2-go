package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	path := os.Args[1]

	xmlFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println(err)
	}

	var fb2 FB2

	xml.Unmarshal(byteValue, &fb2)

	for i := 0; i < len(fb2.Description.TitleInfo.Genres); i++ {
		fmt.Println("Genre: " + fb2.Description.TitleInfo.Genres[i])
	}

	var author string = fb2.Description.TitleInfo.Author[0].FName + " " +
		fb2.Description.TitleInfo.Author[0].MName + " " +
		fb2.Description.TitleInfo.Author[0].LName

	var bookTitle string = fb2.Description.TitleInfo.BookTitle

	var anotation string = fb2.Description.TitleInfo.Annotation.P[0] + "\n" +
		fb2.Description.TitleInfo.Annotation.P[1]

	var bookYear string = fb2.Description.TitleInfo.Date

	fmt.Println(author + "\n" + bookTitle + "\n" + bookYear + "\n" + anotation)
}
