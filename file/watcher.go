package file

import "os"

type FileWatcher struct {
}

type FileDescriptor struct {
}

func (w *FileWatcher) Watch(files map[string]os.FileInfo) {
	for file, fInfo := range files {
		info, err := os.Lstat(file)
		if err != nil {
			continue
		}
		if info.IsDir() {
			continue
		}
		// this file is symlink
		if info.Mode()&os.ModeSymlink > 0 {
			info, err = os.Stat(file)
		}
		//
		if fInfo.Size() >= info.Size() {

		} else if fInfo.Size() < info.Size() {

		}
		files[file] = info
	}
}
