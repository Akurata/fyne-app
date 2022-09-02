package main

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"time"

	"fyne-app/animation"
	"fyne-app/assets"
	"fyne-app/theme"

	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

const (
	appID string = "com.alexkurata.app"
)

var initialCanvasSize = fyne.NewSize(600, 600)

type catFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func getCatFact() string {
	req, err := http.NewRequest("GET", "https://catfact.ninja/fact", nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var output catFact
	err = json.Unmarshal(respBody, &output)
	if err != nil {
		panic(err)
	}

	return output.Fact
}

func main() {
	rand.Seed(time.Now().Unix())

	// Make the canvas
	app := fyneApp.NewWithID(appID)
	app.SetIcon(assets.AssetDvdPng)

	window := app.NewWindow("DVD")
	window.CenterOnScreen()
	window.RequestFocus()

	// Set initial size, and render
	window.Resize(initialCanvasSize)
	window.SetIcon(assets.AssetDvdPng)
	window.SetPadded(false)

	root := window.Canvas()
	root.Content().Resize(initialCanvasSize)

	// Load theme
	settings := app.Settings()
	settings.SetTheme(&theme.CustomTheme{})

	// Make a simple rectangle
	// rect := canvas.NewRectangle(theme.Primary)
	// rect.Resize(fyne.NewSize(50, 50))
	// rect.Refresh()

	// Set initial size
	// dvdImg := canvas.NewImageFromResource(assets.AssetDvdBlueSvg)
	// imgSize := fyne.NewSize(imgWidth*imgScale, imgHeight*imgScale)
	// dvdImg.SetMinSize(imgSize)
	// dvdImg.Resize(imgSize)
	// dvdImg.Refresh()

	// animation.BounceDvD(window, rect)

	// window := fyneApp.NewWindow("Hello World")

	// // winSize := window.Content().Size()

	// textContentBind := binding.NewString()
	// textContentBind.Set("Get a cat fact")

	// text := widget.NewLabelWithData(textContentBind)
	// text.Wrapping = fyne.TextWrapWord

	// button := widget.NewButtonWithIcon("Click", fyneTheme.UploadIcon(), func() {
	// 	// fyneApp.SendNotification(&fyne.Notification{
	// 	// 	Title:   "Test Alert",
	// 	// 	Content: "Some text body to put in the alert",
	// 	// })
	// 	// text.SetText(getCatFact())

	// 	// animation.BounceDvD(window, rect)

	// 	animation.BounceDvD(root, rect)
	// })

	// buttonContainer := container.NewPadded(button)

	dvdImg := animation.NewDvD(assets.AssetDvdBlueSvg)
	content := container.NewWithoutLayout(dvdImg)
	root.SetContent(content)

	// button.Resize(fyne.NewSize(50, 50))
	// // buttonContainer := container.NewPadded(button)

	// container := container.New(
	// 	layout.NewVBoxLayout(),
	// 	layout.NewSpacer(),
	// 	text,
	// 	layout.NewSpacer(),
	// 	container.New(
	// 		layout.NewCenterLayout(),
	// 		container.New(
	// 			layout.NewCenterLayout(),
	// 			button,
	// 		),
	// 	),
	// 	layout.NewSpacer(),
	// )

	// fyneApp.SetIcon(assets.AssetIconPng)

	go animation.BounceDvD(content, dvdImg)

	window.ShowAndRun()

	// window.Run

	// window.SetContent(container)
	// window.ShowAndRun()
}
