package main

import (
	"fmt"

	"./gosketch"
)

func main() {
	i, err := gosketch.Read("./progressive-web-app-onboarding-richcullen.sketch")
	if err != nil {
		panic(err)
	}
	fmt.Println(i)
}
