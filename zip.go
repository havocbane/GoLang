package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"os"
)

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}

const pathPrefix = "C:\\Users\\jbane\\tmp\\"
const archiveFilename = "test.zip"
const unzipPathPrefix = "unzip\\"
const bufferSize = 80
var srcFilenames = [2]string{"test1.txt", "test2.txt"}

func main() {
	// Create a new zip containing the files in srcFilenames
	zipPath := pathPrefix + archiveFilename
	fmt.Println(fmt.Sprintf("Creating archive '%s'", zipPath))
	zipArchive, err := os.Create(zipPath)
	errorCheck(err)
	zipWriter := zip.NewWriter(zipArchive)

	for _, srcFilename := range srcFilenames {
		fmt.Println(fmt.Sprintf("Adding '%s' to archive", srcFilename))
		zipFile, err := zipWriter.Create(srcFilename)
		errorCheck(err)

		srcFilePath := pathPrefix + srcFilename
		srcFile, err := os.Open(srcFilePath)
		errorCheck(err)

		srcFileReader := bufio.NewReader(srcFile)
		buf := make([]byte, bufferSize)
		for {
			n, err := srcFileReader.Read(buf)
			if 0 == n || io.EOF == err {
				break
			}
			_, err = zipFile.Write(buf[:n])
			errorCheck(err)
		}
		srcFile.Close()
	}

	err = zipWriter.Close()
	errorCheck(err)
	zipArchive.Close()
	fmt.Println("Done creation.")

	// Extract files from newly created archive
	zipExtractor(zipPath, pathPrefix + unzipPathPrefix)
}

func zipExtractor(zipPath string, dstPath string) {
	fmt.Println(fmt.Sprintf("Extracting archive '%s' to '%s'", zipPath, dstPath))
	zipReader, err := zip.OpenReader(zipPath)
	errorCheck(err)
	defer zipReader.Close()

	err = os.MkdirAll(dstPath, 0666)
	errorCheck(err)

	for _, srcFile := range zipReader.File {
		fmt.Println(fmt.Sprintf("Extracting '%s' from archive", srcFile.Name))
		dstFilePath := dstPath + srcFile.Name
		dstFile, err := os.Create(dstFilePath)
		errorCheck(err)

		srcFileOpened, err := srcFile.Open()
		errorCheck(err)

		buf := make([]byte, bufferSize)
		srcReader := bufio.NewReader(srcFileOpened)
		for {
			n, err := srcReader.Read(buf)
			if 0 == n || io.EOF == err {
				break
			}
			dstFile.Write(buf[:n])
		}
		srcFileOpened.Close()
		dstFile.Close()
	}
	fmt.Println("Done extracting.")
}
