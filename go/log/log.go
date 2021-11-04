package log

import (
	"fmt"

	"github.com/y-yeah/go-python/go/storage"
)

func PrintOne() {
	fmt.Println(storage.GetData())
	fmt.Println(storage.SetData(10))
}

func PrintTwo() {
	fmt.Println(storage.GetData())
}
