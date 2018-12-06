package main

import (
	"archive/zip"
	"fmt"
	"os"
)

// https://golang.org/pkg/archive/zip/#example_Reader のコードを参考にさせていただいた。

func main() {
	// 書き出すzipを作成
	file, err := os.Create("sample.zip")
	if err != nil {
		fmt.Println(err)
	}

	// zipからWriterを作成
	w := zip.NewWriter(file)

	// zipに格納するファイルの元となる構造体を作成。
	var files = []struct {
		Name, Body string
	}{
		{"sample1.csv", "hoge1,foo1,bar1"},
		{"sample2.csv", "hoge2,foo2,bar2"},
		{"sample3.csv", "hoge3,foo3,bar3"},
	}
	for _, file := range files {
		// 与えられた名前でファイルを生成する。
		// ファイルの中身を書き込むためのWriterを返す。
		f, err := w.Create(file.Name)
		if err != nil {
			fmt.Println(err)
		}

		// ファイルのボディを書き込む。
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			fmt.Println(err)
		}
	}

	// Flushする。
	if err := w.Flush();err != nil {
		fmt.Println(err)
	}

	// 終わったらCloseする。
	defer w.Close()
}

