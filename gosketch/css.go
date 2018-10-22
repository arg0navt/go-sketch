package gosketch

import (
	"encoding/json"
	"errors"
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
	BackgroundColor string
	BackgroundImage string
	BorderRadius    float64
	Border          []string
	BoxShadow       string
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
	BoxShadow       string
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
			var newBlock BlockCss
			newBlock.getStyleBlock(&item, &newPage.Css)
		}
		// getStyle(&page.Layers, &newPage.Css)
		result = append(result, newPage)
	}
	json.NewEncoder(w).Encode(s.Pages)
}

func (b *BlockCss) getStyle(l *interface{}, result *[]interface{}) {
	var shadow string
	b.Width = a.Frame.Width
	b.Height = a.Frame.Height
	b.Left = a.Frame.X
	b.Top = a.Frame.Y
	b.BorderRadius = 0
	b.BackgroundColor = a.BackgroundColor.getFormatsColor().RGBA
	for index, item := range a.Style.Shadows {
		s, err := item.getShadow()
		if err == nil {
			if index > 0 {
				shadow = shadow + ", "
			}
			shadow = shadow + s
		}
	}
	b.BoxShadow = shadow
	fmt.Println(b)
}

func (c *Color) getFormatsColor() ColorCss {
	rgba := "rgba(" + strconv.Itoa(int(c.Red*255)) + ", " + strconv.Itoa(int(c.Green*255)) + ", " + strconv.Itoa(int(c.Blue*255)) + ", " + strconv.FormatFloat(c.Alpha, 'f', 2, 64) + ")"
	hex := "#" + strconv.FormatInt(int64(c.Red*255), 16) + strconv.FormatInt(int64(c.Green*255), 16) + strconv.FormatInt(int64(c.Blue*255), 16)
	return ColorCss{RGBA: rgba, HEX: hex}
}

func (s *Shadow) getShadow() (string, error) {
	var result string
	if s.IsEnabled == true {
		x := strconv.Itoa(int(s.OffsetX)) + "px "
		y := strconv.Itoa(int(s.OffsetY)) + "px "
		blur := strconv.Itoa(int(s.BlurRadius)) + "px "
		color := s.Color.getFormatsColor().RGBA
		result = x + y + blur + color
	} else {
		return result, errors.New("Disabled shadow")
	}
	return result, nil
}

// func getBorder(b []Border) []string {
// 	var result []string
// 	for _, item := range b {
// 		if item.IsEnabled == true {
// 			t := strconv.Itoa(int(item.Thickness)) + "px "
// 			color := getFormatsColor(item.Color).RGBA
// 			result = append(result, t+"solid "+color)
// 		}
// 	}
// 	return result
// }
