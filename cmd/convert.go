package main

import (
	"log"
	"os"
	"tpot_scribe/pkg/convert"
)

func main() {

	filepath := "test/paragraph.docx"

	// Save as HTML
	h, err := convert.DocxToHTML(filepath)
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("test/output.html", []byte(h), 0644)

	// Save as JSON
	j, err := convert.DocxToJSON(filepath)
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("test/output.json", []byte(j), 0644)

	// Save as YAML
	y, err := convert.DocxToYAML(filepath)
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("test/output.yaml", []byte(y), 0644)
}
