package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Frank-Mayer/fyneflow"
	"image/color"
)

var (
	logo   *canvas.Image
	bottom *fyne.Container

	flow *fyneflow.Flow[string]
)

func Init() fyne.Window {
	rsc, err := fyne.LoadResourceFromURLString("https://raw.githubusercontent.com/Frank-Mayer/inno-lab/main/logo.png")
	if err != nil {
		panic(err)
	}
	logo = canvas.NewImageFromResource(rsc)
	bottom = container.New(layout.NewCenterLayout(), container.New(layout.NewGridWrapLayout(fyne.NewSize(75, 25)), logo), layout.NewSpacer())

	myApp := app.New()
	myWindow := myApp.NewWindow("UIUI")

	flow = fyneflow.NewFlow[string](myWindow)
	flow.Set("CreateUI", CreateUI)
	flow.Set("CreateScenario", CreateScenario)
	flow.Set("CameraLook", CameraLook)
	flow.Set("ShowPic", ShowPic)

	return myWindow
}

func CreateUI() fyne.CanvasObject {

	text1 := canvas.NewText("Bevor es losgeht...", color.White)
	text1.TextSize = 30
	text1.Alignment = fyne.TextAlignCenter
	text1.TextStyle = fyne.TextStyle{Monospace: true}

	text2 := canvas.NewText("Wie möchten Sie angesprochen werden?", color.White)
	text2.TextSize = 30
	text2.Alignment = fyne.TextAlignCenter
	text2.TextStyle = fyne.TextStyle{Monospace: true}

	button1 := widget.NewButton("Frau", func() {
		_ = flow.UseStateStr("gender", "Woman").Set("Woman")
		_ = flow.GoTo("CameraLook")
	})
	button2 := widget.NewButton("Herr", func() {
		_ = flow.UseStateStr("gender", "Man").Set("Man")
		_ = flow.GoTo("CameraLook")
	})
	//Non-Binary Button?

	top := container.New(layout.NewGridLayout(1), layout.NewSpacer(), container.New(layout.NewGridLayout(1), text1, text2, layout.NewSpacer()))
	middle := container.New(layout.NewGridLayout(1),
		container.New(layout.NewGridLayout(2),
			button1, button2), layout.NewSpacer())

	stack := container.NewStack(

		container.NewBorder(top,
			bottom, nil, nil,
			middle))

	return stack
}

func CreateScenario() fyne.CanvasObject {
	text1 := canvas.NewText("In welchem Szenario möchten Sie sich sehen?", color.White)
	text1.TextSize = 40
	text1.Alignment = fyne.TextAlignCenter
	text1.TextStyle = fyne.TextStyle{Monospace: true}

	button1 := widget.NewButton("Politiker:in", func() {
		_ = flow.GoTo("ShowPic")
	})
	button2 := widget.NewButton("Astronaut:in", func() {
		_ = flow.GoTo("ShowPic")
	})
	button3 := widget.NewButton("Im Urlaub", func() {
		_ = flow.GoTo("ShowPic")
	})
	button4 := widget.NewButton("Imagewechsel", func() {
		_ = flow.GoTo("ShowPic")
	})
	button5 := widget.NewButton("Pop-Star", func() {
		_ = flow.GoTo("ShowPic")
	})
	button6 := widget.NewButton("Fußballer:in", func() {
		_ = flow.GoTo("ShowPic")
	})
	button7 := widget.NewButton("Auf Abenteuer", func() {
		_ = flow.GoTo("ShowPic")
	})
	button8 := widget.NewButton("Model", func() {
		_ = flow.GoTo("ShowPic")
	})
	button9 := widget.NewButton("Im TED-Talk", func() {
		_ = flow.GoTo("ShowPic")
	})

	buttonBeenden := widget.NewButton("Beenden und Bilder löschen.", func() {
		_ = flow.GoTo("CreateUI")
		//flow.ClearStates ?
	})

	top := container.New(layout.NewGridLayout(1), layout.NewSpacer(), text1)

	middle := container.New(layout.NewGridLayout(1), layout.NewSpacer(),
		container.New(layout.NewGridLayout(3), button1, button2, button3, button4, button5, button6, button7, button8, button9), layout.NewSpacer())

	stack := container.NewStack(
		container.NewHBox(container.NewVBox(buttonBeenden)),
		container.NewBorder(top, bottom, nil, nil,
			middle))

	return stack
}

func CameraLook() fyne.CanvasObject {
	buttonBeenden := widget.NewButton("Beenden und Bilder löschen.", func() {
		flow.GoTo("CreateUI")
		//flow.ClearStates ?
	})

	text1 := canvas.NewText("Um Ihr Szenario zu erstellen wird ein Bild von Ihnen aufgenommen.", color.White)
	text1.TextSize = 20
	text1.Alignment = fyne.TextAlignCenter
	text1.TextStyle = fyne.TextStyle{Monospace: true}

	text2 := canvas.NewText("Schauen Sie gerade in die Kamera über dem Bildschirm.", color.White)
	text2.TextSize = 20
	text2.Alignment = fyne.TextAlignCenter
	text2.TextStyle = fyne.TextStyle{Monospace: true}

	text3 := canvas.NewText("Drücken Sie die Taste, wenn Sie bereit sind!", color.White)
	text3.TextSize = 15
	text3.Alignment = fyne.TextAlignCenter
	text3.TextStyle = fyne.TextStyle{Monospace: true}

	button1 := widget.NewButton("Bereit?", func() {
		flow.GoTo("CreateScenario")
	})

	stack := container.NewStack(
		container.NewHBox(container.NewVBox(buttonBeenden)),
		container.NewBorder(nil, bottom, nil, nil,
			container.NewCenter(
				container.NewVBox(
					container.New(layout.NewGridLayout(1), text1, text2, layout.NewSpacer(), text3, button1),
				))))

	return stack
}

func ShowPic() fyne.CanvasObject {
	buttonBeenden := widget.NewButton("Beenden und Bilder löschen.", func() {
		flow.GoTo("CreateUI")
		//flow.ClearStates ?

	})

	text1 := canvas.NewText("Bitte warten...", color.White)
	text1.TextSize = 50
	text1.Alignment = fyne.TextAlignCenter
	text1.TextStyle = fyne.TextStyle{Monospace: true}

	urlString, err := fyne.LoadResourceFromURLString("https://upload.wikimedia.org/wikipedia/commons/thumb/1/15/Cat_August_2010-4.jpg/2560px-Cat_August_2010-4.jpg")
	if err != nil {
		panic(err)
	}

	pic := canvas.NewImageFromResource(urlString)

	stack := container.NewStack(
		container.NewHBox(container.NewVBox(buttonBeenden)),
		container.NewBorder(nil, bottom, nil, nil,
			container.NewCenter(text1,
				container.NewVBox(container.NewGridWrap(fyne.NewSize(620, 620), pic),
					widget.NewButton("Weitere Szenarien auswählen.", func() {
						flow.GoTo("CreateScenario")
					})))))

	return stack
}
