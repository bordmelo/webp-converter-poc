package main

import (
	"fmt"

	"github.com/nickalie/go-webpbin"
)

func main() {
	err := webpbin.NewCWebP().
		Quality(80).
		InputFile("image.jpeg").
		OutputFile("image.webp").
		Run()

	fmt.Println(err)
}
