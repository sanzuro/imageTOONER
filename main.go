package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

type kolor struct {
	r, g, b, a uint32
}

func main() {

	file, err := os.Open("photo.png")

	if err != nil {
		fmt.Print(" error occured on the first step ")
		os.Exit(0)
	}
	img, err := png.Decode(file)
	rect := img.Bounds()

	nimg := image.NewRGBA(rect)

	var kernal [3][3]kolor

	for y := 0; y < rect.Dy(); y++ {
		for x := 0; x < rect.Dx(); x++ {

			if x > 0 && y > 0 && y < rect.Dy()-1 && x < rect.Dx()-1 {
				for p := 0; p < 3; p++ {
					for q := 0; q < 3; q++ {
						kernal[p][q].r, kernal[p][q].g, kernal[p][q].b, kernal[p][q].a = img.At(x-p+1, y-q+1).RGBA()
					}
				}
				var r, g, b uint8
				for p := 0; p < 3; p++ {
					for q := 0; q < 3; q++ {
						r += uint8(kernal[p][q].r) / 9
						g += uint8(kernal[p][q].g) / 9
						b += uint8(kernal[p][q].b) / 9
					}
				}
				nimg.Set(x, y, color.NRGBA{
					R: r,
					G: g,
					B: b,
					A: 255,
				})
			} else {
				nimg.Set(x, y, color.NRGBA{
					R: 0,
					G: 0,
					B: 0,
					A: 255,
				})
			}

		}

	}

	nfile, err := os.Create("newPhoto.png")
	defer nfile.Close()
	err = png.Encode(nfile, nimg)

	// f , _ := os.Create("new.png")
	f := image.NewRGBA(rect)

	for y := 0; y < rect.Dy(); y++ {
		for x := 0; x < rect.Dx(); x++ {

			r1, g1, b1, _ := img.At(x, y).RGBA()
			r2, g2, b2, _ := nimg.At(x, y).RGBA()
			f.Set(x, y, color.NRGBA{
				R: uint8(-r1 + r2),
				G: uint8(-g1 + g2),
				B: uint8(-b1 + b2),
				A: 255,
			})
			h, j := hsv(f.At(x, y))
			fmt.Printf("%v %v \n ", h, j)
		}
	}
	k, _ := os.Create("n.png")
	_ = png.Encode(k, f)
	defer k.Close()

}

func hsv(k color.Color) (int32, uint8, float32) {
	r1, g1, b1, _ := k.RGBA()
	r, g, b := uint8(r1), uint8(g1), uint8(b1)
	var max, min, kolor uint8
	if r > b {
		if r > g {
			max = r
			kolor = 0
		} else {
			max = g
			kolor = 1
		}
	} else {
		if b > g {
			max = b
			kolor = 2
		} else {
			max = g
			kolor = 1
		}
	}

	if r < b {
		if r < g {
			min = r

		} else {
			min = g
		}
	} else {
		if b < g {
			min = b
		} else {
			min = g
		}
	}
	if max-min == 0 {
		return int32(0), uint8(0), unit32(0)
	}
	if kolor == 0 {
		return int32(((g-b)/(max-min) + 0) * 60), uint8((max + min) / 2), float32(())
	}
	if kolor == 1 {
		return int32(((r-b)/(max-min) + 2) * 60), uint8((max - min) / 2)
	}
	if kolor == 2 {
		return int32(((r-g)/(max-min) + 4) * 60), uint8((max - min) / 2)
	}

	return int32(0), uint8(0)

}
