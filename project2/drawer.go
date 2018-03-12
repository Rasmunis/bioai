package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

func getCoords2(node, maxX, maxY int) (x, y int) {
	x = node % maxX
	y = int(node / maxX)
	return
}

func DrawBnW(segments []map[int]empty, maxX, maxY int, original image.Image) (image.Image, image.Image) {
	img := image.NewRGBA(image.Rect(0, 0, maxX, maxY))
	b := original.Bounds()
	img2 := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(img2, img2.Bounds(), original, b.Min, draw.Src)

	for _, seg := range segments {
		for pix, _ := range seg {
			_, right := seg[pix+1]
			_, left := seg[pix-1]
			_, up := seg[pix-maxX]
			_, down := seg[pix+maxX]
			x, y := getCoords2(pix, maxX, maxY)
			switch {
			case !right && (pix+1)%maxX != 0:
				img.Set(x, y, color.RGBA{0, 0, 0, 255})
				img2.Set(x, y, color.RGBA{50, 205, 50, 255})
			case !left && pix%maxX != 0:
				img.Set(x, y, color.RGBA{0, 0, 0, 255})
				img2.Set(x, y, color.RGBA{50, 205, 50, 255})
			case !up && (pix-maxX) > 0:
				img.Set(x, y, color.RGBA{0, 0, 0, 255})
				img2.Set(x, y, color.RGBA{50, 205, 50, 255})
			case !down && (pix+maxX) < maxX*maxY:
				img.Set(x, y, color.RGBA{0, 0, 0, 255})
				img2.Set(x, y, color.RGBA{50, 205, 50, 255})
			default:
				img.Set(x, y, color.RGBA{255, 255, 255, 255})
			}
		}
	}
	f1, _ := os.OpenFile("BnWsegments.png", os.O_WRONLY|os.O_CREATE, 0600)
	f2, _ := os.OpenFile("GreenSegments.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f1.Close()
	defer f2.Close()
	png.Encode(f1, img)
	jpeg.Encode(f2, img2, nil)
	return img, img2
}
