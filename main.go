package main

import (
	"image"
	"image/png" // register the PNG format with the image package
	"os"
	"image/color"
	"time"
	"fmt"
)


func makeGray(src image.Image) (*image.Gray) {
	// Create a new grayscale image
	bounds := src.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	rectangle := image.Rect(0, 0, w, h)
	gray := image.NewGray(rectangle)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			oldColor := src.At(x, y)
			grayColor := color.GrayModel.Convert(oldColor)
			gray.Set(x, y, grayColor)
		}
	}

	return gray
}

func main() {
	start := time.Now()

	infile, err := os.Open("test1.png")
	if err != nil {
		// replace this with real error handling
		panic(err)
	}
	defer infile.Close()

	// Decode will figure out what type of image is in the file on its own.
	// We just have to be sure all the image packages we want are imported.
	src, _, err := image.Decode(infile)
	if err != nil {
		// replace this with real error handling
		panic(err)
	}

	gray := makeGray(src)

	// create integral image
	integral := integralImage(gray)

	min, max := -1, 2

	kernel := [][]int {
		{min, min, min},
		{max, max, max},
		{min, min, min},
	}

	bounds := gray.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	rectangle := image.Rect(0, 0, w, h)
	final := image.NewGray(rectangle)

	kernelLength := len(kernel)

	for y := 0; y < h - kernelLength; y++ {
		for x := 0; x < w - kernelLength; x++ {
			val := convolutionAt(kernel, gray, x, y)
			final.Set(x, y, color.Gray{uint8(val)})
		}
	}

	// Encode the grayscale image to the output file
	outfile, err := os.Create("output.png")
	if err != nil {
		// replace this with real error handling
		panic(err)
	}
	defer outfile.Close()
	png.Encode(outfile, final)

	elapsed := time.Since(start)
	fmt.Printf("shit took about: %s", elapsed)
}

func integralImage(img *image.Gray) [][]int {
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	integral := make([][]int, w)
	for i := range a {
		a[i] = make([]int, h)
	}

	integral[0][0] = int(img.GrayAt(0, 0))

	// First pass
	for x := 1; x < w; x++ {
		integral[x][0] = integral[x-1][0] + int(img.GrayAt(x, 0))
	}

	// First pass
	for y := 1; y < h; y++ {
		integral[0][y] = integral[0][y-1] + int(img.GrayAt(y, 0))
	}

	for y := 1; y < h; y++ {
		for x := 1; x < w; x++ {
			diag := integral[x-1][y-1]
			top := integral[x][y-1]
			left := integral[x-1][y]
			me := int(img.GrayAt(x, y))

			integral[x][y] = top + left + me - diag
		}
	}

	return integral
}

func convolutionAt(kernel [][]int, img *image.Gray, centerX int, centerY int) int {
	sum := 0
	kernelTotal := 0
	kernelLength := len(kernel)

	for y := 0; y < kernelLength; y++ {
		for x := 0; x < kernelLength; x++ {
			targetPixel := img.GrayAt(centerX + x, centerY + y)
			sum += kernel[y][x] * int(targetPixel.Y)
			kernelTotal += kernel[y][x]
		}
	}

	if kernelTotal == 0 {
		kernelTotal = 1
	}
	result := sum / kernelTotal

	return result
}