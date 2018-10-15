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
				a, err := getArtboard(&lMap)
				if err != nil {
					return err
				}
				getLayers(&a.Layers)
				(*layers)[index] = a
			case "group":
				g, err := getGroup(&lMap)
				if err != nil {
					return err
				}
				getLayers(&g.Layers)
				(*layers)[index] = g
			case "shapeGroup":
				sG, err := getShapeGroup(&lMap)
				if err != nil {
					return err
				}
				getLayers(&sG.Layers)
				(*layers)[index] = sG
			case "text":
				t, err := getText(&lMap)
				if err != nil {
					return err
				}
				(*layers)[index] = t
			case "symbolInstance":
				sI, err := getSymbolInstance(&lMap)
				if err != nil {
					return err
				}
				(*layers)[index] = sI
			case "symbolMaster":
				sM, err := getSymbolMaster(&lMap)
				if err != nil {
					return err
				}
				getLayers(&sM.Layers)
				(*layers)[index] = sM
			}
		} else {
			return errors.New("not type is map")
		}
	}
	return nil
}

func getArtboard(layer *map[string]interface{}) (Artboard, error) {
	var result Artboard
	lByte, _ := json.Marshal(layer)
	err := json.Unmarshal(lByte, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func getGroup(layer *map[string]interface{}) (Group, error) {
	var result Group
	lByte, _ := json.Marshal(layer)
	err := json.Unmarshal(lByte, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func getShapeGroup(layer *map[string]interface{}) (ShapeGroup, error) {
	var result ShapeGroup
	lByte, _ := json.Marshal(layer)
	err := json.Unmarshal(lByte, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func getText(layer *map[string]interface{}) (Text, error) {
	var result Text
	lByte, _ := json.Marshal(layer)
	err := json.Unmarshal(lByte, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func getSymbolInstance(layer *map[string]interface{}) (SymbolInstance, error) {
	var result SymbolInstance
	lByte, _ := json.Marshal(layer)
	err := json.Unmarshal(lByte, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func getSymbolMaster(layer *map[string]interface{}) (SymbolMaster, error) {
	var result SymbolMaster
	lByte, _ := json.Marshal(layer)
	err := json.Unmarshal(lByte, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
