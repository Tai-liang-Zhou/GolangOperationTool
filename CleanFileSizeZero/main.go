package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func DeleteZeroSizeFile() {
	if err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		fileInfo, err := d.Info()
		if err != nil {
			fmt.Println(err)
		}
		fileSize := float64(fileInfo.Size()) / 1024 / 1024
		fmt.Printf("file path: %s\nfile size: %.5f MB\n", path, fileSize)

		if fileInfo.Mode().IsRegular() && fileSize == 0.0 {
			fmt.Printf("file siez 0 MB, remove file %s ", path)
			// os.Remove(path)
		}
		return nil

	}); err != nil {
		fmt.Println(err)
	}

}

func main() {

	DeleteZeroSizeFile()

}
