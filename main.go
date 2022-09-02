package main

import (
	"math/rand"
	"time"

	"fyne-app/animation"
	"fyne-app/assets"
	"fyne-app/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

const (
	appID string = "com.alexkurata.app"
)

var initialCanvasSize = fyne.NewSize(600, 600)

func main() {
	rand.Seed(time.Now().Unix())

	// Make the application
	fyneApp := app.NewWithID(appID)
	fyneApp.SetIcon(assets.AssetDvdPng)

	// Load theme
	settings := fyneApp.Settings()
	settings.SetTheme(&theme.CustomTheme{})

	// Make the content window in the app
	window := fyneApp.NewWindow("DVD")
	window.CenterOnScreen()
	window.RequestFocus()

	// Set window initial size, and render
	window.Resize(initialCanvasSize)
	window.SetIcon(assets.AssetDvdPng)
	window.SetPadded(false)

	// Create content root in window
	root := window.Canvas()
	root.Content().Resize(initialCanvasSize)

	// Create moving DvD object
	dvdImg := animation.NewDvD(assets.AssetDvdBlueSvg)

	// Put object in container and append to the root
	content := container.NewWithoutLayout(dvdImg)
	root.SetContent(content)

	// Start bounce animation in background
	go animation.BounceDvD(content, dvdImg)

	// Start app
	window.ShowAndRun()
}
