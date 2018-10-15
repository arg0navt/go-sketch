package gosketch

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
				getLayers(&jsonPage.Layers)
				pagesMap[keyName] = jsonPage
			}
		}
	}
	all.Pages = pagesMap
	json.NewEncoder(w).Encode(&all)
}

func getLayers(layers *[]interface{}) {
	for index, layer := range *layers {
		lMap, ok := layer.(map[string]interface{})
		if ok {
			switch lMap["_class"] {
			case "artboard":
				a := getArtboard(&lMap)
				getLayers(&a.Layers)
				(*layers)[index] = a
			case "group":
				g := getGroup(&lMap)
				getLayers(&g.Layers)
				(*layers)[index] = g
			case "text":
				t := getText(&lMap)
				(*layers)[index] = t
			case "symbolInstance":
				s := getSymbol(&lMap)
				(*layers)[index] = s
			}
		}
	}
}

func getArtboard(layer *map[string]interface{}) Artboard {
	var result Artboard
	lByte, _ := json.Marshal(layer)
	err := json.Unmarshal(lByte, &result)
	if err != nil {
		log.Fatalln("error:", err)
	}
	return result
}

func getGroup(layer *map[string]interface{}) Group {
	var result Group
	lByte, _ := json.Marshal(layer)
	err := json.Unmarshal(lByte, &result)
	if err != nil {
		log.Fatalln("error:", err)
	}
	return result
}

func getText(layer *map[string]interface{}) Text {
	var result Text
	lByte, _ := json.Marshal(layer)
	err := json.Unmarshal(lByte, &result)
	if err != nil {
		log.Fatalln("error:", err)
	}
	return result
}

func getSymbol(layer *map[string]interface{}) SymbolInstance {
	var result SymbolInstance
	lByte, _ := json.Marshal(layer)
	err := json.Unmarshal(lByte, &result)
	if err != nil {
		log.Fatalln("error:", err)
	}
	return result
}
