package main

import (
	"archive/zip"
	"fmt"
	"os"
)

func main() {

	rdr, err := os.Open("bigfile.zip")
	if err != nil {
		panic("failed to open file: " + err.Error())
	}

	finfo, err := rdr.Stat()
	if err != nil {
		panic("failed to read file info: " + err.Error())
	}

	zrdr, err := zip.NewReader(rdr, finfo.Size())
	if err != nil {
		panic("failed to open a zip reader: " + err.Error())
	}

	for _, f := range zrdr.File {
		fmt.Println(f.FileInfo())
	}
}
