package gosketch

import (
	"encoding/json"
	"net/http"
)

type PageCss struct {
	ID  string
	Css []interface{}
}

type BlockCss struct {
	Width           float64
	Height          float64
	Top             float64
	Left            float64
	BackgroundColor ColorCss
	BackgroundImage string
	BorderRadius    float64
	BorderWidth     float64
	BorderColor     ColorCss
	BorderStyle     string
	Children        []interface{}
}

type TextCss struct {
	FontSize        float64
	FontColor       ColorCss
	FontFamily      string
	FontWeight      float64
	Width           float64
	Height          float64
	Top             float64
	Left            float64
	BackgroundColor ColorCss
	BackgroundImage string
	BorderRadius    float64
	BorderWidth     float64
	BorderColor     ColorCss
	BorderStyle     string
	Children        []interface{}
}

type ColorCss struct {
	HEX  string
	RGBA string
}

func (s *SketchFile) GetCSS(w http.ResponseWriter, r *http.Request) {
	var result []interface{}
	for key, page := range s.Pages {
		blocks := make([]interface{}, 0)
		newPage := PageCss{ID: key, Css: blocks}
		for _, item := range page.Layers {
			switch item.(type) {
			case Artboard:
				getStyleArtboard(item.(Artboard), &newPage.Css)
			}
		}
		// getStyle(&page.Layers, &newPage.Css)
		result = append(result, newPage)
	}
	json.NewEncoder(w).Encode(s.Pages)
}

// func getStyle(layer *[]interface{}, result *[]interface{}) {

// }

func getStyleArtboard(a Artboard, result *[]interface{}) {
	var newBlock BlockCss
	newBlock.Width = a.Frame.Width
	newBlock.Height = a.Frame.Height
	newBlock.Left = a.Frame.X
	newBlock.Top = a.Frame.Y
	newBlock.BackgroundColor = getFormatsColor(a.BackgroundColor)
}

func getFormatsColor(c Color) ColorCss {
	return ColorCss{}
}
