package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var colors []rl.Color = []rl.Color{rl.Red, rl.Black, rl.Blue, rl.Brown, rl.DarkBlue, rl.Gold, rl.Lime, rl.Magenta}

var s_width int
var s_height int

var rectsCount int
var rects [rectsCount]rect

func main() {
	rl.InitWindow(1200, 800, "raylib [core] example - basic window")
	knowCount()
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	s_width = rl.GetScreenWidth()
	s_height = rl.GetScreenHeight()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		drawRects()
		changePercs()
		drawRects()
		rl.ClearBackground(rl.Beige)
		rl.EndDrawing()
	}
}

type rect struct {
	w     int32
	h     int32
	x     int32
	y     int32
	color rl.Color
}

func knowCount() {
	rectsCount = (s_width / 20) * (s_height / 20)
}

func randomRects(percs int) {
	for i := 0; i < s_width; i += 20 {
		for j := 0; j < s_height; j += 20 {
			if rand.Intn(100) > percs {
				for i <= rectsCount {
					rects[i] = rect{
						w:     20,
						h:     20,
						x:     int32(i),
						y:     int32(j),
						color: colors[rand.Intn(len(colors)-1)],
					}
				}
			}
		}
	}
}

func drawRects() {
	for _, v := range rects {
		rl.DrawRectangle(v.x, v.y, v.w, v.h, v.color)
	}
}

var percs int = 60

func changePercs() {
	if rl.IsKeyPressed(rl.KeyUp) {
		randomRects(percs + 1)
	} else if rl.IsKeyPressed(rl.KeyDown) {
		randomRects(percs - 1)
	}
}
