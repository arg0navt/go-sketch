package gosketch

// Document document.json
type Document struct {
	Class                  string         `json:"_class"`
	ObjecttID              string         `json:"do_objectID"`
	Assets                 DocumentAssets `json:"assets"`
	ColorSpace             int            `json:"colorSpace"`
	CurrentPageIndex       int            `json:"currentPageIndex"`
	EnableLayerInteraction bool           `json:"enableLayerInteraction"`
	EnableSliceInteraction bool           `json:"enableSliceInteraction"`
	Pages                  []DocumentPage `json:"pages"`
}

// DocumentAssets documents.json "assets"
type DocumentAssets struct {
	Class  string          `json:"_class"`
	Colors []DocumentColor `json:"colors"`
}

// DocumentColor documents.json "assets > colors"
type DocumentColor struct {
	Class string `json:"_class"`
	Alpha int    `json:"alpha"`
	Blue  int    `json:"blue"`
	Green int    `json:"green"`
	Red   int    `json:"red"`
}

// DocumentPage document.json "pages"
type DocumentPage struct {
	Class     string `json:"_class"`
	RefClsass string `json:"_ref_class"`
	Ref       string `json:"_ref"`
}
