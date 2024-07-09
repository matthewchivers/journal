package editor

type Editor interface {
	OpenFile(filePath string) error
}
