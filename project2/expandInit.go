package main

import (
	"image"
	"image/color"
	"math"
	"sort"
)

func expandInit(mst, treeSize []int, avrgColor []color.Color, nodenr int, img *image.Image) {

	x, y := getxy(nodenr, img)
	up, ri, do, le := nextTo(nodenr, img)
	nbors := make([]int, 4)
	nbors = append(nbors[:0], up, ri, do, le)
	isPointedto := false
	treeSize[nodenr] = 1
	avrgColor[nodenr] = (*img).At(x, y)
	r, g, b, _ := (*img).At(x, y).RGBA()

	for _, j := range nbors {
		if pointsTo(mst[j], j, img) == nodenr {
			expandInit(mst, treeSize, avrgColor, j, img)
			treeSize[nodenr] += treeSize[j]
			segr, segg, segb, _ := avrgColor[j].RGBA()
			r += segr * uint32(treeSize[j])
			g += segg * uint32(treeSize[j])
			b += segb * uint32(treeSize[j])
			isPointedto = true
		}
	}
	if isPointedto {
		r /= uint32(treeSize[nodenr])
		g /= uint32(treeSize[nodenr])
		b /= uint32(treeSize[nodenr])
	}
	avrgColor[nodenr] = color.NRGBA{uint8(r / 0x101), uint8(g / 0x101), uint8(b / 0x101), 255}
}

type contrast struct {
	index    int
	contrast float64
}

func initPop(mst []int, cutnum, popSize int, img *image.Image) [][]int {
	N := len(mst)
	var root int
	pop := make([][]int, popSize)
	treeSize := make([]int, N)
	avrgColor := make([]color.Color, N)
	for i, dir := range mst {
		if dir == 4 {
			root = i
			break
		}
	}
	expandInit(mst, treeSize, avrgColor, root, img)
	contrastlist := make([]contrast, N)
	for i := 0; i < N; i++ {
		contrastlist[i] = getContrast(avrgColor[i], img, pointsTo(mst[i], i, img))
		contrastlist[i].index = i
	}
	sort.Slice(contrastlist, func(i, j int) bool {
		return contrastlist[i].contrast > contrastlist[j].contrast
	})
	i := 0
	var j int
	for len(pop) < popSize {
		j = 0
		for whatever := 0; whatever < cutnum; whatever++ {
			pop[i][contrastlist[j].index] = 4
			j++
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
