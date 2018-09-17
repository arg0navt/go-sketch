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

type Color struct {
	Alpha int
	Blue  int
	Green int
	Red   int
}

type Border struct {
	IsEnabled bool
	Color     Color
	FillType  int
	Position  int
	Thickness int
}

type GradientStop struct {
	Color    Color
	Position int
}

type Gradient struct {
	ElipseLength          int
	From                  string
	GradientType          int
	ShouldSmoothenOpacity int
	Stops                 []string
	to                    string
}

type GraphicsContextSettings struct {
	BlendMode int
	Opacity   int
}

type InnerShadow struct {
	IsEnabled       bool
	BlurRadius      int
	Color           Color
	ContextSettings GraphicsContextSettings
	OffsetX         int
	offsetY         int
	Spread          int
}

type Fill struct {
	IsEnabled        bool
	Color            Color
	FillType         int
	Gradient         Gradient
	NoiseIndex       int
	NoiseIntensity   int
	PatternFillType  int
	PatternTileScale int
}

type Shadow struct {
	IsEnabled       bool
	BlurRadius      int
	Color           Color
	ContextSettings GraphicsContextSettings
	OffsetX         int
	OffsetY         int
	Spread          int
}

type Blur struct {
	IsEnabled   bool
	Center      string
	MotionAngle int
	Radius      int
	Type        int
}

type Rect struct {
	ConstrainProportions bool
	Height               int
	Width                int
	X                    int
	Y                    int
}

type EncodedAttributes struct {
	NSKern                          int
	MSAttributedStringFontAttribute map[string]string
	NSParagraphStyle                map[string]string
	NSColor                         map[string]string
}

type TextStyle struct {
	EncodedAttributes EncodedAttributes
}

type BorderOptions struct {
	DoObjectID    string
	IsEnabled     bool
	LineCapStyle  int
	LineJoinStyle int
}

type ColorControls struct {
	IsEnabled  bool
	Brightness int
	Contrast   int
	Hue        int
	Saturation int
}

type Style struct {
	Blur                []Blur
	Borders             []Border
	BorderOptions       BorderOptions
	ContextSettings     GraphicsContextSettings
	ColorControls       ColorControls
	EndDecorationType   int
	Folls               []Fill
	InnerShadows        []InnerShadow
	MiterLimit          int
	Shadows             []Shadow
	SharedObjectID      string
	StartDecorationType int
	TextStyle           TextStyle
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
