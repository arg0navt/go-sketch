package gosketch

// Meta meta.json
type Meta struct {
	Commit               string
	PagesAndArtboards    map[string]MetaPage
	Version              int
	Fonts                []string
	CompatibilityVersion string
	App                  string
	Autosaved            int
	Variant              string
	Created              MetaCreated
	AppVersion           string
	Build                int
}

// MetaPage meta.json "pagesAndArtboards"
type MetaPage struct {
	Name      string
	Artboards map[string]MetaArtboard
}

// MetaArtboard meta.json "pagesAndArtboards > artboards"
type MetaArtboard struct {
	Name string
}

// MetaCreated meta.json "created"
type MetaCreated struct {
	Commit               string
	AppVersion           string
	Build                int
	App                  string
	CompatibilityVersion int
	Version              int
	Variant              string
}

// PagesList return ID pages
func (s *SketchFile) PagesList() []string {
	var result []string
	for key := range s.Meta.PagesAndArtboards {
		result = append(result, key)
	}
	return result
}
