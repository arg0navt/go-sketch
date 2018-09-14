package gosketch

type Document struct {
	Class                  string         `json:"_class"`
	ObjecttID              string         `json:"do_objectID"`
	Assets                 DocumentAssets `json:"assets"`
	ColorSpace             int64          `json:"colorSpace"`
	CurrentPageIndex       int64          `json:"currentPageIndex"`
	EnableLayerInteraction bool           `json:"enableLayerInteraction"`
	EnableSliceInteraction bool           `json:"enableSliceInteraction"`
	Pages                  []DocumentPage `json:"pages"`
}

type DocumentAssets struct {
	Class  string          `json:"_class"`
	Colors []DocumentColor `json:"colors"`
}

type DocumentColor struct {
	Class string `json:"_class"`
	Alpha int64  `json:"alpha"`
	Blue  int    `json:"blue"`
	Green int    `json:"green"`
	Red   int    `json:"red"`
}

type DocumentPage struct {
	Class     string `json:"_class"`
	RefClsass string `json:"_ref_class"`
	Ref       string `json:"_ref"`
}
