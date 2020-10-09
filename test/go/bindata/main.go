package main

import "strings"

func main() {
	goMod := MustAssetString("go.mod")

	if !strings.HasPrefix(goMod, "module github.com/sagikazarmark/mypleasings") {
		println(goMod)
		panic("expected a go.mod file")
	}

	println("ok")
}
