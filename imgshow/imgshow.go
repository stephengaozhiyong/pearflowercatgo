package imgshow

import (
	"image"
	"image/color"
	"math"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
)

func ReadImage(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func ShowImgOnWindow(img image.Image) {
	myApp := app.New()
	myWindow := myApp.NewWindow("Image Viewer")
	imageCanvas := canvas.NewImageFromImage(img)
	imageCanvas.FillMode = canvas.ImageFillOriginal
	myWindow.SetContent(imageCanvas)

	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.SetIcon(theme.DocumentIcon())
	myWindow.ShowAndRun()
}

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)
	for i := 0; i < dx; i++ {
		item := make([]uint8, dy)
		for j := 0; j < dy; j++ {
			item[j] = uint8((float64(i) * math.Log(float64(j))))
		}
		res[i] = item
	}
	return res
}

func GetImage(f func(dx, dy int) [][]uint8) image.Image {
	const (
		dx = 256
		dy = 256
	)
	data := f(dx, dy)
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
		}
	}
	return m
}

type Image struct {
	w int
	h int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	// return image.Rect(0, 0, img.w, img.h)
	return image.Rect(0, 0, 256, 256)
}

func (img Image) At(x, y int) color.Color {
	v := uint8((float64(y) * math.Log(float64(x))))
	return color.RGBA{v, v, 255, 255}
}
