package easyfs

import (
	"io/fs"
	"os"
	"testing/fstest"
)

type FileMode = fs.FileMode

type EasyFS struct {
	fstest.MapFS
}

// returns a fresh new filesystem
func NewFS() EasyFS {
	return EasyFS{fstest.MapFS{}}
}

// deprecated
//func (m EasyFS) AddFile(name, content string) error {
//	return m.WriteFile(name, []byte(content), 0666)
//}

// depreciated
//func (m EasyFS) AddDir(name string) error {
//	return m.AddDir(name)
//}

func (m EasyFS) AddZip(name string, content []byte) error {
	m.MapFS[name] = &fstest.MapFile{
		Data: content,
	}
	return nil
}

func (m EasyFS) Mkdir(name string) error {
	m.MapFS[name] = &fstest.MapFile{
		Mode: os.ModeDir,
	}
	return nil
}

// WriteFile writes data to a file named by filename. perm is not used but cn be set to
func (m EasyFS) WriteFile(name string, content []byte, perm FileMode) error {
	//perm is not implimented
	if name[0] == '/' {
		name = name[1:] // FS filesystem in go cannot start with /
	}
	m.MapFS[name] = &fstest.MapFile{
		Data: content,
	}
	return nil

}
