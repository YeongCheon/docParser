package main

import (
	"fmt"
	"os"

	"archive/zip"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

var lineCharacter string

func init() {
	lineCharacter = "\n"
}

func GetDocumentContent(filename string) (documentContent string) {
	if !strings.HasSuffix(filename, "docx") {
		fmt.Println("Not a valid docx file")
		return
	}

	r, err := zip.OpenReader(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	var f *zip.File
	var found bool

	for _, f = range r.File {
		if strings.EqualFold(f.Name, "word/document.xml") {
			found = true
			break
		}
	}

	if !found {
		log.Fatal("Not a valid docxfile")
		return
	}

	rc, err := f.Open()

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		log.Fatal(err)
	}
	defer rc.Close()

	documentContent = string(data)
	documentContent = strings.Replace(documentContent, "<w:rPr>", lineCharacter, -1)

	reg, err := regexp.Compile("<.*?>")
	if err != nil {
		log.Fatal(err)
		return
	}
	return reg.ReplaceAllString(documentContent, "")
}

func SearchKeywords(data string, keywords map[string]bool) {
	for keyword, _ := range keywords {
		if strings.Contains(data, keyword) {
			keywords[keyword] = true
		}
	}
}

func SaveTextFile(fileName, fileContent string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(fileContent)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	/*
		if len(os.Args) < 3 {
			fmt.Println("\n Usage: GoFinder <filePath> <keyword1> [<keyword2>...n]" +
				"\n e.g. GoFinder /tmp/sample.docx C++ Java\n")
			return
		}

		keywords := make(map[string]bool)
		for _, keyword := range os.Args[2:] {
			keywords[keyword] = false
		}

		filePath := os.Args[1]

		documentContent := GetDocumentContent(filePath)
		fmt.Println(documentContent)
		if len(documentContent) == 0 {
			fmt.Println("Empty document")
			return
		}

		SearchKeywords(documentContent, keywords)

		for keyword, exists := range keywords {
			fmt.Println(keyword, exists)
		}
	*/
}
