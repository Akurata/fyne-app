package animation

import (
	"fmt"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type coordinates struct {
	x float32
	y float32
}

func pickRandomPoint(width float32, height float32) coordinates {
	boundsArea := 2 * (height + width)
	pickEdge := rand.Float32() * boundsArea

	offset := width
	if pickEdge < offset { // Top
		fmt.Println("RANDOM PICK: TOP")
		return coordinates{pickEdge, 0}
	}

	if pickEdge < offset+height { // Right
		fmt.Println("RANDOM PICK: RIGHT")
		return coordinates{width, pickEdge - offset}
	}
	offset += height

	if pickEdge < offset+width { // Bottom
		fmt.Println("RANDOM PICK: BOTTOM")
		return coordinates{pickEdge - offset, height}
	}
	offset += width

	if pickEdge < offset+height { // Left
		fmt.Println("RANDOM PICK: LEFT")
		return coordinates{0, pickEdge - offset}
	}

	panic("Invalid random pick coordinates")
}

func getDelta(bounds fyne.Size, box *canvas.Rectangle, target coordinates) coordinates {
	pos := box.Position()

	fmt.Println("Target: ", target)

	// Calculate the slope to get to the target
	slope := coordinates{
		x: target.x - pos.X,
		y: target.y - pos.Y,
	}

	return coordinates{
		x: slope.x / (bounds.Width - pos.X),
		y: slope.y / (bounds.Height - pos.Y),
	}
}

func BounceDvD(canvas fyne.Canvas, box *canvas.Rectangle) {
	fmt.Println("\n--- BounceDvD ---")

	bounds := canvas.Size()

	pos := box.Position()
	size := box.Size()

	fmt.Println("Position: ", pos)
	fmt.Println("Size: ", size)

	// START: Pick container edge
	targetCoords := pickRandomPoint(bounds.Width, bounds.Height)
	delta := getDelta(bounds, box, targetCoords)
	fmt.Println("Delta: ", delta)

	speed := float32(2)
	for {
		bounds = canvas.Size()
		pos = box.Position()

		move := fyne.NewPos(
			pos.X+(delta.x*speed),
			pos.Y+(delta.y*speed),
		)

		// Determine if move hits bounds
		if move.Y <= 0 || move.Y >= (bounds.Height-size.Height) {
			delta.y *= -1
			continue
		}

		if move.X <= 0 || move.X >= (bounds.Width-size.Width) {
			delta.x *= -1
			continue
		}

		box.Move(move)
		time.Sleep(time.Second / 60)
		box.Refresh()
	}

	fmt.Println("DONE")
	return
}
