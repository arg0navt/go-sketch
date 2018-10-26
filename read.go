package gosketch

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

//SketchFile informarions of sketch file
type SketchFile struct {
	Document Document
	Meta     Meta
	Pages    map[string]Page
}

// Read : It's takes information about sketch and returns json map[]
func Read(path string) (*SketchFile, error) {
	var all SketchFile
	re, err := unzip(path)
	if err != nil {
		return &all, err
	}
	pagesMap := make(map[string]Page)
	for _, f := range re.File {
		formatJSON := strings.HasSuffix(f.Name, ".json")
		if formatJSON && f.Name != "user.json" {
			file, err := f.Open()
			if err != nil {
				return &all, nil
			}
			defer file.Close()
			byteValue, err := ioutil.ReadAll(file)
			if err != nil {
				return &all, nil
			}
			if f.Name == "meta.json" {
				var jsonMeta Meta
				json.Unmarshal(byteValue, &jsonMeta)
				all.Meta = jsonMeta
			} else if f.Name == "document.json" {
				var jsonDocument Document
				json.Unmarshal(byteValue, &jsonDocument)
				all.Document = jsonDocument
			} else {
				var jsonPage Page
				json.Unmarshal(byteValue, &jsonPage)
				keyName := strings.TrimSuffix(f.Name, ".json")
				keyName = keyName[6:]
				pagesMap[keyName] = jsonPage
			}
		}
	}
	all.Pages = pagesMap
	return &all, nil
}
