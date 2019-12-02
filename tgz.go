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

func UnTgz(file io.Reader, initialPath, outName string) {
	initialPathTgz := initialPath
	FIRST_CHAR := 0
	if initialPath[FIRST_CHAR] == "/"[FIRST_CHAR] {
		initialPathTgz = strings.Replace(initialPathTgz, "/", "", 1)
	}
	if initialPath[len(initialPath)-1] != "/"[FIRST_CHAR] {
		initialPathTgz += "/"
	}

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
			filename := header.Name

			oldFolder := folder
			folder = strings.Replace(folder, initialPathTgz, "", 1)
			if oldFolder == folder && initialPathTgz != "" {
				continue
			}

			filename = strings.Replace(filename, initialPathTgz, "", 1)

			err = os.MkdirAll(path.Join(outName, folder), os.ModePerm)
			HandleError(err)

			outFile, err := os.Create(path.Join(outName, filename))
			HandleError(err)
			defer outFile.Close()

			_, err = io.Copy(outFile, fileTar)
			HandleError(err)
		}
	}

	defer fileGzip.Close()
}
