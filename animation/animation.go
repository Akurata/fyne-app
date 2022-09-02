package animation

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"fyne-app/assets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

const (
	speed float32 = 2 // Speed scale of the image animation

	imgScale  float32 = 0.25 // The factor to scale the image
	imgWidth  float32 = 512  // Static DvD image width
	imgHeight float32 = 261  // Static DvD image height
)

// Index of possible DvD images to use
var dvdIndex = map[string]*fyne.StaticResource{
	assets.AssetDvdBlueSvg.StaticName:   assets.AssetDvdBlueSvg,
	assets.AssetDvdGreenSvg.StaticName:  assets.AssetDvdGreenSvg,
	assets.AssetDvdPurpleSvg.StaticName: assets.AssetDvdPurpleSvg,
	assets.AssetDvdRedSvg.StaticName:    assets.AssetDvdRedSvg,
	assets.AssetDvdYellowSvg.StaticName: assets.AssetDvdYellowSvg,
}

//
// Bounce DvD animation
//
// Classic DvD bouncing in the screen animation.
//
// Setup steps:
//  => Block until the starting viewport is setup.
//  => Move the DvD image into the center of the view.
//  => Pick a random vector relative to the center of the view
//
// Animation loop:
//  => Determine if the view dimensions have changed
//     and reposition the image back into view if the
//     the resize causes the image to go out of frame.
//  => Determine which side is smaller and scale the
//     relative vector speed.
//  => Calculate the next move.
//  => Determine if the move hits a boundary
//  => If the move does hit a side:
//     => Invert the (x if LR, y if TB) vector input
//     => Replace the displayed object with a different color
//     => Reset to top of the animation loop
//  => Otherwise, process the move
//
func BounceDvD(wrap *fyne.Container, box *canvas.Image) {
	fmt.Println("\n--- BounceDvD ---")

	// Block until container is actual render size
	bounds, size := blockUntilReady(wrap, box, time.Millisecond)

	// Move object to the relative center
	startPos := fyne.NewPos(
		(bounds.Width/2)-(size.Width/2),   // Container middle X
		(bounds.Height/2)-(size.Height/2), // Container middle Y
	)
	move(box, startPos)

	// Pick a random point in a radius around the starting position
	targetCoords := pickRandomPoint(startPos, float64(startPos.X*0.5))

	// Get a vector from the starting position to the target
	// within the relative container bounds
	vector := getVector(bounds, startPos, targetCoords)

	for {
		// Resize the box if the bounds state changed
		bounds = wrap.Size()
		pos := box.Position()

		maxX := bounds.Width - size.Width
		maxY := bounds.Height - size.Height

		// Catch if out of bounds and move into frame
		if pos.X > maxX {
			box.Move(fyne.NewPos(maxX, pos.Y))
			box.Refresh()
			continue
		}
		if pos.Y > maxY {
			box.Move(fyne.NewPos(pos.X, maxY))
			box.Refresh()
			continue
		}

		// Determine scale factor
		scaleX := float32(1)
		scaleY := float32(1)
		if bounds.Height > bounds.Width {
			scaleX = bounds.Height / bounds.Width
		} else {
			scaleY = bounds.Width / bounds.Height
		}

		// Calculate next coordinates
		newPos := fyne.NewPos(
			pos.X+(vector.X*(speed*scaleX)),
			pos.Y+(vector.Y*(speed*scaleY)),
		)

		// Determine if move hits bounds
		if newPos.Y <= 0 || newPos.Y >= maxY {
			vector.Y *= -1
			box = handleSideBounce(wrap, box)
			continue
		}

		if newPos.X <= 0 || newPos.X >= maxX {
			vector.X *= -1
			box = handleSideBounce(wrap, box)
			continue
		}

		// Process move
		move(box, newPos)
	}
}

// Create a DvD image object
func NewDvD(asset fyne.Resource) *canvas.Image {
	dvdImg := canvas.NewImageFromResource(asset)
	imgSize := fyne.NewSize(imgWidth*imgScale, imgHeight*imgScale)
	dvdImg.SetMinSize(imgSize)
	dvdImg.Resize(imgSize)
	dvdImg.Refresh()
	return dvdImg
}

// Convenience wrapper for moving the canvas object and triggering its refresh
func move(obj *canvas.Image, pos fyne.Position) {
	obj.Move(pos)
	time.Sleep(time.Second / 120) // Throttle
	obj.Refresh()
}

// Since the animation can start before the viewport is fully setup,
// this method will block until the viewport dimensions have changed
// which should be enough time to start calculating the initial vectors
func blockUntilReady(wrap *fyne.Container, box *canvas.Image, interval time.Duration) (fyne.Size, fyne.Size) {
	bounds := wrap.Size()
	size := box.Size()
	for bounds.Width == size.Width && bounds.Height == size.Height {
		bounds = wrap.Size()
		size = box.Size()
		time.Sleep(interval)
	}

	return bounds, size
}

// Get random point in circumference around center point
func pickRandomPoint(center fyne.Position, radius float64) fyne.Position {
	theta := (rand.Float64() * 360) * (math.Pi * 2)
	x := radius * math.Cos(theta)
	y := radius * math.Sin(theta)
	return fyne.Position{
		X: float32(x) + center.X,
		Y: float32(y) + center.Y,
	}
}

// Calculate the slope to get to the target, and scale the vector to viewport
func getVector(bounds fyne.Size, pos fyne.Position, target fyne.Position) fyne.Position {
	return fyne.Position{
		X: (target.X - pos.X) / (bounds.Width - pos.X),
		Y: (target.Y - pos.Y) / (bounds.Height - pos.Y),
	}
}

// Pick a different color to change the DvD icon to on edge bounce
func handleSideBounce(wrap *fyne.Container, box *canvas.Image) *canvas.Image {
	currName := box.Resource.Name()
	imgPool := []*fyne.StaticResource{}

	// Make a pool of random images to change to that are not the current one
	for name, img := range dvdIndex {
		if name != currName {
			imgPool = append(imgPool, img)
		}
	}

	// Pick a random index from pool
	next := imgPool[rand.Intn(len(imgPool))]

	// Create the image resource and move it into position
	dvdImg := NewDvD(next)
	dvdImg.Move(box.Position())

	// Delete the current dvd image
	wrap.RemoveAll()

	// Add the new one to the container
	wrap.Add(dvdImg)

	// Re-render both
	dvdImg.Refresh()
	box.Refresh()

	// Return new image to update pointer
	return dvdImg
}
