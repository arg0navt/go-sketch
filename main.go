package main

import (
	"./gosketch"
)

func main() {
	i, err := gosketch.Read("./progressive-web-app-onboarding-richcullen.sketch")
	if err != nil {
		panic(err)
	}
	i.GetCSS("41CC057E-153E-4215-A787-8105A6BE3DE6")
}
