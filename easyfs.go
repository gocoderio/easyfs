package easyfs

import (
	"io/fs"
	"os"
	"reflect"
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

// returns an existing filesystem if there is one
func GetFS() EasyFS {
	if checkGlobalVar() {
		return easyFS //if there is a global var, return the existing filesystem
	}
	return EasyFS{fstest.MapFS{}} // if not return a new one
}
func checkGlobalVar() bool {
	var zeroVal string
	v := reflect.ValueOf(&easyFS).Elem().Interface()
	return v != zeroVal
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
