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
	BackgroundColor string
	BackgroundImage string
	BorderRadius    float64
	BorderWidth     float64
	BorderColor     string
	BorderStyle     string
	Children        []interface{}
}

type TextCss struct {
	FontSize        float64
	FontColor       string
	FontFamily      string
	FontWeight      float64
	Width           float64
	Height          float64
	Top             float64
	Left            float64
	BackgroundColor string
	BackgroundImage string
	BorderRadius    float64
	BorderWidth     float64
	BorderColor     string
	BorderStyle     string
	Children        []interface{}
}

func (s *SketchFile) GetCSS(w http.ResponseWriter, r *http.Request) {
	var result []interface{}
	for key, page := range s.Pages {
		blocks := make([]interface{}, 0)
		newPage := PageCss{ID: key, Css: blocks}
		getStyle(&page.Layers, &newPage.Css)
		result = append(result, newPage)
	}
	json.NewEncoder(w).Encode(&result)
}

func getStyle(layer *[]interface{}, result *[]interface{}) {

}
