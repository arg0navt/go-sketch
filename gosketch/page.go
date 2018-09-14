package gosketch

//Page jsons from folder pages/
type Page struct {
	Class                 string    `json:"_class"`
	ObjectID              string    `json:"do_objectID"`
	Frame                 PageFrame `json:"frame"`
	IsFlippedHorizontal   bool      `json:"isFlippedHorizontal"`
	IsFlippedVertical     bool      `json:"isFlippedVertical"`
	IsLocked              bool      `json:"isLocked"`
	IsVisible             bool      `json:"isVisible"`
	LayerListExpandedType int64     `json:"layerListExpandedType"`
	Name                  string    `json:"name"`
	NameIsFixed           bool      `json:"nameIsFixed"`
	ResizingConstraint    int64     `json:"resizingConstraint"`
	ResizingType          int64     `json:"resizingType"`
	Rotation              int       `json:"rotation"`
	ShouldBreakMaskChain  bool      `json:"shouldBreakMaskChain"`
	Style                 PageStyle `json:"style"`
	HasClickThrough       bool      `json:"hasClickThrough"`
	Layers                map[string]interface{}
}

//PageFrame jsons from folder pages/ "frame"
type PageFrame struct {
	Class                string `json:"_class"`
	ConstrainProportions bool   `json:"constrainProportions"`
	Height               int64  `json:"height"`
	Width                int64  `json:"width"`
	X                    int64  `json:"x"`
	Y                    int64  `json:"y"`
}

//PageStyle jsons from folder pages/ "style"
type PageStyle struct {
	Class               string `json:"_class"`
	EndDecorationType   int64  `json:"endDecorationType"`
	MiterLimit          int64  `json:"miterLimit"`
	StartDecorationType int64  `json:"startDecorationType"`
}
