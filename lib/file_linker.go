package lib

import (
	"log"
	"os"
)

type FileLinker struct {
	Path string // relative path, like "./"
	SourceFiles []SourceFile
}

func CreateFileLinker(path string) FileLinker {
	return FileLinker{Path: path}
}

func (linker FileLinker) WriteExplodedFiles(targetDirectory string) {
	for i := range linker.SourceFiles {
		sourceFile := linker.SourceFiles[i]

		_ = os.MkdirAll(targetDirectory, os.ModePerm)

		file, err := os.OpenFile(targetDirectory + sourceFile.Name, os.O_RDWR | os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		_, err = file.Write([]byte(sourceFile.Content))
		if err != nil {
			log.Fatal(err)
		}
	}
}