package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

func main() {
	file, _ := os.Open("./86016/Test image.jpg")
	img, _ := jpeg.Decode(file)
	r, g, b, _ := img.At(100, 100).NRGBA()
	fmt.Println("Hello world", r, g, b, img.Bounds())
}

type Node struct {
	X, Y       int
	neighbours []*Node
	root       *Node
}

func prims(img image.Image) {
	// get image dimensions
	// maxX := img.Bounds().Max.X
	// maxY := img.Bounds().Max.Y

	// // initialize mst
	// mst := make([]byte, maxX*maxY)
	//
	// // keep track of which nodes are in the current mst
	// nodesInMst := make([]int, 0)
	//
	// // generate random coordinate of initial mst-node
	// randX := rand.Intn(maxX)
	// randY := rand.Intn(maxY)
	//
	// // while the mst doesn't contain all pixels
	// for len(nodesInMst) < maxX*maxY {
	//
	// 	// check all edges from mst and pick the smallest one
	// 	for _, n := range nodesInMst {
	//
	// 	}
	// }
}

func euclRGBdist() {}
