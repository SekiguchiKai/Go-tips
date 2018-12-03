package main

import (
	"fmt"
	"bytes"
	"os"
)

func main() {
	// bytes.Bufferの使い方

	var buff bytes.Buffer

	// []byteにする
	b := []byte("I love Go!\n")
	write1(buff, b)
	s := "I love TypeScript!"
	write2(buff, s)

	// 標準出力に書き込む
	buff.WriteTo(os.Stdout)
}

func write1(buff bytes.Buffer, b []byte) {
	// Bufferに書き込み
	buff.Write(b)
}

func write2(buff bytes.Buffer, s string) {
	// Bufferに書き込み
	fmt.Fprintf(&buff, s)
}

// https://golang.org/pkg/bytes/#Buffer のtype BufferのExampleを参考にさせていただいている。