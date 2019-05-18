package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// 源文件
	var srcname = "data.rar"
	var destdfile = "data.zip"
	if err := Zip(destdfile, srcname); err != nil {
		log.Fatal(err)
	}
}

func Zip(dest, src string) (err error) {
	// 创建准备写入的文件
	fw, err := os.Create(dest)
	defer fw.Close()
	if err != nil {
		log.Fatal("创建文件错误", err)
		return err
	}

	// 通过 *File 创建zip.Write
	zw := zip.NewWriter(fw)
	defer func() {
		// 检测以下是否成功关闭
		if err := zw.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// 下面奖文件写入zw,
	return filepath.Walk(src, func(path string, fi os.FileInfo, errBack error) (err error) {
		if errBack != nil {
			return errBack
		}
		// 通过文件信息，创建zip的文件信息
		fh, err := zip.FileInfoHeader(fi)
		if err != nil {
			return err
		}
		// 替换文件信息中的文件名称
		fh.Name = strings.TrimPrefix(path, string(filepath.Separator))

		if fi.IsDir() {
			fh.Name += "/"
		}
		//写入文件信息
		w, err := zw.CreateHeader(fh)
		if err != nil {
			return err
		}
		if !fh.Mode().IsRegular() {
			return nil
		}

		// 打开要压缩的文件
		fr, err := os.Open(path)
		defer fr.Close()
		if err != nil {
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
