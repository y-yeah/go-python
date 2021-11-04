package log

import (
	"fmt"

	"github.com/y-yeah/go-python/storage"
)

func PrintOne() {
	fmt.Println(storage.GetData())
	fmt.Println(storage.SetData(10))
}

func PrintTwo() {
	fmt.Println(storage.GetData())
}
