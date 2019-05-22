package main

import (
	"flag"
	"fmt"
	"log"
	"simplezip/unzip"
	"simplezip/zip"
	"strings"
)

func main() {
	operate := flag.String("o", "", "please input your order,such as -o=zip/unzip")
	srcname := flag.String("srcname", "", "file name")
	target := flag.String("target", "", "target file  name")
	flag.Parse()
	if *srcname == "" || *operate == "" {
		fmt.Println("请输入正确的参数")
		return
	}
	switch *operate {
	case "zip":
		if err := zip.Zip(*target, *srcname); err != nil {
			// log.Fatal("Error:", err)
			errs := fmt.Sprintf("%s", err)
			switch {
			case strings.Contains(errs, "no such file or director"):
				fmt.Println("源文件不存在，请检查后重试")
			default:
				fmt.Println("系统错误，请重试")
			}
		}
	case "unzip":
		if err := unzip.Unzip(*target, *srcname); err != nil {
			log.Fatal("Error:", err)
			fmt.Println("系统错误，请重试")
		}
	default:
		fmt.Println("执行的命令错误，目前仅支持：zip/unzip")
	}
}
