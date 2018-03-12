package main

import (
	"fmt"
	"image/jpeg"
	"math"
	"math/rand"
	"os"
	"sort"
	"time"
)

/*
"sort"
"time"
   "image"
   "fmt"
   "github.com/twmb/algoimpl/go/graph"*/

func main() {

	popSize := 1
	genNum := 1

	file, _ := os.Open("./147091/Test image.jpg")
	img, _ := jpeg.Decode(file)
	//file, _ := os.Open("./out.png")
	//img, _ := png.Decode(file)

	P := make([]*Solution, popSize, popSize)

	mst, _ := Prims(img)

	//Genomes := Cutter(mst, edges, popSize, 100, len(mst), 500, &img)
	Genomes := initPop(mst, 100, popSize, &img)
	pop := make([]Solution, popSize, popSize)

	for i := 0; i < popSize; i++ {
		pop[i].Genome = Genomes[i]
		P[i] = &(pop[i])
	}
	for i := 0; i < genNum; i++ {
		fmt.Print(i)
		fmt.Println(" ")
		for _, sol := range P {
			sol.FitDif, sol.FitCon = fitness(*sol, &img)
		}
		nextGen := make([]Solution, 0, popSize)
		NG := make([]*Solution, 0, popSize)
		for j := 0; j < popSize; j++ {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			index := int(math.Pow(r.Float64(), 2)) * popSize
			child := mutate(pop[index])
			child.FitDif, child.FitCon = fitness(child, &img)
			nextGen = append(nextGen, child)
			NG = append(NG, &child)
		}
		F := nonDominatedRank(append(NG, P...))
		Q := make([]*Solution, 0, popSize)
		i := 0
		for len(F[i])+len(Q) < popSize {
			Q = append(Q, F[i]...)
			i++
		}

		sort.Slice(F[i], func(k, j int) bool {
			return (*(F[i][j])).Dist > (*(F[i][k])).Dist
		})
		n := popSize - len(Q)
		for j := 0; j < n; j++ {
			Q = append(Q, F[i][j])
		}
		P = Q
	}
	segments := findSegments(*(P[0]), &img)
	DrawBnW(segments, img.Bounds().Max.X, img.Bounds().Max.Y, img)
}
