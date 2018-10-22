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

type Color struct {
	Alpha float64
	Blue  float64
	Green float64
	Red   float64
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
	Opacity   float64
}

type InnerShadow struct {
	IsEnabled       bool
	BlurRadius      int
	Color           Color
	ContextSettings GraphicsContextSettings
	OffsetX         int
	OffsetY         int
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
	MotionAngle float64
	Radius      float64
	Type        float64
}

type Rect struct {
	ConstrainProportions bool
	Height               float64
	Width                float64
	X                    float64
	Y                    float64
}

type EncodedAttributes struct {
	NSKern float64
}

type TextStyle struct {
	EncodedAttributes EncodedAttributes
}

type BorderOptions struct {
	DoObjectID    string `json:"Do_objectID"`
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
	Blur                Blur
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
	DoObjectID string `json:"Do_objectID"`
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
	DoObjectID   string `json:"Do_objectID"`
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
	Class                             string `json:"_class"`
	DoObjectID                        string `json:"Do_objectID"`
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
	Class                 string `json:"_class"`
	DoObjectID            string `json:"Do_objectID"`
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
	DoObjectID            string `json:"do_objectID"`
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
	Class                          string `json:"_class"`
	DoObjectID                     string `json:"Do_objectID"`
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
	DoObjectID            string `json:"do_objectID"`
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
	Class                          string `json:"_class"`
	DoObjectID                     string `json:"Do_objectID"`
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
	SymbolID                       string
	VerticalSpacing                int
}

type Group struct {
	Class                 string `json:"_class"`
	DoObjectID            string `json:"Do_objectID"`
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
	Class                         string `json:"_class"`
	DoObjectID                    string `json:"Do_objectID"`
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
	FixedRadius                   float64
	HasConvertedToNewRoundCorners bool
}

type Oval struct {
	Class                 string `json:"_class"`
	DoObjectID            string `json:"Do_objectID"`
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
	Class                            string `json:"_class"`
	BackgroundColor                  Color
	DoObjectID                       string `json:"Do_objectID"`
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
