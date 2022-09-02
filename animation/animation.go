package animation

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"fyne-app/assets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	// "fyne.io/fyne/v2/container"
)

const (
	imgScale  float32 = 0.25
	imgWidth  float32 = 512
	imgHeight float32 = 261
)

var dvdIndex = map[string]*fyne.StaticResource{
	assets.AssetDvdBlueSvg.StaticName:   assets.AssetDvdBlueSvg,
	assets.AssetDvdGreenSvg.StaticName:  assets.AssetDvdGreenSvg,
	assets.AssetDvdOrangeSvg.StaticName: assets.AssetDvdOrangeSvg,
	assets.AssetDvdPurpleSvg.StaticName: assets.AssetDvdPurpleSvg,
	assets.AssetDvdRedSvg.StaticName:    assets.AssetDvdRedSvg,
	assets.AssetDvdYellowSvg.StaticName: assets.AssetDvdYellowSvg,
}

type State struct {
	Bounds fyne.Size
}

// Get random point in circumference around center
func pickRandomPoint(center fyne.Position, radius float64) fyne.Position {
	theta := (rand.Float64() * 90) * (math.Pi * 2)
	x := radius * math.Cos(theta)
	y := radius * math.Sin(theta)
	return fyne.Position{
		X: float32(x) + center.X,
		Y: float32(y) + center.Y,
	}
}

// // Pick a random position across the whole canvas
// func pickRandomPos(boundsSize fyne.Size, boxSize fyne.Size) fyne.Position {
// 	minWidth := boundsSize.Width * 0.25
// 	maxWidth := (boundsSize.Width * 0.75) - boxSize.Width

// 	minHeight := boundsSize.Height * 0.25
// 	maxHeight := (boundsSize.Height * 0.75) - boxSize.Height

// 	randX := minWidth + rand.Float32()*(maxWidth-minWidth)
// 	randY := minHeight + rand.Float32()*(maxHeight-minHeight)

// 	return fyne.NewPos(randX, randY)
// }

func getVector(bounds fyne.Size, pos fyne.Position, target fyne.Position) fyne.Position {
	// Calculate the slope to get to the target
	return fyne.Position{
		X: (target.X - pos.X) / (bounds.Width - pos.X),
		Y: (target.Y - pos.Y) / (bounds.Height - pos.Y),
	}
}

func NewDvD(asset fyne.Resource) *canvas.Image {
	dvdImg := canvas.NewImageFromResource(asset)
	imgSize := fyne.NewSize(imgWidth*imgScale, imgHeight*imgScale)
	dvdImg.SetMinSize(imgSize)
	dvdImg.Resize(imgSize)
	dvdImg.Refresh()
	return dvdImg
}

func changeColor(wrap *fyne.Container, box *canvas.Image) *canvas.Image {
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

func BounceDvD(wrap *fyne.Container, box *canvas.Image) {
	fmt.Println("\n--- BounceDvD ---")
	bounds := wrap.Size()
	size := box.Size()

	// Block until container is actual render size
	for {
		bounds = wrap.Size()
		size = box.Size()
		if bounds.Width == size.Width && bounds.Height == size.Height {
			time.Sleep(time.Millisecond)
		} else {
			break
		}
	}

	// Move to random starting position
	// startPos := pickRandomPos(initSize, size)
	startPos := fyne.NewPos((bounds.Width/2)-(size.Width/2), (bounds.Height/2)-(size.Height/2))
	box.Move(startPos)
	box.Refresh()

	fmt.Println("Position: ", box.Position())
	fmt.Println("Size: ", box.Size())

	// START: Pick container edge
	targetCoords := pickRandomPoint(startPos, float64(startPos.X*0.5))
	fmt.Println("Target", targetCoords)
	vector := getVector(bounds, startPos, targetCoords)
	fmt.Println("Vector: ", vector)

	speed := float32(2)

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
		move := fyne.NewPos(
			pos.X+(vector.X*(speed*scaleX)),
			pos.Y+(vector.Y*(speed*scaleY)),
		)

		// Determine if move hits bounds
		if move.Y <= 0 || move.Y >= maxY {
			vector.Y *= -1
			box = changeColor(wrap, box)
			continue
		}

		if move.X <= 0 || move.X >= maxX {
			vector.X *= -1
			box = changeColor(wrap, box)
			continue
		}

		box.Move(move)
		time.Sleep(time.Second / 120)
		box.Refresh()

	}
}
