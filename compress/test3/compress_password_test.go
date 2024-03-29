package compress

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	ezip "github.com/alexmullins/zip"
)

func TestCompressZipWithPassword(t *testing.T) {
	files := []string{"1.txt", "2.txt", "3.txt"}
	var raw bytes.Buffer
	zipw := ezip.NewWriter(&raw)
	defer zipw.Close()
	for _, file := range files {
		f, _ := os.Open(file)
		fInfo, _ := os.Stat(file)
		w, _ := zipw.Encrypt(fInfo.Name(), "123")
		io.Copy(w, f)
		f.Close()
	}
	zipw.Close()
	f, err := os.OpenFile("./a.zip", os.O_CREATE|os.O_WRONLY, 0666)
	defer f.Close()
	if err != nil {
		t.Error(err)
	}
	raw.WriteTo(f)
}

// CompressPathToZip 压缩文件夹
func CompressPathToZip(path, targetFile string) error {
	d, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()

	f, err := os.Open(path)
	if err != nil {
		return err
	}

	err = compress(f, "", w)

	return err
}

// EncryptZip 加密压缩文件
func EncryptZip(src, desc, password string) error {
	zipfile, err := os.Create(desc)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := ezip.NewWriter(zipfile)
	defer archive.Close()

	filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := ezip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = strings.TrimPrefix(path, filepath.Dir(src)+"/")
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}
		// 设置密码
		header.SetPassword(password)
		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})
	return err
}

func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
