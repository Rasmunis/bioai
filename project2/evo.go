package main

import (
    "os"
    "math/rand"
    "time"
    "math"
/*
    "image"
    "fmt"
    "github.com/twmb/algoimpl/go/graph"*/

)
import "image/jpeg"


func main() {
    
    popSize := 20
    genNum := 20
    
    file, _ := os.Open("./86016/Test image.jpg")
    img, _ := jpeg.Decode(file)
    P := make([]*Solution, popSize, popSize)
    pop := make([]Solution, popSize, popSize)
    for i:= 0; i<popSize; i++ {
        pop[i].Genome =init(img)
        P[i]=&(pop[i])

    }
    for i:= 0; i<genNum; i++{
        for _, sol := range(P){
            sol.FitDif, sol.FitCon = fitness(*sol, &img)
        }
        nextGen:= make([]Solution, 0, popSize)
        NG := make([]*Solution, 0, popSize)
        for j:= 0; j< genNum; j++{
            r:=rand.New(rand.NewSource(time.Now().UnixNano()))
            index := int(math.Pow(r.Float64(),2))* popSize
            child:=mutate(pop[index])
            child.FitDif, child.FitCon = fitness(child, &img)
            nextGen = append(nextGen, child)
            NG= append(NG, &child)
        }
        F:=nonDominatedRank(append(NG, P...))
        Q := make([]*Solution,0,popSize)
        i:=0
        for len(F[i])+ len(Q) < popSize{
            Q = append(Q,F[i]...)
            i++
        }
        sort.Slice(F[i], func(k,j int)bool {
            return (*(F[i][j])).Dist>(*(F[i][k])).Dist})
        
            
        }
        
    }
}











