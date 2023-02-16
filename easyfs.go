package easyfs

import (
	"os"
	"testing/fstest"
)

type EasyFS struct {
	fstest.MapFS
}

// returns a fresh new filesystem
func NewFS() EasyFS {
	return EasyFS{fstest.MapFS{}}
}

func (m EasyFS) AddFile(name, content string) error {
	if name[0] == '/' {
		name = name[1:] // FS filesystem in go cannot start with /
	}
	m.MapFS[name] = &fstest.MapFile{
		Data: []byte(content),
	}
	return nil
}

func (m EasyFS) AddDir(name string) error {
	m.MapFS[name] = &fstest.MapFile{
		Mode: os.ModeDir,
	}
	return nil
}

func (m EasyFS) AddZip(name string, content []byte) error {
	m.MapFS[name] = &fstest.MapFile{
		Data: content,
	}
	return nil
}
