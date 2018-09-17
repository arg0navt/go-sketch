package gosketch

import "fmt"

//Page jsons from folder pages/
type Page struct {
	Class                 string        `json:"_class"`
	ObjectID              string        `json:"do_objectID"`
	Frame                 PageFrame     `json:"frame"`
	IsFlippedHorizontal   bool          `json:"isFlippedHorizontal"`
	IsFlippedVertical     bool          `json:"isFlippedVertical"`
	IsLocked              bool          `json:"isLocked"`
	IsVisible             bool          `json:"isVisible"`
	LayerListExpandedType int           `json:"layerListExpandedType"`
	Name                  string        `json:"name"`
	NameIsFixed           bool          `json:"nameIsFixed"`
	ResizingConstraint    int           `json:"resizingConstraint"`
	ResizingType          int           `json:"resizingType"`
	Rotation              int           `json:"rotation"`
	ShouldBreakMaskChain  bool          `json:"shouldBreakMaskChain"`
	Style                 PageStyle     `json:"style"`
	HasClickThrough       bool          `json:"hasClickThrough"`
	Layers                []interface{} `json:"layers"`
}

//PageFrame jsons from folder pages/ "frame"
type PageFrame struct {
	Class                string `json:"_class"`
	ConstrainProportions bool   `json:"constrainProportions"`
	Height               int    `json:"height"`
	Width                int    `json:"width"`
	X                    int    `json:"x"`
	Y                    int    `json:"y"`
}

//PageStyle jsons from folder pages/ "style"
type PageStyle struct {
	Class               string `json:"_class"`
	EndDecorationType   int    `json:"endDecorationType"`
	MiterLimit          int    `json:"miterLimit"`
	StartDecorationType int    `json:"startDecorationType"`
}

type Layer struct {
	Name   string
	Layers interface{}
}

// GetLyersPage get layers elements by page
func (s *SketchFile) GetLyersPage(page string) []interface{} {
	return s.Pages[page].Layers
}

func getLayer(l interface{}) {
	mapLayer, ok := l.(map[string]interface{})
	if ok {
		fmt.Println(mapLayer["name"])
	}
	children, okCh := mapLayer["layers"].([]interface{})
	if okCh {
		for _, childrenLayer := range children {
			getLayer(childrenLayer)
		}
	}
}

// GetCSS get style css by layrs page
func (s *SketchFile) GetCSS(pageID string) {
	page := s.Pages[pageID]
	for _, l := range page.Layers {
		getLayer(l)
	}
}
