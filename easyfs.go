package easyfs

import (
	"io/fs"
	"reflect"
)

var EFS EasyFS //deprecated as public- use GetFS() instead

type EasyFS struct {
	MapFS
}

// returns a fresh new filesystem
func NewFS() EasyFS {
	return EasyFS{MapFS{}}
}

// returns an existing filesystem if there is one
func GetFS() EasyFS {
	if checkGlobalVar() {
		println("error in GetFS- variable not found")
		return EFS //if there is a global var, return the existing filesystem
	}
	return EasyFS{MapFS{}} // if not return a new one
}
func checkGlobalVar() bool {
	var zeroVal string
	v := reflect.ValueOf(&EFS).Elem().Interface()
	return v != zeroVal
}

// perm is unused, but you need to pass in something, like 0777
func (m EasyFS) Mkdir(name string, perm fs.FileMode) error {
	m.MapFS[name] = &MapFile{
		Mode: fs.ModeDir,
		//Mode: ModeDir,
	}
	return nil
}

// WriteFile writes data to a file named by filename. perm is not used but cn be set to
func (m EasyFS) WriteFile(name string, data []byte, perm fs.FileMode) error {
	//perm is not implimented
	if name[0] == '/' {
		name = name[1:] // FS filesystem in go cannot start with /
	}
	m.MapFS[name] = &MapFile{
		Data: data,
	}
	return nil

}
