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
	//Genome contains the graphstructure
	Genome []int
	//Cuts contains the indexes such that Genome[index]==4
	Cuts []int
	//all the following variables are set by the fitness function
	FitCon, FitDif float64
	Dist           float64
	SegmentSlice   []map[int]empty
	//SegmentColor [i] is the color of segments[i]
	SegmentColor []color.Color
}

type empty struct{}

func mutate(sol Solution) Solution {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	segnr := r.Int() % len(sol.SegmentSlice)
	for i, _ := range sol.SegmentSlice[segnr] {
		if sol.Genome[i] == 4 {
			sol.Genome[i] = r.Int() % 4
		}
	}
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
	case 0:
		if y == bounds.Min.Y {
			return x, y
		} else {
			return x, y - 1
		}
	case 1:
		if x >= bounds.Max.X-1 {
			return x, y
		} else {
			return x + 1, y
		}
	case 2:
		if y >= bounds.Max.Y-1 {
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

	l := getindex(Max(x-1, xmin), y, xmax, ymin)
	r := getindex(Min(x+1, xmax-1), y, xmax, ymin)
	d := getindex(x, Min(y+1, ymax-1), xmax, ymin)
	u := getindex(x, Max(y-1, ymin), xmax, ymin)

	return u, r, d, l
}

func expand(index, count int, visited *map[int]empty, segmentSlice *[]map[int]empty, sol *Solution, im *image.Image) {
	(*visited)[index] = empty{}
	(*segmentSlice)[count][index] = empty{}
	u, l, d, r := nextTo(index, im)
	nbors := [4]int{u, l, d, r}
	target := pointsTo((*sol).Genome[index], index, im)
	_, ok := (*visited)[target]
	if !ok {
		expand(target, count, visited, segmentSlice, sol, im)
	}
	for _, j := range nbors {
		_, ok := (*visited)[j]
		if !ok && pointsTo(sol.Genome[j], j, im) == index {
			expand(j, count, visited, segmentSlice, sol, im)
		}
	}
}

func findSegments(sol Solution, im *image.Image) []map[int]empty {
	bounds := (*im).Bounds()
	size := (bounds.Max.X - bounds.Min.X) * (bounds.Max.Y - bounds.Min.Y)
	segments := make([]*[]int, size, size)
	for i := 0; i < size; i++ {
		segments[i] = &[]int{i}
	}
	count := 0
	segmentSlice := make([]map[int]empty, 0)
	visited := make(map[int]empty)
	for i, _ := range sol.Genome {
		_, ok := visited[i]
		if !ok {
			segmentSlice = append(segmentSlice, make(map[int]empty))
			expand(i, count, &visited, &segmentSlice, &sol, im)
			count++
		}
	}
	return segmentSlice
}

func fitness(sol *Solution, im *image.Image) {

	segmentSlice := findSegments(*sol, im)
	sol.SegmentSlice = segmentSlice
	fmt.Println(len(segmentSlice))
	var r, g, b uint32
	sol.SegmentColor = make([]color.Color, len(segmentSlice))
	for i, segment := range segmentSlice {
		r = 0
		g = 0
		b = 0
		for i, _ := range segment {

			x, y := getxy(i, im)
			pr, pg, pb, _ := (*im).At(x, y).RGBA()

			r += pr
			g += pg
			b += pb
		}
		d := uint32(len(segment))
		r /= d
		g /= d
		b /= d

		sol.SegmentColor[i] = color.NRGBA{uint8(r / 0x101), uint8(g / 0x101), uint8(b / 0x101), 255}
	}
	nbors := make([]int, 0, 4)
	fitdif := 0.0
	fitcon := 0.0
	for segnum, segment := range segmentSlice {
		sr, sg, sb, _ := sol.SegmentColor[segnum].RGBA()
		for i, _ := range segment {
			pr, pg, pb, _ := (*im).At(getxy(i, im)).RGBA()
			rd := pr - sr
			gd := pg - sg
			bd := pb - sb
			fitdif += math.Sqrt(float64(rd*rd + gd*gd + bd*bd))
			u, d, l, f := nextTo(i, im)
			nbors = append(nbors[:0], u, d, l, f)
			for j := 0; j < 4; j++ {
				nbor := nbors[j]
				var seg map[int]empty
				for _, seg = range segmentSlice {
				}
				_, ok := seg[nbor]
				if ok {
					break
				}
				s2r, s2g, s2b, _ := sol.SegmentColor[segnum].RGBA()
				rd = s2r - sr
				bd = s2b - sb
				gd = s2g - sg
				fitcon += math.Sqrt(float64(rd*rd + gd*gd + bd*bd))
			}
		}
	}
	sol.FitDif = fitdif
	sol.FitCon = fitcon
}

/*func expandInit(mst []int, treeSize *[]int, avrgColor *[]color.Color, nodenr int, img *image.Image) {

	x, y := getxy(nodenr, img)
	up, ri, do, le := nextTo(nodenr, img)
	nbors := make([]int, 4)
	nbors = append(nbors[:0], up, ri, do, le)
	isPointedto := false
	(*treeSize)[nodenr] = 1
	(*avrgColor)[nodenr] = (*img).At(x, y)
	r, g, b, _ := (*img).At(x, y).RGBA()

	for _, j := range nbors {
		if pointsTo(mst[j], j, img) == nodenr {
			expandInit(mst, treeSize, avrgColor, j, img)
			(*treeSize)[nodenr] += (*treeSize)[j]
			segr, segg, segb, _ := (*avrgColor)[j].RGBA()
			r += segr * uint32((*treeSize)[j])
			g += segg * uint32((*treeSize)[j])
			b += segb * uint32((*treeSize)[j])
			isPointedto = true
		}
	}
	if isPointedto {
		r /= uint32((*treeSize)[nodenr])
		g /= uint32((*treeSize)[nodenr])
		b /= uint32((*treeSize)[nodenr])
	}
	(*avrgColor)[nodenr] = color.NRGBA{uint8(r / 0x101), uint8(g / 0x101), uint8(b / 0x101), 255}
}

type contrast struct {
	index    int
	contrast float64
}

func initPop(mst []int, cutnum, popSize int, img *image.Image) [][]int {
	N := len(mst)
	var root int
	treeSize := make([]int, N)
	avrgColor := make([]color.Color, N)
	for i, dir := range mst {
		if dir == 4 {
			root = i
			break
		}
	}
	expandInit(mst, &treeSize, &avrgColor, root, img)
	contrastlist := make([]contrast, N)
	for i := 0; i < N; i++ {
		contrastlist[i] = getContrast(avrgColor[i], img, pointsTo(mst[i], i, img))
		contrastlist[i].index = i
	}
	sort.Slice(contrastlist, func(i, j int) bool {
		return contrastlist[i].contrast > contrastlist[j].contrast
	})
	pop := make([][]int, popSize)
	for i := 0; i < popSize; i++ {
		pop[i] = make([]int, N)
		copy(pop[i], mst)
		for whatever := 0; whatever < cutnum; whatever++ {
			pop[i][contrastlist[N-whatever-1].index] = 4
		}
	}
	return pop
}

func getContrast(col color.Color, img *image.Image, target int) contrast {

	var ret contrast
	ret.index = 0
	r, g, b, _ := col.RGBA()
	x, y := getxy(target, img)
	pr, pg, pb, _ := (*img).At(x, y).RGBA()
	dr := r - pr
	dg := g - pg
	db := b - pb
	ret.contrast = math.Sqrt(float64(dr*dr + dg*dg + db*db))
	return ret
}
*/
