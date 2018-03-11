package main

import (
	"fmt"
	"image/jpeg"
	"math"
	"math/rand"
	"os"
	"sort"
	"time"
	/*
	   "image"
	   "fmt"
	   "github.com/twmb/algoimpl/go/graph"*/)

func main() {

	popSize := 2
	genNum := 2

	file, _ := os.Open("./86016/Test image.jpg")
	img, _ := jpeg.Decode(file)

	P := make([]*Solution, popSize, popSize)

	mst, edges := Prims(img)
	Genomes := Cutter(mst, edges, popSize, 50, 100)
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
		fmt.Println(len(NG) + len(P))
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
