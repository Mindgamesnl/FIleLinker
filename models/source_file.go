package models

type SourceFile struct {
	Name string // the name of the file, may include relative paths and file extensions
	Content string // file content as represented by a string
}
