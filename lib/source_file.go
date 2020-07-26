package lib

import (
	"io/ioutil"
	"strings"
)

var commentPrefixes = [2]string{"#", "//"}

type SourceFile struct {
	Name string // the name of the file, may include relative paths and file extensions
	Content string // file content as represented by a string
}

func (linker FileLinker) ReadFromRootFile(rootFileName string) {
	var content string
	parsedContent, _ := ioutil.ReadFile(rootFileName)
	content = string(parsedContent)
	currentLine := 0

	var isWriting = false
	var currentTarget SourceFile

	lineScanner:
	for _, line := range strings.Split(strings.TrimSuffix(content, "\n"), "\n") {
		// check for all types
		for i := range commentPrefixes {
			prefix := commentPrefixes[i]

			if strings.HasPrefix(line, prefix + "FL:START:") {
				sourceFileName := strings.Replace(line, prefix + "FL:START:", "", -1)
				currentTarget = SourceFile{Name: sourceFileName, Content: ""}
				isWriting = true
				continue lineScanner
			}

			if strings.HasPrefix(line, prefix + "FL:END:") {
				linker.SourceFiles = append(linker.SourceFiles, currentTarget)
				isWriting = false
				continue lineScanner
			}
		}

		if isWriting {
			currentTarget.Content += line + "\n"
		}

		currentLine++
	}
}