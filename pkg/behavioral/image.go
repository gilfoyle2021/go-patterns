package strategy

import (
	"image"
	"image/color"
	"image/draw"
)

type ImageSquare struct {
	DestinationFilePath string
}

func (i *ImageSquare) Print() error {
	width := 80
	height := 80

	origin := image.Point{0, 0}

	bgImage := image.NewRGBA(image.Rectangle{
		Min: origin,
		Max: image.Point{X: width, Y: height},
	})
	bgColor := image.Uniform{
		color.RGBA{
			R: 70,
			G: 70,
			B: 70,
			A: 0,
		},
	}

	// quality := &jpeg.Options{Quality: 75}
	draw.Draw(bgImage, bgImage.Bounds(), &bgColor, origin, draw.Src)

	return nil
}
