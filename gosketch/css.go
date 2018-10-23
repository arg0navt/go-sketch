package gosketch

import (
	"encoding/json"
	"errors"
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

type MapColor struct {
	Value map[string]interface{}
}

type MapShadow struct {
	Value map[string]interface{}
}

func (s *SketchFile) GetCSS(w http.ResponseWriter, r *http.Request) {
	result := make([]interface{}, 0)
	for key, page := range s.Pages {
		blocks := make([]interface{}, 0)
		for _, item := range page.Layers {
			result := checkTypeLayer(&item)
			if result != nil {
				blocks = append(blocks, result)
			}
		}
		result = append(result, PageCss{ID: key, Css: blocks})
	}
	json.NewEncoder(w).Encode(result)
}

func checkTypeLayer(layer *map[string]interface{}) interface{} {
	switch (*layer)["_class"] {
	case "artboard", "group", "shapeGroup", "symbolMaster":
		var block BlockCss
		block.css(layer)
		return block
	default:

	}
	return nil
}

func (block *BlockCss) css(layer *map[string]interface{}) {
	frameM, okF := (*layer)["frame"].(map[string]interface{})
	if okF {
		block.Width = frameM["width"].(float64)
		block.Height = frameM["height"].(float64)
		block.Left = frameM["x"].(float64)
		block.Top = frameM["y"].(float64)
	}
	bkgM, ok := (*layer)["backgroundColor"].(map[string]interface{})
	if ok {
		bkg := MapColor{Value: bkgM}
		block.BackgroundColor = bkg.colorRGBA()
	}
	style, ok := (*layer)["style"].(map[string]interface{})
	if ok {
		shadowS, ok := style["shadow"].([]map[string]interface{})
		if ok {
			for index, item := range shadowS {
				shd := MapShadow{Value: item}
				itemShadow, err := shd.boxShadow()
				if err == nil {
					if index > 0 {
						block.BoxShadow = block.BoxShadow + ", "
					}
					block.BoxShadow = block.BoxShadow + itemShadow
				}
			}
		}
	}
	children := make([]interface{}, 0)
	for _, item := range (*layer)["layers"].([]interface{}) {
		childrenItem, ok := item.(map[string]interface{})
		if ok {
			result := checkTypeLayer(&childrenItem)
			if result != nil {
				children = append(children, result)
			}
		}
	}
	block.Children = children
}

func (c *MapColor) colorRGBA() string {
	r := strconv.Itoa(int((*c).Value["red"].(float64) * 255))
	g := strconv.Itoa(int((*c).Value["green"].(float64) * 255))
	b := strconv.Itoa(int((*c).Value["blue"].(float64) * 255))
	a := strconv.FormatFloat((*c).Value["alpha"].(float64), 'f', 2, 64)
	return "rgba(" + r + ", " + g + ", " + b + ", " + a + ")"
}

func (c *MapColor) colorHex() string {
	h := strconv.FormatInt(int64((*c).Value["red"].(float64)*255), 16)
	e := strconv.FormatInt(int64((*c).Value["green"].(float64)*255), 16)
	x := strconv.FormatInt(int64((*c).Value["blue"].(float64)*255), 16)
	return "#" + h + e + x
}

func (s *MapShadow) boxShadow() (string, error) {
	if s.Value["isEnabled"].(bool) {
		x := strconv.Itoa(int((*s).Value["offsetX"].(float64))) + "px "
		y := strconv.Itoa(int((*s).Value["offsetY"].(float64))) + "px "
		blur := strconv.Itoa(int((*s).Value["blurRadius"].(float64))) + "px "
		c := MapColor{Value: (*s).Value["color"].(map[string]interface{})}
		color := c.colorRGBA()
		return x + y + blur + color, nil
	}
	return "", errors.New("Disabled shadow")
}
