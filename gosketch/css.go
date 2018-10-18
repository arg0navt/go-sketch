package gosketch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	Border          []string
	Shadow          []string
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
	Border          []string
	Shadow          []string
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

func getStyleArtboard(a Artboard, result *[]interface{}) {
	var newBlock BlockCss
	newBlock.Width = a.Frame.Width
	newBlock.Height = a.Frame.Height
	newBlock.Left = a.Frame.X
	newBlock.Top = a.Frame.Y
	newBlock.BorderRadius = 0
	newBlock.BackgroundColor = getFormatsColor(a.BackgroundColor)
	newBlock.Shadow = getShadow(a.Style.Shadows)
	newBlock.Border = getBorder(a.Style.Borders)
	fmt.Println(newBlock)
}

func getFormatsColor(c Color) ColorCss {
	rgba := "rgba(" + strconv.Itoa(int(c.Red*255)) + ", " + strconv.Itoa(int(c.Green*255)) + ", " + strconv.Itoa(int(c.Blue*255)) + ", " + strconv.FormatFloat(c.Alpha, 'f', 2, 64) + ")"
	hex := "#" + strconv.FormatInt(int64(c.Red*255), 16) + strconv.FormatInt(int64(c.Green*255), 16) + strconv.FormatInt(int64(c.Blue*255), 16)
	return ColorCss{RGBA: rgba, HEX: hex}
}

func getShadow(s []Shadow) []string {
	var result []string
	for _, item := range s {
		if item.IsEnabled == true {
			x := strconv.Itoa(int(item.OffsetX)) + "px "
			y := strconv.Itoa(int(item.OffsetY)) + "px "
			blur := strconv.Itoa(int(item.BlurRadius)) + "px "
			color := getFormatsColor(item.Color).RGBA
			result = append(result, x+y+blur+color)
		}
	}
	return result
}

func getBorder(b []Border) []string {
	var result []string
	for _, item := range b {
		if item.IsEnabled == true {
			t := strconv.Itoa(int(item.Thickness)) + "px "
			color := getFormatsColor(item.Color).RGBA
			result = append(result, t+"solid "+color)
		}
	}
	return result
}
