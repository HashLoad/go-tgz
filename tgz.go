package tgz

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func UnTgz(file io.Reader, outName string) {
	fileGzip, err := gzip.NewReader(file)
	HandleError(err)

	fileTar := tar.NewReader(fileGzip)
	for {
		header, err := fileTar.Next()
		if err == io.EOF {
			break
		}
		HandleError(err)

		if header.Typeflag != tar.TypeDir {
			paths := strings.Split(header.Name, "/")
			folder := strings.Replace(header.Name, paths[len(paths)-1], "", -1)

			err = os.MkdirAll(path.Join(outName, folder), os.ModePerm)
			HandleError(err)

			outFile, err := os.Create(path.Join(outName, header.Name))
			HandleError(err)
			defer outFile.Close()

			_, err = io.Copy(outFile, fileTar)
			HandleError(err)
		}
	}

	defer fileGzip.Close()
}
