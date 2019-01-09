package gosketch

import (
	"html/template"
	"net/http"
	"strconv"
)

var tmpl = `<html>
<head>
    <title>Hello World!</title>
</head>
<body>

    {{ . }}
</body>
</html>
`

func (css *Css) CreateHTML(w http.ResponseWriter) {
	t := template.New("main")
	t, _ = t.Parse(tmpl)
	result := "<div style='position:relative;width:1000px;height:1000px'>"
	for _, item := range css.Value {
		for _, i := range item.Struct {
			result = result + getElement(i)
		}
	}
	result = result + "</div>"
	t.Execute(w, template.HTML(result))
}

func getElement(block BlockCss) string {
	background := ""
	if block.BackgroundColor != "" {
		background = "background:" + block.BackgroundColor + ";"
	}
	typeBlock := "div"
	if block.Font != nil {
		typeBlock = "span"
	}
	top := "top:" + strconv.FormatFloat(block.Top, 'f', 0, 64) + "px;"
	left := "left:" + strconv.FormatFloat(block.Left, 'f', 0, 64) + "px;"
	width := "width:" + strconv.FormatFloat(block.Width, 'f', 0, 64) + "px;"
	height := "height:" + strconv.FormatFloat(block.Height, 'f', 0, 64) + "px;"
	position := "position: absolute;"
	r := "<" + typeBlock + "style=' + top + left + width + height + position + background "'>"
	for _, item := range block.Children {
		r = r + getElement(item)
	}
	if typeBlock == "span" {
		font, ok := block.Font.(Font)
		if ok {
			r = r + font.Text
		}
	}
	r = r + "</" + typeBlock + ">"
	return r
}
