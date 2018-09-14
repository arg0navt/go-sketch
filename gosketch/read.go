package gosketch

import (
	"encoding/json"
	"io/ioutil"
)

type ReadSketch interface {
	setDocument(map[string]interface{})
	setMeta(map[string]interface{})
	setUser(map[string]interface{})
	GetDocument() map[string]interface{}
	GetUser() map[string]interface{}
	GetMeta() map[string]interface{}
}

type SketchFile struct {
	Document map[string]interface{}
	User     map[string]interface{}
	Meta     map[string]interface{}
}

func (s *SketchFile) setDocument(j map[string]interface{}) {
	s.Document = j
}

func (s *SketchFile) setMeta(j map[string]interface{}) {
	s.Meta = j
}

func (s *SketchFile) setUser(j map[string]interface{}) {
	s.User = j
}

func (s *SketchFile) GetDocument() map[string]interface{} {
	return s.Document
}

func (s *SketchFile) GetUser() map[string]interface{} {
	return s.User
}

func (s *SketchFile) GetMeta() map[string]interface{} {
	return s.Meta
}

// GetJSON : It's takes information about sketch and returns json map[]
func Read(src string) (*ReadSketch, error) {
	var all ReadSketch = &SketchFile{}
	r, err := unzip(src)
	if err != nil {
		return &all, err
	}

	for _, f := range r.File {
		if f.Name == "meta.json" || f.Name == "document.json" || f.Name == "user.json" {
			file, err := f.Open()
			if err != nil {
				return &all, err
			}
			defer file.Close()
			byteValue, err := ioutil.ReadAll(file)
			if err != nil {
				return &all, err
			}
			jsonDocument := make(map[string]interface{})
			json.Unmarshal(byteValue, &jsonDocument)
			if f.Name == "meta.json" {
				all.setDocument(jsonDocument)
			}
			if f.Name == "document.json" {
				all.setMeta(jsonDocument)
			}
			if f.Name == "user.json" {
				all.setUser(jsonDocument)
			}
		}
	}
	return &all, nil
}
