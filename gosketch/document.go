package gosketch

// Document document.json
type Document struct {
	ObjecttID              string `json:"Do_objectID"`
	Assets                 DocumentAssets
	ColorSpace             int
	CurrentPageIndex       int
	EnableLayerInteraction bool
	EnableSliceInteraction bool
	Pages                  []DocumentPage
}

// DocumentAssets documents.json "assets"
type DocumentAssets struct {
	Class  string `json:"_class"`
	Colors []DocumentColor
}

// DocumentColor documents.json "assets > colors"
type DocumentColor struct {
	Class string `json:"_class"`
	Alpha int
	Blue  int
	Green int
	Red   int
}

// DocumentPage document.json "pages"
type DocumentPage struct {
	Class     string `json:"_class"`
	RefClsass string `json:"_ref_class"`
	Ref       string `json:"_ref"`
}
