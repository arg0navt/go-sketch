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
	Font            Font
}

type Font struct {
	Size   float64
	Color  ColorCss
	Family string
	Weight float64
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
			var block BlockCss
			block.cssBlock(item)
			blocks = append(blocks, block)
		}
		result = append(result, PageCss{ID: key, Css: blocks})
	}
	json.NewEncoder(w).Encode(result)
}

func (block *BlockCss) cssBlock(layer map[string]interface{}) {
	frame, okF := layer["frame"].(map[string]interface{})
	if okF {
		block.getPosition(frame)
	}
	bkgM, ok := layer["backgroundColor"].(map[string]interface{})
	if ok {
		bkg := MapColor{Value: bkgM}
		block.BackgroundColor = bkg.colorRGBA()
	}
	style, ok := layer["style"].(map[string]interface{})
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
	childrenMaps, ok := layer["layers"].([]interface{})
	if ok {
		block.getChildren(childrenMaps)
	}
}

func (block *BlockCss) getChildren(childrenMaps []interface{}) {
	children := make([]interface{}, 0)
	for _, child := range childrenMaps {
		child, ok := child.(map[string]interface{})
		if ok {
			var block BlockCss
			block.cssBlock(child)
			children = append(children, block)
		}
	}
	block.Children = children
}

func (block *BlockCss) getPosition(frame map[string]interface{}) {
	block.Width = frame["width"].(float64)
	block.Height = frame["height"].(float64)
	block.Left = frame["x"].(float64)
	block.Top = frame["y"].(float64)
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
