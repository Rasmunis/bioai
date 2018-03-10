package drawer

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func getCoords(node, maxX, maxY int) (x, y int) {
	x = node % maxX
	y = int(node / maxY)
	return
}

type Empty struct{}

func DrawBnW(segments []map[int]Empty, maxX, maxY int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, maxX, maxY))

	for _, seg := range segments {
		for pix := range seg {
			_, right := seg[pix+1]
			_, left := seg[pix-1]
			_, up := seg[pix-maxX]
			_, down := seg[pix+maxX]
			x, y := getCoords(pix, maxX, maxY)
			switch {
			case !right && (pix+1)%maxX != 0:
				img.Set(x, y, color.RGBA{0, 0, 0, 255})
			case !left && pix%maxX != 0:
				img.Set(x, y, color.RGBA{0, 0, 0, 255})
			case !up && (pix-maxX) > 0:
				img.Set(x, y, color.RGBA{0, 0, 0, 255})
			case !down && (pix+maxX) < maxX*maxY:
				img.Set(x, y, color.RGBA{0, 0, 0, 255})
			default:
				img.Set(x, y, color.RGBA{255, 255, 255, 255})
			}
		}
	}
	f, _ := os.OpenFile("BnWsegments.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
	return img
}
