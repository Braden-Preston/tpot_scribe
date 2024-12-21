package convert

import (
	"fmt"
	"strconv"

	"github.com/gomutex/godocx/wml/ctypes"
)

func CssFromProps(para *ctypes.ParagraphProp, run *ctypes.RunProperty) CSS {
	css := CSS{}
	if para != nil {
		if para.Justification != nil {
			css.TextAlign = string(para.Justification.Val)
		}
	}
	if run != nil {
		if run.Color != nil {
			css.Color = fmt.Sprintf("#%s", run.Color.Val)
		}
		if run.Size != nil {
			css.FontSize = fmt.Sprintf("%dpx", run.Size.Value)
		}
		if run.Bold != nil {
			isBold, _ := strconv.ParseBool(string(*run.Bold.Val))
			css.Bold = isBold
		}
		if run.Italic != nil {
			isItalic, _ := strconv.ParseBool(string(*run.Italic.Val))
			css.Italic = isItalic
		}
		if run.Underline != nil {
			isUnderline := string(run.Underline.Val) == "single"
			css.Underline = isUnderline
		}
		if run.Shading != nil {
			if string(*run.Shading.Fill) != "auto" {
				css.Background = fmt.Sprintf("#%s", *run.Shading.Fill)
			}
		}
	}
	return css
}
