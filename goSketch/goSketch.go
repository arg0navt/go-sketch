package goSketch

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// GetFiles: unzip sketch file
func GetFiles(src string, dir string) error {
	err := unzip(src, dir)
	if err != nil {
		return err
	}
	return nil
}

func unzip(src string, dir string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()
	for _, f := range r.File {
		fpath := filepath.Join(dir, f.Name)

		file, err := f.Open()
		if err != nil {
			panic(err)
		}
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}
		newFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			fmt.Println(1)
			return err
		}
		defer newFile.Close()
		_, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Println(2)
			return err
		}
		defer file.Close()
	}
	return nil
}
