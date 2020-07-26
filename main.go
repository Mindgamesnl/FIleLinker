package main

import (
	"github.com/Mindgamesnl/FileLinker/lib"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	fileLinker := lib.CreateFileLinker("./")
	fileLinker = fileLinker.ReadFromRootFile("test_sources/example.js")
	fileLinker.WriteExplodedFiles("./out/")
	spew.Dump(fileLinker)


}
