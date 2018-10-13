package gosketch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

//SketchFile informarions of sketch file
type SketchFile struct {
	Document Document
	Meta     Meta
	Pages    map[string]Page
}

// Read : It's takes information about sketch and returns json map[]
func Read(w http.ResponseWriter, r *http.Request) {
	var all SketchFile
	re, err := unzip("./progressive-web-app-onboarding-richcullen.sketch")
	if err != nil {
		panic(err)
	}
	pagesMap := make(map[string]Page)
	for _, f := range re.File {
		formatJSON := strings.HasSuffix(f.Name, ".json")
		if formatJSON && f.Name != "user.json" {
			file, err := f.Open()
			if err != nil {
				panic(err)
			}
			defer file.Close()
			byteValue, err := ioutil.ReadAll(file)
			if err != nil {
				panic(err)
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
	json.NewEncoder(w).Encode(&all)
}
