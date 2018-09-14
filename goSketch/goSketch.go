package gosketch

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// GetFiles : unzip sketch file
func GetFiles(src string, dir string) error {
	r, err := unzip(src)
	if err != nil {
		return err
	}
	count := 1
	ch := make(chan int)
	for _, f := range r.File {
		go pushFileToDir(f, dir, ch)
	}
	for count < len(r.File) {
		select {
		case <-ch:
			{
				count++
			}
		}
	}

	return nil
}

func unzip(src string) (*zip.ReadCloser, error) {
	r, err := zip.OpenReader(src)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func pushFileToDir(f *zip.File, dir string, ch chan int) error {
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
	ch <- 1
	defer file.Close()
	return nil
}
