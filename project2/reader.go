package main

import (
	"container/heap"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"math/rand"
	"os"

	"./edgeHeap"
)

func main() {
	file, _ := os.Open("./86016/Test image.jpg")
	img, _ := jpeg.Decode(file)
	color1 := img.At(100, 100)
	color2 := img.At(100, 200)
	fmt.Println("Color 1", color1)
	fmt.Println("Color 2", color2)
	fmt.Println("Diff", euclRGBdist(color1, color2))
	prims(img)
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

func generateWeightedGraph(img image.Image) {
	type Node struct {
	}
}

func prims(img image.Image) []int {
	// get image dimensions
	maxX := img.Bounds().Max.X
	maxY := img.Bounds().Max.Y

	// initialize mst
	mst := make([]int, maxX*maxY)

	type empty struct{}
	// keep track of which nodes are in the current mst
	nodesInMst := make(map[int]empty)

	// generate random coordinate of initial mst-node
	randX := rand.Intn(maxX)
	randY := rand.Intn(maxY)

	// initial node
	initNode := randX * randY

	// initialize edgeHeap
	edges := make(edgeHeap.EdgeHeap, 0)

	// get edges from initial nodes and add to edgeHeap
	getEdges(&edges, initNode, -1, img)

	// add initial node to nodes in mst
	nodesInMst[initNode] = empty{}

	// while the mst doesn't contain all pixels
	for len(nodesInMst) < maxX*maxY {
		// get edge with smallest rgb difference
		bestEdge := heap.Pop(&edges)

		// get source and destination node of the edge
		src := bestEdge.(edgeHeap.Edge).Src
		dest := bestEdge.(edgeHeap.Edge).Dest

		// check if destination is already in the mst (cycle)
		_, cycle := nodesInMst[dest]
		if !cycle {
			// give the source a direction
			mst[src] = bestEdge.(edgeHeap.Edge).Direction
			// add destination to mst, without a direction yet (points to self)
			mst[dest] = 4
			// mark that the new node is in the mst
			nodesInMst[dest] = empty{}
			// add the edges from the new node to the heap
			getEdges(&edges, dest, src, img)
		}
	}
	return mst
}

func getCoords(node, maxX, maxY int) (x, y int) {
	x = node % maxX
	y = node % maxY
	return
}

func getEdges(h *edgeHeap.EdgeHeap, node int, previousNode int, img image.Image) {
	maxX := img.Bounds().Max.X
	maxY := img.Bounds().Max.Y
	x1, y1 := getCoords(node, maxX, maxY)
	color1 := img.At(x1, y1)

	if ((node+1)%maxX) != 0 && (node+1) != previousNode {
		x2, y2 := getCoords(node+1, maxX, maxY)
		heap.Push(h, edgeHeap.Edge{euclRGBdist(color1, img.At(x2, y2)), node, node + 1, 1})
	}

	if node%maxX != 0 && (node-1) != previousNode {
		x2, y2 := getCoords(node-1, maxX, maxY)
		heap.Push(h, edgeHeap.Edge{euclRGBdist(color1, img.At(x2, y2)), node, node - 1, 3})
	}

	if ((node + maxX) < maxX*maxY) && (node+maxX) != previousNode {
		x2, y2 := getCoords(node+maxX, maxX, maxY)
		heap.Push(h, edgeHeap.Edge{euclRGBdist(color1, img.At(x2, y2)), node, node + maxX, 2})
	}

	if ((node - maxX) > 0) && (node-maxX) != previousNode {
		x2, y2 := getCoords(node-maxX, maxX, maxY)
		heap.Push(h, edgeHeap.Edge{euclRGBdist(color1, img.At(x2, y2)), node, node - maxX, 0})
	}

}
