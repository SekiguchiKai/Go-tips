package main

import "fmt"

type Hoge struct {
	// 構造体に関数型のフィールドを持たせることができる!
	Foo func(bar string) string
}

func main() {
	f := func(bar string) string {
		return fmt.Sprintf("%sだ!", bar)
	}

	h := &Hoge{Foo: f}

	fmt.Println(h.Foo("sekky0905"))
}
