package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

// def ascii characters
var charMap = []string{
	"@", "#", "S", "%", "?", "*", "+", ";", ":", ",", ".",
}

func main() {

  // open and decode file
	if len(os.Args) < 2 {
		fmt.Println("invalid file location")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

  defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

  // resize, convert to ascii and print to stdout
	resizedImg := resize(img, 100, 0)

	asciiArt := convertToASCII(resizedImg)

	fmt.Println(asciiArt)
}

// resize original image to keep aspect ratio
func resize(img image.Image, width, height int) image.Image {
	bounds := img.Bounds()
	imgWidth := bounds.Max.X - bounds.Min.X
	imgHeight := bounds.Max.Y - bounds.Min.Y

	aspectRatio := float64(imgWidth) / float64(imgHeight)
	newWidth := width
	newHeight := int(float64(newWidth) / aspectRatio)

	if height > 0 {
		newHeight = height
		newWidth = int(float64(newHeight) * aspectRatio)
	}

	resized := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			sx := int(float64(x) / float64(newWidth) * float64(imgWidth))
			sy := int(float64(y) / float64(newHeight) * float64(imgHeight))
			resized.Set(x, y, img.At(sx, sy))
		}
	}

	return resized
}

// convert pixels to ascii chars
func convertToASCII(img image.Image) string {
	bounds := img.Bounds()
	asciiArt := ""

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := img.At(x, y)
			r, g, b, _ := pixel.RGBA()
			brightness := (r + g + b) / 3
			charIndex := int(brightness * uint32(len(charMap)) / 65535)
			asciiArt += charMap[charIndex]
		}
		asciiArt += "\n"
	}

	return asciiArt
}
