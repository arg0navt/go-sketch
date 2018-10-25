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
	Borders         []string
	BoxShadow       string
	Children        []interface{}
	Font            interface{}
}

type Font struct {
	Size   float64
	Color  string
	Family string
	Weight float64
}

type MapColor struct {
	Value map[string]interface{}
}

type MapShadow struct {
	Value map[string]interface{}
}

type ProtocolWood struct {
	Item  BlockCss
	Index int
}

func (s *SketchFile) GetCSS(w http.ResponseWriter, r *http.Request) {
	result := make([]interface{}, 0)
	for key, page := range s.Pages {
		structureBranches := make([]interface{}, len(page.Layers))
		growBranch := make(chan ProtocolWood)
		countWoods := make(chan int)
		countChildren := len(page.Layers)
		countW := len(page.Layers)
		for index, item := range page.Layers {
			go cssBlock(item, index, growBranch, countWoods)
		}
		for countW > 0 {
			select {
			case newBranch := <-growBranch:
				structureBranches[newBranch.Index] = newBranch.Item
				countChildren = countChildren - 1
				if countChildren == 0 {
					result = append(result, PageCss{ID: key, Css: structureBranches})
				}
			case c := <-countWoods:
				countW = countW + c - 1
				fmt.Println(countW)
			}
		}

	}
	json.NewEncoder(w).Encode(result)
}

func cssBlock(layer map[string]interface{}, index int, growBranch chan<- ProtocolWood, countWoods chan<- int) {
	var block BlockCss
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
		borders, ok := style["borders"].([]interface{})
		if ok {
			go block.getBorders(borders)
		}
	}
	if layer["_class"] == "artboard" || layer["_class"] == "group" || layer["_class"] == "shapeGroup" || layer["_class"] == "symbolMaster" {
		block.Font = nil
	} else {
		block.Font = Font{}
	}
	childrenMaps, ok := layer["layers"].([]interface{})
	growBranch <- ProtocolWood{Item: block, Index: index}
	if ok {
		block.getChildren(childrenMaps, countWoods)
	}
}

func (block *BlockCss) getChildren(childrenMaps []interface{}, countWoods chan<- int) {
	structureBranches := make([]interface{}, len(childrenMaps))
	growBranch := make(chan ProtocolWood)
	count := len(childrenMaps)
	countWoods <- len(childrenMaps)
	for index, child := range childrenMaps {
		child, ok := child.(map[string]interface{})
		if ok {
			go cssBlock(child, index, growBranch, countWoods)
		}
	}
	for count > 0 {
		select {
		case newBranch := <-growBranch:
			structureBranches[newBranch.Index] = newBranch.Item
			count = count - 1
			if count == 0 {
				block.Children = structureBranches
			}
		}
	}
}

func (block *BlockCss) getBorders(borders []interface{}) {
	result := make([]string, 0)
	for _, border := range borders {
		border, ok := border.(map[string]interface{})
		if ok && border["isEnabled"].(bool) {
			width := strconv.Itoa(int(border["thickness"].(float64)))
			color, ok := border["color"].(map[string]interface{})
			if ok {
				color := MapColor{Value: color}
				colorString := color.colorRGBA()
				result = append(result, width+"px solid "+colorString) // TODO: only solid
			}
		}
	}
	if len(result) > 0 {
		block.Borders = result
	}
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
