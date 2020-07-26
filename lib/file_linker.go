package lib

type FileLinker struct {
	Path string // relative path, like "./"
	SourceFiles []SourceFile
}

func CreateFileLinker(path string) FileLinker {
	return FileLinker{Path: path}
}