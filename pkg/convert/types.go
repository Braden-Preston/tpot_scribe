package convert

type Document struct {
	Styles []NamedStyle   `json:"styles" yaml:"styles,omitempty"`
	Blocks []BlockElement `json:"blocks" yaml:"blocks,omitempty"`
}

type BlockElement struct {
	Tag      string          `json:"tag" yaml:"tag,omitempty"`
	Text     string          `json:"text,omitempty" yaml:"text,omitempty"`
	CSS      *CSS            `json:"css,omitempty" yaml:"css,omitempty"`
	Children []InlineElement `json:"children,omitempty" yaml:"children,omitempty"`
}

type InlineElement struct {
	Tag  string `json:"tag" yaml:"tag,omitempty"`
	Text string `json:"text" yaml:"text,omitempty"`
	CSS  *CSS   `json:"css,omitempty" yaml:"css,omitempty"`
}

type NamedStyle struct {
	Name string `json:"name" yaml:"name,omitempty"`
	CSS  *CSS   `json:"css,omitempty" yaml:"css,omitempty"`
}

type CSS struct {
	Color      string `json:"color,omitempty" yaml:"color,omitempty"`
	Background string `json:"background,omitempty" yaml:"background,omitempty"`
	FontSize   int    `json:"fontSize,omitempty" yaml:"fontSize,omitempty"`
	FontFamily int    `json:"fontFamily,omitempty" yaml:"fontFamily,omitempty"`
	Bold       bool   `json:"bold,omitempty" yaml:"bold,omitempty"`
	Italic     bool   `json:"italic,omitempty" yaml:"italic,omitempty"`
	Underline  bool   `json:"underline,omitempty" yaml:"underline,omitempty"`
	TextAlign  string `json:"textAlign,omitempty" yaml:"textAlign,omitempty"`
}
