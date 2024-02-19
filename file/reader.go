package file

import (
	"io/fs"
	"os"
	"path/filepath"
)

type Scanner struct {
	Path string
}

var scanner *Scanner

func DefaultScanner() *Scanner {
	if scanner == nil {
		scanner = NewScanner("")
	}
	return scanner
}

func NewScanner(path string) *Scanner {
	if path == "" {
		path, _ = os.Getwd()
	}
	return &Scanner{
		Path: path,
	}
}

func (s *Scanner) Ls() []fs.FileInfo {
	files := make([]fs.FileInfo, 0)
	filepath.Walk(s.Path, func(path string, info fs.FileInfo, err error) error {
		if info.Mode().IsDir() {
			return nil
		}
		files = append(files, info)
		return nil
	})
	return files
}
