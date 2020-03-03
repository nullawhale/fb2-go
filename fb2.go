package main

import "encoding/xml"

type FB2 struct {
	FictionBook xml.Name `xml:"FictionBook"`
	Description struct {
		TitleInfo struct {
			Genres []string `xml:"genre"`
			Author []struct {
				FName string `xml:"first-name"`
				MName string `xml:"middle-name"`
				LName string `xml:"last-name"`
			} `xml:"author"`
			BookTitle  string `xml:"book-title"`
			Annotation struct {
				P []string `xml:"p"`
			} `xml:"annotation"`
			Date       string `xml:"date"`
			Lang       string `xml:"lang"`
			SrcLang    string `xml:"src-lang"`
			Translator struct {
				FName string `xml:"first-name"`
				MName string `xml:"middle-name"`
				LName string `xml:"last-name"`
			} `xml:"translator"`
		} `xml:"title-info"`

		PublishInfo struct {
			BookName  string `xml:"book-name"`
			Publisher string `xml:"publisher"`
			City      string `xml:"city"`
			Year      int    `xml:"year"`
			ISBN      string `xml:"isbn"`
			Sequence  struct {
				Name string `xml:"name"`
			} `xml:"sequence"`
		} `xml:"publish-info"`

		CustomInfo []struct {
			InfoType string `xml:"info-type"`
		} `xml:"custom-info"`
	} `xml:"description"`

	Body struct {
		Sections []struct {
			P []string `xml:"p"`
		} `xml:"section"`
	} `xml:"body"`
}
