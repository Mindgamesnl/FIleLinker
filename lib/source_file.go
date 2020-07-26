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

func (linker FileLinker) ReadFromRootFile(rootFileName string) FileLinker {
	var content string
	parsedContent, _ := ioutil.ReadFile(linker.Path + rootFileName)
	content = string(parsedContent)
	currentLine := 0

	var isWriting = false
	var currentTarget SourceFile
	var rootFile SourceFile
	var usedNotePrefix string;

	lineScanner:
	for _, line := range strings.Split(strings.TrimSuffix(content, "\n"), "\n") {
		// first line? if so, check for a root note if its even a file
		if currentLine == 0 {
			foundRoot := false
			for i := range commentPrefixes {
				prefix := commentPrefixes[i]
				usedNotePrefix = prefix
				if strings.HasPrefix(line, prefix + "FL:ROOT:") {
					foundRoot = true
					rootFileName := strings.Replace(line, prefix + "FL:ROOT:", "", -1)
					rootFile = SourceFile{Name: linker.Path + rootFileName, Content: ""}
					currentLine++
					continue lineScanner
				}
			}

			if !foundRoot {
				rootFile = SourceFile{Name: linker.Path + "unknown", Content: content}
				currentLine++
				break lineScanner
			}
		}

		// check for all types
		for i := range commentPrefixes {
			prefix := commentPrefixes[i]

			if strings.HasPrefix(line, prefix + "FL:START:") {
				sourceFileName := strings.Replace(line, prefix + "FL:START:", "", -1)
				currentTarget = SourceFile{Name: linker.Path + sourceFileName, Content: ""}
				isWriting = true
				currentLine++
				continue lineScanner
			}

			if strings.HasPrefix(line, prefix + "FL:END") {
				linker.SourceFiles = append(linker.SourceFiles, currentTarget)
				isWriting = false
				currentLine++

				// re-add import to the original file
				nameWithoutRoot := strings.Replace(currentTarget.Name, linker.Path, "", -1)
				rootFile.Content += usedNotePrefix + "import<" + nameWithoutRoot + ">" + "\n"
				continue lineScanner
			}
		}

		if isWriting {
			currentTarget.Content += line + "\n"
		} else {
			rootFile.Content += line + "\n"
		}

		currentLine++
	}
	linker.SourceFiles = append(linker.SourceFiles, rootFile)
	return linker
}