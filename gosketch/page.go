package gosketch

//Page jsons from folder pages/
type Page struct {
	ObjectID              string `json:"Do_objectID"`
	Frame                 PageFrame
	IsFlippedHorizontal   bool
	IsFlippedVertical     bool
	IsLocked              bool
	IsVisible             bool
	LayerListExpandedType int
	Name                  string
	NameIsFixed           bool
	ResizingConstraint    int
	ResizingType          int
	Rotation              int
	ShouldBreakMaskChain  bool
	Style                 PageStyle
	HasClickThrough       bool
	Layers                []interface{}
}

//PageFrame jsons from folder pages/ "frame"
type PageFrame struct {
	ConstrainProportions bool
	Height               int
	Width                int
	X                    int
	Y                    int
}

//PageStyle jsons from folder pages/ "style"
type PageStyle struct {
	EndDecorationType   int
	MiterLimit          int
	StartDecorationType int
}
