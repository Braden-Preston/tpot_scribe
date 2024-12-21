package convert

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"

	"github.com/gomutex/godocx"
	"gopkg.in/yaml.v3"
)

func ParseDocument(filepath string) (*Document, error) {

	document, err := godocx.OpenDocument(filepath)
	if err != nil {
		return nil, err
	}
	defer document.Close()

	parsedDocument := Document{}

	// Get named styles for parsed document
	for _, style := range document.DocStyles.StyleList {
		namedStyle := NamedStyle{Name: style.Name.Val}
		css := CssFromProps(style.ParaProp, style.RunProp)
		if (css != CSS{}) {
			namedStyle.CSS = &css
		}
		parsedDocument.Styles = append(parsedDocument.Styles, namedStyle)
	}

	// Get block elements for parse document with inlined styles
	for _, documentChild := range document.Document.Body.Children {
		paragraphNode := documentChild.Para
		paragraph := paragraphNode.GetCT()

		css := CssFromProps(paragraph.Property, paragraph.Property.RunProperty)
		block := BlockElement{
			Tag:      "p",
			Children: nil,
		}
		if (css != CSS{}) {
			block.CSS = &css
		}

		// For each element in paragraph, a.k.a "run"
		children := []InlineElement{}
		for _, paragraphChild := range paragraph.Children {
			runNode := paragraphChild.Run

			// Get the text content of all the runs
			carriageReturn := false
			texts := []string{}
			for _, run := range runNode.Children {
				if run.Text != nil {
					textContent := run.Text.Text
					texts = append(texts, textContent)
				}
				if run.Break != nil {
					breakType := *run.Break.BreakType
					if breakType == "textWrapping" {
						carriageReturn = true
					}
				}
			}
			textContent := strings.Join(texts, "")
			textContent = strings.TrimSpace(textContent)

			// Create a new inline element
			inlineElement := InlineElement{
				Tag:  "span",
				Text: textContent,
			}

			if textContent == "" {
				inlineElement.Tag = "br"
			}

			// Get CSS props
			css := CssFromProps(paragraph.Property, runNode.Property)
			if (css != CSS{}) {
				inlineElement.CSS = &css
			}

			// Append inline element to block's children
			if carriageReturn {
				children = append(children, InlineElement{
					Tag: "br",
				})
			} else {
				children = append(children, inlineElement)
			}

		}

		if len(children) != 1 {
			block.Children = append(block.Children, children...)
		} else {
			block.Text = children[0].Text
			block.CSS = children[0].CSS
		}

		parsedDocument.Blocks = append(parsedDocument.Blocks, block)
	}
	return &parsedDocument, nil
}

func DocxToHTML(filepath string) (string, error) {

	ctx := context.Background()

	document, err := ParseDocument(filepath)
	if err != nil {
		return "", err
	}

	buff := bytes.NewBuffer([]byte{})
	err = DocumentComponent("braden", *document).Render(ctx, buff)
	if err != nil {
		return "", err
	}
	html := strings.ReplaceAll(buff.String(), "styles=", "style=")
	return html, err
}

func DocxToJSON(filepath string) (string, error) {

	document, err := ParseDocument(filepath)
	if err != nil {
		return "", err
	}

	buff := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(buff)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(document)
	if err != nil {
		return "", err
	}

	return buff.String(), err
}

func DocxToYAML(filepath string) (string, error) {

	document, err := ParseDocument(filepath)
	if err != nil {
		return "", err
	}

	buff := bytes.NewBuffer([]byte{})
	encoder := yaml.NewEncoder(buff)

	err = encoder.Encode(document)
	if err != nil {
		return "", err
	}

	return buff.String(), err
}
