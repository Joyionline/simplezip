package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	err error
	fw  *os.File
	ir  io.ReadCloser
)

func main() {
	var src = "data.zip"
	var destfile = "data"
	if err := Unzip(destfile, src); err != nil {
		log.Fatal(err)
	}
}

func Unzip(dest, src string) (err error) {
	zr, err := zip.OpenReader(src)
	defer zr.Close()
	if err != nil {
		log.Fatal("Error:", err)
		return err
	}

	if dest != "" {
		if err := os.MkdirAll(dest, 0755); err != nil {
			log.Fatal("Error:", err)
			return err
		}
	}

	for _, file := range zr.File {
		path := filepath.Join(dest, file.Name)
		// 如果是目录，就创建目录
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(path, file.Mode()); err != nil {
				log.Fatal("Error:", err)
				return err
			}
			continue
		}
		ir, err = file.Open()
		if err != nil {
			ir.Close()
			log.Fatal("Error:", err)
			return err
		}
		fw, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
		if err != nil {
			fClose()
			log.Fatal("Error:", err)
			return err
		}

		n, err := io.Copy(fw, ir)
		if err != nil {
			fClose()
			log.Fatal("Error:", err)
			return err
		}
		fmt.Printf("成功解压 %s ,共写入了 %d 个字符的数据\n", path, n)

		// fw.Close()
		// ir.Close()
	}
	return nil
}

func fClose() {
	fw.Close()
	ir.Close()
}
