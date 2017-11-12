package Scaner

import (
	"path/filepath"
	"os"
	"fmt"
)

var formats  = map[string]bool{ ".txt": true,
								".doc": true,
								".jpg": true,
							}

var FileNames = make(chan string)

func scanner(path string, info os.FileInfo, err error) error {
	if formats[filepath.Ext(info.Name())] {
		fmt.Println(path)
		FileNames <- path
	}
	return nil
}

// ScanDirecory recursively scans directory "path"
func ScanDirectory(path string) {
	defer close(FileNames)
	filepath.Walk(path, scanner)
}