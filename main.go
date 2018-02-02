package main

import (
	"fmt"
	"image"
	"os"
	"time"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	plasma, err := loadPicture("plasma.png")
	if err != nil {
		panic(err)
	}

	star, err := loadPicture("star.png")
	if err != nil {
		panic(err)
	}

	var (
		frames = 0
		second = time.Tick(time.Second)
	)

	for !win.Closed() {

		//win.Clear(pixel.RGBA{0,0,0,0})
		win.Clear(colornames.Maroon)
		win.SetMatrix(pixel.IM)


		s1 := pixel.NewSprite(star, star.Bounds())
		s2 := pixel.NewSprite(plasma, plasma.Bounds())

		for x := 0; x < 7; x++ {
			for y := 0; y < 5; y++ {

				switch x {
				case 0: win.SetComposeMethod(pixel.ComposeOver)
				case 1: win.SetComposeMethod(pixel.ComposeRover)
				case 2: win.SetComposeMethod(pixel.ComposeOut)
				case 3: win.SetComposeMethod(pixel.ComposeIn)
				case 4: win.SetComposeMethod(pixel.ComposeOut)
				case 5: win.SetComposeMethod(pixel.ComposeRin)
				case 6: win.SetComposeMethod(pixel.ComposeRout)
				}

				s1.Draw(win, pixel.IM.Moved(pixel.V(90.0 + float64(x)*140, 110.0 + float64(y)*140)))

				switch y {
				case 0: win.SetComposeMethod(pixel.ComposeXor)
				case 1: win.SetComposeMethod(pixel.ComposeRover)
				case 2: win.SetComposeMethod(pixel.ComposeOut)
				case 3: win.SetComposeMethod(pixel.ComposeIn)
				case 4: win.SetComposeMethod(pixel.ComposeOut)
				case 5: win.SetComposeMethod(pixel.ComposeRin)
				case 6: win.SetComposeMethod(pixel.ComposeRout)
				}

				s2.Draw(win, pixel.IM.Moved(pixel.V(90.0 + float64(x)*140, 110.0 + float64(y)*140)))

			}
		}

		win.Update()

		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
			frames = 0
		default:
		}
	}
}

func main() {
	pixelgl.Run(run)
}