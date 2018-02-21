package main

import (
	"fmt"
	"image"
	"os"
	"time"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"pixel/imdraw"
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

	var (
		frames= 0
		second= time.Tick(time.Second)
	)

	music, _ := os.Open("song.mp3")
	s, format, _ := mp3.Decode(music)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(s)

	imd := imdraw.New(nil)

	const millisPerBeat = 60000 / 126

	pulse := 0.0

	ticker := time.NewTicker(time.Millisecond * millisPerBeat)
	go func() {
		for range ticker.C {
			pulse = 1.0
		}
	}()

	for !win.Closed() {

		pulse *= 0.999

		imd.Clear()
		imd.SetMatrix(pixel.IM.Scaled(win.Bounds().Center(), pulse))
		imd.Color = pixel.RGB(1, 0, 0)
		imd.Push(pixel.V(200, 100))
		imd.Color = pixel.RGB(0, 1, 0)
		imd.Push(pixel.V(800, 100))
		imd.Color = pixel.RGB(0, 0, 1)
		imd.Push(pixel.V(500, 700))
		imd.Polygon(0)

		win.Clear(pixel.RGBA{0, 0, 0, 0})

		imd.Draw(win)

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
