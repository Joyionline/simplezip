package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	desc := "test.zip"
	src := "test.txt"
	zip(desc, src)
}

func zip(desc, src string) (err error) {
	file, err := os.Create(src)
	if err != nil {
		log.Fatal("Create the file error: ", err)
		return
	}
	if file != nil {
		defer file.Close()
	}
	w := tar.NewWriter(file)
	defer func() {
		if err := w.Close(); err != nil {
			log.Fatal("Close the Writer ", err)
		}
	}()

	return filepath.Walk(src, func(path string, fi os.FileInfo, errBack error) (err error) {
		if errBack != nil {
			return errBack
		}
		// 通过文件信心，创建zip的文件信息
		fh, err := tar.FileInfoHeader(fi, path)
		if err != nil {
			return err
		}

		fh.Name = strings.TrimPrefix(path, string(filepath.Separator))
		if fi.IsDir() {
			fh.Name += "/"
		}

		err = w.WriteHeader(fh)
		if err != nil {
			log.Fatal("get the header error", err)
		}

		fmt.Println("当前的模式数据是:", fh.Mode)

		fr, err := os.Open(path)
		defer fr.Close()
		if err != nil {
			log.Fatal(err)
			return err
		}

		n, err := io.Copy(w, fr)
		if err != nil {
			return err
		}
		fmt.Printf("成功压缩文件：%s,共写入了 %d 个字符的数据\n", path, n)
		return nil
	})

}
