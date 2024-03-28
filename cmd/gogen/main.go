package main

import (
	"generator/internal/generator"
)

func main() {
	g := generator.New()
	if err := g.Generate(); err != nil {
		panic(err)
	}
}
