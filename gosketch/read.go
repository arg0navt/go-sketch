package gosketch

import (
	"encoding/json"
	"errors"
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
				getLayers(&jsonPage.Layers)
				pagesMap[keyName] = jsonPage
			}
		}
	}
	all.Pages = pagesMap
	return &all, nil
}

func getLayers(layers *[]interface{}) error {
	for index, layer := range *layers {
		lMap, ok := layer.(map[string]interface{})
		if ok {
			switch lMap["_class"] {
			case "artboard":
				getArtboard(&lMap, &(*layers)[index])
			case "group":
				getGroup(&lMap, &(*layers)[index])
			case "shapeGroup":
				getShapeGroup(&lMap, &(*layers)[index])
			case "text":
				getText(&lMap, &(*layers)[index])
			case "symbolInstance":
				getSymbolInstance(&lMap, &(*layers)[index])
			case "symbolMaster":
				getSymbolMaster(&lMap, &(*layers)[index])
			}
		} else {
			return errors.New("not type is map")
		}
	}
	return nil
}

func getArtboard(layer *map[string]interface{}, result *interface{}) {
	var r Artboard
	lByte, _ := json.Marshal(layer)
	err := json.Unmarshal(lByte, &r)
	if err != nil {
		panic(err)
	}
	*result = r
}

func getGroup(layer *map[string]interface{}, result *interface{}) {
	var r Group
	lByte, _ := json.Marshal(layer)
	err := json.Unmarshal(lByte, &r)
	if err != nil {
		panic(err)
	}
	*result = r
}

func getShapeGroup(layer *map[string]interface{}, result *interface{}) {
	var r ShapeGroup
	lByte, _ := json.Marshal(layer)
	err := json.Unmarshal(lByte, &r)
	if err != nil {
		panic(err)
	}
	*result = r
}

func getText(layer *map[string]interface{}, result *interface{}) {
	var r Text
	lByte, _ := json.Marshal(layer)
	err := json.Unmarshal(lByte, &r)
	if err != nil {
		panic(err)
	}
	*result = r
}

func getSymbolInstance(layer *map[string]interface{}, result *interface{}) {
	var r SymbolInstance
	lByte, _ := json.Marshal(layer)
	err := json.Unmarshal(lByte, &r)
	if err != nil {
		panic(err)
	}
	*result = r
}

func getSymbolMaster(layer *map[string]interface{}, result *interface{}) {
	var r SymbolMaster
	lByte, _ := json.Marshal(layer)
	err := json.Unmarshal(lByte, &r)
	if err != nil {
		panic(err)
	}
	*result = r
}
