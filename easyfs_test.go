package easyfs

import (
	"testing"
)

func TestAddFile(t *testing.T) {
	// Add a file to the filesystem
	FS := NewFS()
	err := FS.WriteFile("hello.txt", []byte("Hello, world!"), 0777)
	if err != nil {
		t.Errorf("Error adding file to filesystem: %v", err)
	}

	// Check that the file was added successfully
	if _, err := FS.Open("hello.txt"); err != nil {
		t.Errorf("Error opening file: %v", err)
	}
}

func TestAddDir(t *testing.T) {
	// Add a directory to the filesystem
	FS := NewFS()
	err := FS.Mkdir("mydir")
	if err != nil {
		t.Errorf("Error adding directory to filesystem: %v", err)
	}

	// Check that the directory was added successfully
	if _, err := FS.Open("mydir"); err != nil {
		t.Errorf("Error opening directory: %v", err)
	}
}
