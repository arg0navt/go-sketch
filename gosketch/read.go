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
				go getPageLayers(jsonPage)
				// pagesMap[keyName] = jsonPage
			}
		}
	}
	all.Pages = pagesMap
	json.NewEncoder(w).Encode(&all)
}

func getLayers(layers []map[string]interface{}) {
	for index, layer := range layers {
		switch layer["_class"] {
		case "artboard":
			a := getArtboard(&layer)
			layers[index] = a
		case "group":
			g := getGroup(lMap)
			page.Layers[index] = g
		case "text":
			t := getText(lMap)
			page.Layers[index] = t
		case "symbolInstance":
			s := getSymbol(lMap)
			page.Layers[index] = s
		}
		// l, _ := json.Marshal(layer)
		// err2 := json.Unmarshal(l)
		// if err2 != nil {
		// 	panic(err2)
		// }
	}
	// var artboatd Artboard

	// 	getArtboard(&artboatd)
	// }
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

// func getArtboard(a *Artboard) {
// 	for index, layer := range a.Layers {
//
// 		if ok {

// 			switch l["_class"] {
// 			case "group":
// 				var dst Group
// 				l, _ := json.Marshal(layer)
// 				err := json.Unmarshal(l, &dst)
// 				if err != nil {
// 					log.Fatalln("error:", err)
// 				}
// 				a.Layers[index] = dst
// 				getLayers(&dst.Layers)
// 			}

// 		}
// 	}
// }

// func getLayers(l *[]interface{}) {
// 	for index, layer := range *l {
// 		lMap, ok := layer.(map[string]interface{})
// 		if ok {
// 			switch lMap["_class"] {
// 			case "group":
// 				group := getGroup(&lMap)
// 				(*l)[index] = group
// 				getLayers(&group.Layers)
// 			case "text":
// 				text := getText(&lMap)
// 				(*l)[index] = text
// 			case "symbolInstance":
// 				symbol := getSymbol(&lMap)
// 				(*l)[index] = symbol
// 			}
// 		}

// 	}
// }

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
