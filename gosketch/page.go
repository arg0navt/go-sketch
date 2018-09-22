package gosketch

import (
	"fmt"
)

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
	Height               float64
	Width                float64
	X                    float64
	Y                    float64
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

type SharedStyle struct {
	DoObjectID string
	Name       string
	Value      Style
}

type ExportFormat struct {
	AbsoluteSize     int
	FileFormat       string
	Name             string
	NamingScheme     int
	Scale            int
	VisibleScaleType int
}

type ExportOptions struct {
	ExportFormats []ExportFormat
	LayerOptions  float64
	ShouldTrim    bool
}

type SharedStyleContainer struct {
	Objects []SharedStyle
}

type MSJSONFileReference struct {
	RefClass string
	Red      string
}

type MSAttributedString struct {
	ArchivedAttributedString map[string]string
}

type CurvePoint struct {
	DoObjectID   string
	CornerRadius int
	СurveFrom    string
	CurveMode    int
	CurveTo      string
	HasCurveFrom bool
	HasCurveTo   bool
	Point        string
}

type RulerData struct {
	Base int
}

type Text struct {
	DoObjectID                        string
	ExportOptions                     ExportOptions
	Frame                             Rect
	IsFlippedVertical                 bool
	IsFlippedHorizontal               bool
	IsLocked                          bool
	IsVisible                         bool
	LayerListExpandedType             int
	Name                              string
	NameIsFixed                       bool
	OriginalObjectID                  int
	ResizingType                      int
	Rotation                          int
	ShouldBreakMaskChain              bool
	Style                             Style
	AttributedString                  MSAttributedString
	AutomaticallyDrawOnUnderlyingPath bool
	DontSynchroniseWithSymbol         bool
	GlyphBounds                       string
	HeightIsClipped                   bool
	LineSpacingBehaviour              int
	TextBehaviour                     int
}

type ShapeGroup struct {
	DoObjectID            string
	ExportOptions         ExportOptions
	Frame                 Rect
	IsFlippedVertical     bool
	IsFlippedHorizontal   bool
	IsLocked              bool
	IsVisible             bool
	LayerListExpandedType int
	Name                  string
	NameIsFixed           bool
	OriginalObjectID      string
	ResizingType          int
	Rotation              int
	ShouldBreakMaskChain  bool
	Style                 Style
	HasClickThrough       bool
	Layers                []interface{}
	ClippingMaskMode      int
	HasClippingMask       bool
	WindingRule           int
}

type Path struct {
	IsClosed bool
	Points   []CurvePoint
}

type ShapePath struct {
	DoObjectID            string
	ExportOptions         ExportOptions
	Frame                 Rect
	IsFlippedVertical     bool
	IsFlippedHorizontal   bool
	IsLocked              bool
	IsVisible             bool
	LayerListExpandedType int
	Name                  string
	NameIsFixed           bool
	ResizingType          int
	Rotation              int
	ShouldBreakMaskChain  bool
	BooleanOperation      int
	Edited                bool
	Path                  Path
}

type Artboard struct {
	DoObjectID                     string
	ExportOptions                  ExportOptions
	Frame                          Rect
	IsFlippedVertical              bool
	IsFlippedHorizontal            bool
	IsLocked                       bool
	IsVisible                      bool
	LayerListExpandedType          float64
	Name                           string
	NameIsFixed                    bool
	ResizeType                     float64
	Rotation                       float64
	ShouldBreakMaskChain           bool
	Style                          Style
	HasClickThrough                bool
	Layers                         []interface{}
	BackgroundColor                Color
	HasBackgroundColor             bool
	HorizontalRulerData            RulerData
	IncludeBackgroundColorInExport bool
	IncludeInCloudUpload           bool
	VerticalRulerData              RulerData
}

type Bitmap struct {
	DoObjectID            string
	ExportOptions         ExportOptions
	Frame                 Rect
	IsFlippedVertical     bool
	IsFlippedHorizontal   bool
	IsLocked              bool
	IsVisible             bool
	LayerListExpandedType int
	Name                  string
	NameIsFixed           bool
	ResizeType            int
	Rotation              int
	ShouldBreakMaskChain  bool
	Style                 Style
	ClippingMask          string
	FillReplacesImage     bool
	Image                 MSJSONFileReference
	NineSliceCenterRect   string
	NineSliceScale        string
}

type SymbolInstance struct {
	DoObjectID                     string
	ExportOptions                  ExportOptions
	Frame                          Rect
	IsFlippedVertical              bool
	IsFlippedHorizontal            bool
	IsLocked                       bool
	IsVisible                      bool
	LayerListExpandedType          int
	Name                           string
	NameIsFixed                    bool
	ResizeType                     int
	Rotation                       int
	ShouldBreakMaskChain           bool
	Style                          Style
	HorizontalSpacing              int
	MasterInfluenceEdgeMaxXPadding int
	MasterInfluenceEdgeMaxYPadding int
	MasterInfluenceEdgeMinXPadding int
	MasterInfluenceEdgeMinYPadding int
	SymbolID                       int
	VerticalSpacing                int
}

type Group struct {
	DoObjectID            string
	ExportOptions         ExportOptions
	Frame                 Rect
	IsFlippedVertical     bool
	IsFlippedHorizontal   bool
	IsLocked              bool
	IsVisible             bool
	LayerListExpandedType int
	Name                  string
	NameIsFixed           bool
	OriginalObjectID      string
	ResizeType            int
	Rotation              int
	ShouldBreakMaskChain  bool
	HasClickThrough       bool
	Layers                []interface{}
}

type Rectangle struct {
	DoObjectID                    string
	ExportOptions                 ExportOptions
	Frame                         Rect
	IsFlippedVertical             bool
	IsFlippedHorizontal           bool
	IsLocked                      bool
	IsVisible                     bool
	LayerListExpandedType         int
	Name                          string
	NameIsFixed                   bool
	ResizeType                    int
	Rotation                      int
	ShouldBreakMaskChain          bool
	BooleanOperation              int
	Edited                        bool
	Path                          Path
	FixedRadius                   int
	HasConvertedToNewRoundCorners bool
}

type Oval struct {
	DoObjectID            string
	ExportOptions         ExportOptions
	Frame                 Rect
	IsFlippedVertical     bool
	IsFlippedHorizontal   bool
	IsLocked              bool
	IsVisible             bool
	LayerListExpandedType int
	Name                  string
	NameIsFixed           bool
	ResizeType            int
	Rotation              int
	ShouldBreakMaskChain  bool
	BooleanOperation      int
	Edited                bool
	Path                  Path
}

type SymbolMaster struct {
	BackgroundColor                  Color
	DoObjectID                       string
	ExportOptions                    ExportOptions
	Frame                            Rect
	HasBackgroundColor               bool
	HasClickThrough                  bool
	HorizontalRulerData              RulerData
	IncludeBackgroundColorInExport   bool
	IncludeBackgroundColorInInstance bool
	IncludeInCloudUpload             bool
	IsFlippedHorizontal              bool
	IsFlippedVertical                bool
	IsLocked                         bool
	IsVisible                        bool
	LayerListExpandedType            int
	Layers                           []interface{}
	Name                             string
	NameIsFixed                      bool
	Rotation                         int
	ShouldBreakMaskChain             bool
	Style                            Style
	SymbolID                         string
	VerticalRulerData                RulerData
}

// GetLyersPage get layers elements by page
func (s *SketchFile) GetLyersPage(page string) []interface{} {
	return s.Pages[page].Layers
}

func getLayer(l interface{}, count int) {
	countChildren := count
	mapLayer, ok := l.(map[string]interface{})
	if ok {
		switch className := mapLayer["_class"].(string); className {
		case "artboard":
			a := getArtboard(mapLayer)
			fmt.Println(a)
		}
	}
	children, okCh := mapLayer["layers"].([]interface{})
	if okCh {
		countChildren++
		for _, childrenLayer := range children {
			getLayer(childrenLayer, countChildren)
		}
	}
}

// GetCSS get style css by layrs page
func (s *SketchFile) GetCSS(pageID string) {
	page := s.Pages[pageID]
	for _, l := range page.Layers {
		getLayer(l, 0)
	}
}

func getArtboard(layer map[string]interface{}) Artboard {
	eo := getExportOptions(layer["exportOptions"])
	f := getFrame(layer["frame"])
	// s := getStyle(layer["style"])
	return Artboard{
		DoObjectID:            layer["do_objectID"].(string),
		ExportOptions:         eo,
		Frame:                 f,
		IsFlippedVertical:     layer["isFlippedVertical"].(bool),
		IsFlippedHorizontal:   layer["isFlippedHorizontal"].(bool),
		IsLocked:              layer["isLocked"].(bool),
		IsVisible:             layer["isVisible"].(bool),
		LayerListExpandedType: layer["layerListExpandedType"].(float64),
		Name:                 layer["name"].(string),
		NameIsFixed:          layer["nameIsFixed"].(bool),
		ResizeType:           layer["resizingType"].(float64),
		Rotation:             layer["rotation"].(float64),
		ShouldBreakMaskChain: layer["shouldBreakMaskChain"].(bool),
	}
}

func getExportOptions(eo interface{}) ExportOptions {
	l, ok := eo.(map[string]interface{})
	if ok {
		return ExportOptions{
			LayerOptions: l["layerOptions"].(float64),
			ShouldTrim:   l["shouldTrim"].(bool),
		}
	}
	return ExportOptions{}
}

func getFrame(f interface{}) Rect {
	l, ok := f.(map[string]interface{})
	if ok {
		return Rect{
			ConstrainProportions: l["constrainProportions"].(bool),
			Height:               l["height"].(float64),
			Width:                l["width"].(float64),
			X:                    l["x"].(float64),
			Y:                    l["y"].(float64),
		}
	}
	return Rect{}
}

// func getStyle(s interface{}) Style {

// }
