package main

import (
	"fmt"
	"log"
	"time"

	img "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/ttf"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return fmt.Errorf("could not initialize SDL: %v", err)
	}
	defer sdl.Quit()

	ttf.Init()
	if err := ttf.Init(); err != nil {
		return fmt.Errorf("could not initialize TTF: %v", err)
	}
	defer ttf.Quit()

	w, r, err := sdl.CreateWindowAndRenderer(1200, 800, sdl.WINDOW_SHOWN)
	if err != nil {
		return fmt.Errorf("could not create window: %v", err)
	}
	defer w.Destroy()

	if err := drawTitle(r); err != nil {
		return fmt.Errorf("could not draw title: %v", err)
	}
	time.Sleep(3 * time.Second)

	if err := drawBackground(r); err != nil {
		return fmt.Errorf("could not draw background: %v", err)
	}
	time.Sleep(3 * time.Second)

	return nil
}

func drawBackground(r *sdl.Renderer) error {
	r.Clear()

	t, err := img.LoadTexture(r, "res/imgs/background.png")
	if err != nil {
		return fmt.Errorf("could not load background: %v", err)
	}
	defer t.Destroy()

	if err := r.Copy(t, nil, nil); err != nil {
		return fmt.Errorf("could not copy background: %v", err)
	}

	// _ = t

	r.Present()
	return nil
}

func drawTitle(r *sdl.Renderer) error {
	r.Clear()

	f, err := ttf.OpenFont("res/fonts/Pacifico.ttf", 15)
	if err != nil {
		return fmt.Errorf("could not load font: %v", err)
	}
	defer f.Close()

	c := sdl.Color{R: 255, G: 100, B: 0, A: 255}
	s, err := f.RenderUTF8Solid("Flappy Gopher", c)
	if err != nil {
		return fmt.Errorf("could not render title: %v", err)
	}
	defer s.Free()

	t, err := r.CreateTextureFromSurface(s)
	if err != nil {
		return fmt.Errorf("could not create texture: %v", err)
	}
	defer t.Destroy()

	if err := r.Copy(t, nil, nil); err != nil {
		return fmt.Errorf("could not copy texture: %v", err)
	}
	r.Present()

	return nil
}
