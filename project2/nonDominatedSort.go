package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"math/rand"
	"sort"
	"time"
)

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type Solution struct {
	Genome         []int
	FitCon, FitDif float64
	Dist           float64
}

type empty struct{}

func mutate(sol Solution) Solution {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	N := len(sol.Genome)
	index := r.Int() % N
	mut := 1 + (r.Int() % 4)
	sol.Genome[index] = (sol.Genome[index] + mut) % 5
	return sol
}

func nonDominatedRank(P []*Solution) [][]*Solution {

	nsqrt := int(math.Sqrt(float64(len(P))))
	rank := make(map[*Solution]int)
	F := make([][]*Solution, 1, nsqrt)
	F[0] = make([]*Solution, 0, nsqrt)
	domCount := make(map[*Solution]int)
	S := make(map[*Solution]map[*Solution]empty)
	for _, indp := range P {
		S[indp] = make(map[*Solution]empty)
		for _, indq := range P {
			if indq.FitCon < indp.FitCon && indq.FitDif < indp.FitDif {
				S[indp][indq] = struct{}{}
			} else if indp.FitCon < indq.FitCon && indp.FitDif < indq.FitDif {
				domCount[indp] += 1
			}
		}
		if domCount[indp] == 0 {
			F[0] = append(F[0], indp)
			rank[indp] = 0
		}
	}

	i := 0
	ind := 1
	for len(F[i]) != 0 {
		i += 1
		F = append(F, make([]*Solution, 0, nsqrt))
		for _, indp := range F[i-1] {
			for indq, _ := range S[indp] {
				if ind == 1 {
					ind = 2
				}

				domCount[indq] -= 1
				if domCount[indq] == 0 {
					rank[indq] = i
					F[i] = append(F[i], indq)
				}
			}
		}
	}
	fmt.Println("end of nonDominatedRank")
	return F
}

func crowdingDistAssign(P []*Solution) {

	l := len(P)
	maxcon := math.Inf(-1)
	mincon := math.Inf(1)
	maxdif := math.Inf(-1)
	mindif := math.Inf(1)
	for _, ind := range P {
		ind.Dist = 0
		maxcon = math.Max(maxcon, ind.FitCon)
		mincon = math.Min(mincon, ind.FitCon)
		maxdif = math.Max(maxdif, ind.FitDif)
		mindif = math.Min(mindif, ind.FitDif)
	}
	fcon := maxcon - mincon
	fdif := maxdif - mindif
	sort.Slice(P, func(i, j int) bool {
		return P[i].FitCon > P[j].FitCon
	})
	(*(P[0])).Dist = math.Inf(1)
	(*(P[l-1])).Dist = math.Inf(1)
	for i := 1; i < l-1; i++ {
		(*(P[i])).Dist += (P[i+1].FitCon - P[i-1].FitCon) / fcon
	}
	sort.Slice(P, func(i, j int) bool {
		return P[i].FitDif > P[j].FitDif
	})
	P[0].Dist = math.Inf(1)
	P[l-1].Dist = math.Inf(1)
	for i := 1; i < l-1; i++ {
		(*(P[i])).Dist += (P[i+1].FitDif - P[i-1].FitDif) / fdif
	}
}

func pointsToAlt(direction, x, y int, im *image.Image) (int, int) {
	bounds := (*im).Bounds()
	switch direction {
	case 5:
		return x, y
	case 0:
		if y <= bounds.Min.Y {
			return x, y
		} else {
			return x, y - 1
		}
	case 1:
		if x >= bounds.Max.X {
			return x, y
		} else {
			return x + 1, y
		}
	case 2:
		if y >= bounds.Max.Y {
			return x, y
		} else {
			return x, y + 1
		}
	case 3:
		if x <= bounds.Min.X {
			return x, y
		} else {
			return x - 1, y
		}
	case 4:
		return x, y
	}
	return x, y
}
func pointsTo(direction, index int, im *image.Image) int {
	bounds := (*im).Bounds()
	x, y := getxy(index, im)
	x, y = pointsToAlt(direction, x, y, im)
	return x + bounds.Max.X*y
}

func getxy(index int, im *image.Image) (int, int) {
	bounds := (*im).Bounds()
	x := index % bounds.Max.X
	y := index / bounds.Max.X
	return x, y
}

func getindex(x, y, xmax, ymin int) int {
	return (x + (y-ymin)*xmax)
}

func nextTo(index int, im *image.Image) (int, int, int, int) {

	bounds := (*im).Bounds()
	xmax := bounds.Max.X
	xmin := bounds.Min.X
	ymax := bounds.Max.Y
	ymin := bounds.Min.Y
	if index > (xmax-xmin)*(ymax-ymin) {
		fmt.Print("index is to large!")
	}
	x, y := getxy(index, im)
	ymax -= 1
	xmax -= 1
	l := getindex(Max(x-1, xmin), y, xmax, ymin)
	r := getindex(Min(x+1, xmax), y, xmax, ymin)
	d := getindex(x, Min(y+1, ymax), xmax, ymin)
	u := getindex(x, Max(y-1, ymin), xmax, ymin)

	return u, r, d, l
}

func fitness(sol Solution, im *image.Image) (float64, float64) {
	bounds := (*im).Bounds()
	size := (bounds.Max.X - bounds.Min.X) * (bounds.Max.Y - bounds.Min.Y)
	segments := make([]*map[int]empty, size, size)
	for i := 0; i < size; i++ {
		segments[i] = &map[int]empty{
			i: empty{}}
	}
	var target int
	for i, dir := range sol.Genome {
		target = pointsTo(dir, i, im)
		if target != i {
			for j, _ := range *segments[target] {
				(*segments[i])[j] = empty{}
			}
			segments[target] = segments[i]
		}
	}
	fmt.Println("Hei :)")
	segmentSet := make(map[*map[int]empty]empty)
	for _, segment := range segments {
		segmentSet[segment] = empty{}
	}
	var r, g, b uint32

	segmentColor := make(map[*map[int]empty]color.Color)
	for segment, _ := range segmentSet {
		r = 0
		g = 0
		b = 0
		for i, _ := range *segment {

			x, y := getxy(i, im)
			pr, pg, pb, _ := (*im).At(x, y).RGBA()

			r += pr
			g += pg
			b += pb
		}
		d := uint32(len(*segment))
		r /= d
		g /= d
		b /= d
		segmentColor[segment] = color.NRGBA{uint8(r / 0x101), uint8(g / 0x101), uint8(b / 0x101), 255}
	}
	nbors := make([]int, 0, 4)
	fitdif := 0.0
	fitcon := 0.0
	for segment, _ := range segmentSet {
		sr, sg, sb, _ := segmentColor[segment].RGBA()
		for i, _ := range *segment {
			pr, pg, pb, _ := (*im).At(getxy(i, im)).RGBA()
			rd := pr - sr
			gd := pg - sg
			bd := pb - sb
			fitdif += math.Sqrt(float64(rd*rd + gd*gd + bd*bd))
			u, d, l, f := nextTo(i, im)
			nbors = append(nbors[:0], u, d, l, f)
			for j := 0; j < 4; j++ {
				nbor := nbors[j]
				seg := segments[nbor]
				s2r, s2g, s2b, _ := segmentColor[seg].RGBA()
				rd = s2r - sr
				bd = s2b - sb
				gd = s2g - sg
				fitcon += math.Sqrt(float64(rd*rd + gd*gd + bd*bd))
			}
		}
	}
	return fitdif, fitcon
}
