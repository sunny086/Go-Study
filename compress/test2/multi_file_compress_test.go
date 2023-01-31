package test2

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

func TestMultiFileCompress(t *testing.T) {
	files := []string{"1.txt", "2.txt", "3.txt"}
	target := "compress.zip"

	err := compress(files, target)
	if err != nil {
		log.Fatalf("Compression failed: %v", err)
	}

	fmt.Printf("Compression successful: %s\n", target)
}

func compress(files []string, target string) error {
	// Create a zip archive
	zipFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// Create a new zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Add each file to the archive
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()

		// Get file information
		info, err := f.Stat()
		if err != nil {
			return err
		}

		// Add a file header to the archive
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// Write the file data to the archive
		_, err = io.Copy(writer, f)
		if err != nil {
			return err
		}
	}

	return nil
}
