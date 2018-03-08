package main

import (
    "sort"
    "math"
    "math/rand"
    "time"
    "image"
)



type Solution struct{genom []rune; fitCon, fitDif float64; dist float64}

type sillyFiller struct{}


func mutate(sol Solution) Solution{
    r:=rand.New(rand.NewSource(time.Now().UnixNano()))
    N:= len(sol.genom)
    index := r.Int() % N
    mut:=1+(r.Int()%4)
    sol.genom[index]=(sol.genom[index]+mut)%5
    return sol
}

func nonDominatedRank(P []*Solution) [][]*Solution {

    nsqrt := int(math.Sqrt(float64(len(P))))
    rank := make(map[*Solution]int)
    F := make([][] *Solution, 1, nsqrt)
    F[0]=make([] *Solution, 1, nsqrt)
    domCount := make( map[*Solution]int)
    S := make(map[*Solution]map[*Solution]sillyFiller)
    for _, indp := range P {
        S[indp]=make(map[*Solution]sillyFiller)
        for _, indq := range P {
            if indq.fitCon < indp.fitCon && indq.fitDif<indp.fitDif{
                S[indp][indq] = struct{}{}
            }else if indp.fitCon<indq.fitCon&& indp.fitDif<indq.fitDif{
                domCount[indp]+=1
            }
        }
        if domCount[indp]==0 {
            F[0]=append(F[0],indp)
            rank[indp]=0
        }
    }
    i :=0
    for len(F[i])!=0{
        i+=1
        F=append(F,make([]*Solution,1,nsqrt))
        for _, indp := range (F[i-1]){
            for indq, _ := range (S[indp]){
                domCount[indq]-=1
                if domCount[indq]==0{
                    rank[indq]=i
                    F[i]=append(F[i],indq)
                }
            }
        }
    }
    return F
}





func crowdingDistAssign(P []*Solution) {
   
    
    l := len(P)
    maxcon := math.Inf(-1)
    mincon := math.Inf(1)
    maxdif := math.Inf(-1)
    mindif := math.Inf(1)
    for _, ind := range(P){
        ind.dist = 0
        maxcon = math.Max(maxcon, ind.fitCon)
        mincon = math.Min(mincon, ind.fitCon)
        maxdif=math.Max(maxdif, ind.fitDif)
        mindif=math.Min(mindif,ind.fitDif)
    }
    fcon := maxcon-mincon
    fdif := maxdif-mindif
    sort.Slice(P, func(i,j int)bool {
        return P[i].fitCon>P[j].fitCon
    })
    P[0].dist = math.Inf(1)
    P[l-1].dist=math.Inf(1)
    for i:=1; i<l-1; i++{
        P[i].dist += (P[i+1].fitCon-P[i-1].fitCon)/fcon
    }
    sort.Slice(P, func(i,j int)bool {
        return P[i].fitDif>P[j].fitDif
    })
    P[0].dist=math.Inf(1)
    P[l-1].dist=math.Inf(1)
    for i:=1; i< l-1; i++{
        P[i].dist += (P[i+1].fitDif-P[i-1].fitDif)/fdif
    }
}
    
func fitnessDif(sol Solution, im image.Image) float64{
    
    
    
}





















