package gosketch

import (
	"encoding/json"
	"io/ioutil"
)

//SketchFile informarions of sketch file
type SketchFile struct {
	Document Document
	Meta     Meta
}

// Read : It's takes information about sketch and returns json map[]
func Read(src string) (*SketchFile, error) {
	var all SketchFile
	r, err := unzip(src)
	if err != nil {
		return &all, err
	}

	for _, f := range r.File {
		if f.Name == "meta.json" || f.Name == "document.json" {
			file, err := f.Open()
			if err != nil {
				return &all, err
			}
			defer file.Close()
			byteValue, err := ioutil.ReadAll(file)
			if err != nil {
				return &all, err
			}
			if f.Name == "meta.json" {
				json.Unmarshal(byteValue, &all.Document)
			}
			if f.Name == "document.json" {
				json.Unmarshal(byteValue, &all.Meta)
			}
		}
	}
	return &all, nil
}
