package gosketch

//Page jsons from folder pages/
type Page struct {
	Class                 string                   `json:"_class"`
	ObjectID              string                   `json:"do_objectID"`
	Frame                 PageFrame                `json:"frame"`
	IsFlippedHorizontal   bool                     `json:"isFlippedHorizontal"`
	IsFlippedVertical     bool                     `json:"isFlippedVertical"`
	IsLocked              bool                     `json:"isLocked"`
	IsVisible             bool                     `json:"isVisible"`
	LayerListExpandedType int                      `json:"layerListExpandedType"`
	Name                  string                   `json:"name"`
	NameIsFixed           bool                     `json:"nameIsFixed"`
	ResizingConstraint    int                      `json:"resizingConstraint"`
	ResizingType          int                      `json:"resizingType"`
	Rotation              int                      `json:"rotation"`
	ShouldBreakMaskChain  bool                     `json:"shouldBreakMaskChain"`
	Style                 PageStyle                `json:"style"`
	HasClickThrough       bool                     `json:"hasClickThrough"`
	Layers                []map[string]interface{} `json:"layers"`
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
