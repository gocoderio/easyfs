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

func TestAddZip(t *testing.T) {
	// Create a sample zip file
	zipContent := []byte{0x50, 0x4b, 0x03, 0x04, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x00}
	FS := NewFS()
	// Add the zip file to the filesystem
	err := FS.AddZip("myzip.zip", zipContent)
	if err != nil {
		t.Errorf("Error adding zip file to filesystem: %v", err)
	}

	// Check that the zip file was added successfully
	if _, err := FS.Open("myzip.zip"); err != nil {
		t.Errorf("Error opening zip file: %v", err)
	}
}
