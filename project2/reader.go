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
	"sort"
	"time"

	"./edgeHeap"
)

func main() {
	file, _ := os.Open("./86016/Test image.jpg")
	img, _ := jpeg.Decode(file)
	// color1 := img.At(100, 100)
	// color2 := img.At(100, 200)
	tree, edgi := Prims(img)
	pop := Cutter(tree, edgi, 2, 10, 20)

	fmt.Println(pop[0])
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

func Prims(img image.Image) ([]int, []edgeHeap.Edge) {
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

	// list of edges in mst, used for splitting into segments later
	edgesInMst := make([]edgeHeap.Edge, 0)

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
			// add edge to edgesInMst
			edgesInMst = append(edgesInMst, bestEdge.(edgeHeap.Edge))

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
	sort.SliceStable(edgesInMst, func(i, j int) bool {
		return edgesInMst[i].W > edgesInMst[j].W
	})
	return mst, edgesInMst
}

func Cutter(mst []int, edgesInMst []edgeHeap.Edge, popSize, cuts, nWorstEdges int) [][]int {
	if cuts > nWorstEdges {
		panic("YOOO, YOU CAN'T REMOVE MORE EDGES THAN THE N WORST EDGES.. (cuts > nWorstEdges)")
	}

	rand.Seed(time.Now().Unix())
	population := make([][]int, 0)

	for i := 0; i < popSize; i++ {
		individual := mst
		worstEdges := edgesInMst[:nWorstEdges]
		for j := 0; j < cuts; j++ {
			randInt := rand.Intn(nWorstEdges - 1 - j)
			fmt.Println("INDEX", randInt, "WORST", worstEdges, "SIZE", len(worstEdges))
			currentEdge := worstEdges[randInt]
			worstEdges = append(worstEdges[:randInt], worstEdges[randInt+1:]...)
			individual[currentEdge.Src] = 4
		}
		population = append(population, individual)
	}
	return population
}

func getCoords(node, maxX, maxY int) (x, y int) {
	x = node % maxX
	y = int(node / maxY)
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
