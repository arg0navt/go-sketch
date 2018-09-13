package main

import (
	"./goSketch"
)

func main() {
	err := goSketch.GetFiles("./unsplash-app-creativepox.sketch", "dir")
	if err != nil {
		panic(err)
	}
}
