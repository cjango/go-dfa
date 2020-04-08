package main

import (
	"dfa/core"
	"runtime"
)

func main() {
	types := []string{"porn", "contraband", "politics", "reviles", "wartering"}

	runtime.GOMAXPROCS(runtime.NumCPU())
	core.LoadKeywords(types)
	//server.Start("9528", types)
}
