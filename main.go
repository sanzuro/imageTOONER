 package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
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
			R1, G1, B1, _ := f.At(x, y).RGBA()
			h, j, m := toHSL(f.At(x, y))
			q, qt, qtt := toRGB(h, j, m)
			fmt.Printf(" { %v %v %v }      , { %v %v %v } ,  { %v %v %v } \n ", uint8(R1), uint8(G1), uint8(B1), h, j, m, q*255, qt*255, 255*qtt)
		}
	}
	k, _ := os.Create("n.png")
	_ = png.Encode(k, f)
	defer k.Close()

}

func toHSL(c color.Color) (float64, float64, float64) {
	var h, s, l float64
	R1, G1, B1, _ := c.RGBA()
	r := float64(uint8(R1)) / 255
	g := float64(uint8(G1)) / 255
	b := float64(uint8(B1)) / 255

	max := math.Max(math.Max(r, g), b)
	min := math.Min(math.Min(r, g), b)

	l = (max + min) / 2

	delta := max - min
	if delta == 0 {

		return 0, 0, l
	}

	if l < 0.5 {
		s = delta / (max + min)
	} else {
		s = delta / (2 - max - min)
	}

	r2 := (((max - r) / 6) + (delta / 2)) / delta
	g2 := (((max - g) / 6) + (delta / 2)) / delta
	b2 := (((max - b) / 6) + (delta / 2)) / delta
	switch {
	case r == max:
		h = b2 - g2
	case g == max:
		h = (1.0 / 3.0) + r2 - b2
	case b == max:
		h = (2.0 / 3.0) + g2 - r2
	}

	switch {
	case h < 0:
		h++
	case h > 1:
		h--
	}

	return h, s, l
}

func toRGB(h, s, l float64) (float64, float64, float64) {

	if s == 0 {
		return l, l, l
	}

	var v1, v2 float64
	if l < 0.5 {
		v2 = l * (1 + s)
	} else {
		v2 = (l + s) - (s * l)
	}

	v1 = 2*l - v2

	r := hueToRGB(v1, v2, h+(1.0/3.0))
	g := hueToRGB(v1, v2, h)
	b := hueToRGB(v1, v2, h-(1.0/3.0))

	return r, g, b
}

func hueToRGB(v1, v2, h float64) float64 {
	if h < 0 {
		h++
	}
	if h > 1 {
		h--
	}
	switch {
	case 6*h < 1:
		return (v1 + (v2-v1)*6*h)
	case 2*h < 1:
		return v2
	case 3*h < 2:
		return v1 + (v2-v1)*((2.0/3.0)-h)*6
	}
	return v1
}
