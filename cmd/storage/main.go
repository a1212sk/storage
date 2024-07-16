package main

import (
	"fmt"
	"github.com/a1212sk/storage/internal/storage"
)

func main() {
	fmt.Println("hi")

	storage := storage.NewStorage()
	file, err := storage.Save("test", []byte("hello world"))
	if err != nil {
		panic(err)
	}
	file, err = storage.GetById(file.Id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Downloaded:", string(file.Data))

}
