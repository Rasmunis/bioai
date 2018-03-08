package reader

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"math/rand"
	"os"
)

func main() {
	file, _ := os.Open("./86016/Test image.jpg")
	img, _ := jpeg.Decode(file)
	color1 := img.At(100, 100)
	color2 := img.At(100, 200)
	fmt.Println("Color 1", color1)
	fmt.Println("Color 2", color2)
	fmt.Println("Diff", euclRGBdist(color1, color2))
}

func randomInit(img image.Image) []int {
	// get image dimensions
	maxX := img.Bounds().Max.X
	maxY := img.Bounds().Max.Y

	initSol := make([]int, maxX*maxY)
	for i := 0; i < maxX*maxY; i++ {
		initSol[i] = rand.Intn(5)
	}
	return initSol
}

func euclRGBdist(color1, color2 color.Color) float64 {
	r1, g1, b1, _ := color1.RGBA()
	r2, g2, b2, _ := color2.RGBA()
	return math.Sqrt(float64(((r1-r2)*(r1-r2) + (g1-g2)*(g1-g2) + (b1-b2)*(b1-b2)) >> 16))
}

// func prims(img image.Image) {
// 	// get image dimensions
// 	maxX := img.Bounds().Max.X
// 	maxY := img.Bounds().Max.Y
//
// 	// initialize mst
// 	mst := make([]byte, maxX*maxY)
//
// 	// keep track of which nodes are in the current mst
// 	nodesInMst := make([]int, 0)
//
// 	// generate random coordinate of initial mst-node
// 	randX := rand.Intn(maxX)
// 	randY := rand.Intn(maxY)
//
// 	initNode := randX * randY
//
// 	edgeHeap := &edgeHeap.EdgeHeap{}
// 	// while the mst doesn't contain all pixels
// 	for len(nodesInMst) < maxX*maxY {
//
// 	}
// }
//
// func getCoords(node, maxX, maxY int) (x, y int) {
// 	x := node % maxX
// 	y = node % maxY
// 	return
// }
//
//
//
// func getEdges(h *edgeHeap.EdgeHeap, node int, previousEdge edgeHeap.Edge, maxX, maxY int) []edgeHeap.Edge {
// 	x, y := getCoords(node, maxX, maxY)
// 	if ((node + 1) % maxX) != 0 {
// 		heap.push(h, edgeHeap.Edge{euclRGBdist(color1, color2)})
// 	}
//
// }
