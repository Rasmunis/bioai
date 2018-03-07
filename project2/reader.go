package main

import (
	"fmt"
	"image/jpeg"
	"os"
)

func main() {
	file, _ := os.Open("./86016/testImg.jpg")
	img, _ := jpeg.Decode(file)

	fmt.Println("Hello world", img.At(100, 100), img.Bounds())
}
