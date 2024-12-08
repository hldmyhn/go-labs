package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"sync"
	"time"
)

var blurMatrix = [3][3]float64{
	{0.0625, 0.125, 0.0625},
	{0.125, 0.25, 0.125},
	{0.0625, 0.125, 0.0625},
}

func applyGaussian(img image.Image, x, y int) color.RGBA {
	var rSum, gSum, bSum, sum float64
	bounds := img.Bounds()

	for ky := -1; ky <= 1; ky++ {
		ny := y + ky
		if ny < bounds.Min.Y || ny >= bounds.Max.Y {
			continue
		}
		for kx := -1; kx <= 1; kx++ {
			nx := x + kx
			if nx < bounds.Min.X || nx >= bounds.Max.X {
				continue
			}
			pixel := img.At(nx, ny).(color.RGBA)
			kernelWeight := blurMatrix[ky+1][kx+1]

			rSum += float64(pixel.R) * kernelWeight
			gSum += float64(pixel.G) * kernelWeight
			bSum += float64(pixel.B) * kernelWeight
			sum += kernelWeight
		}
	}

	return color.RGBA{
		R: uint8(rSum / sum),
		G: uint8(gSum / sum),
		B: uint8(bSum / sum),
		A: 255,
	}
}

func apply(src image.Image, dst draw.Image, y int, wg *sync.WaitGroup) {
	defer wg.Done()
	bounds := src.Bounds()
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		newColor := applyGaussian(src, x, y)
		dst.Set(x, y, newColor)
	}
}

func filter_(img image.Image) draw.Image {
	bounds := img.Bounds()
	dst := image.NewRGBA(bounds)
	var wg sync.WaitGroup

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		wg.Add(1)
		go apply(img, dst, y, &wg)
	}

	wg.Wait()
	return dst
}

func main() {
	dog := "dog_blur.png"

	file, err := os.Open(dog)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	start := time.Now()
	processedImg := filter_(img)
	fmt.Println("4 задание:", time.Since(start))

	outFile, err := os.Create(dog)
	if err != nil {
		panic(err)
	}
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
		}
	}(outFile)

	err = png.Encode(outFile, processedImg)
	if err != nil {
		panic(err)
	}
}
