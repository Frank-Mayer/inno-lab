package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Frank-Mayer/fyneflow"
	"github.com/Frank-Mayer/inno-lab/internal/firebase"
	"image/color"
)

var (
	logo         *canvas.Image
	bottom       *fyne.Container
	flow         *fyneflow.Flow[string]
	gender       string
	promptString string
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
	//TODO withFrank: Prompts an extension übergeben
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
		promptString = "https://s.mj.run/hWqyh5IpNio photographic picture of a " + gender + " as a politican in a press conference --iw 2"
		_ = flow.GoTo("ShowPic")
		_ = flow.UseStateStr("prompt", promptString).Set(promptString)
	})
	button2 := widget.NewButton("Astronaut:in", func() {
		promptString = " ::4 https://s.mj.run/QY2Z4ddoxck ::1 ultrarealistic picture of a (" + gender + ") as an astronaut in a space suit, stars and planets in the background "
		_ = flow.GoTo("ShowPic")
		_ = flow.UseStateStr("prompt", promptString).Set(promptString)
	})
	button3 := widget.NewButton("Im Urlaub", func() {
		promptString = "::4 https://s.mj.run/wb3_lBMHQZw ::1 ultrarealistic picture of a " + gender + " at a holiday resort. warm tones --v 5.2 "
		_ = flow.GoTo("ShowPic")
		_ = flow.UseStateStr("prompt", promptString).Set(promptString)
	})
	button4 := widget.NewButton("Imagewechsel", func() {
		promptString = " ultrarealistic picture of a " + gender + " heavily tattood with piercings in their face, in the background you can see a tattoo parlor --iw 2 "
		_ = flow.GoTo("ShowPic")
		_ = flow.UseStateStr("prompt", promptString).Set(promptString)
	})
	button5 := widget.NewButton("Pop-Star", func() {
		promptString = " https://s.mj.run/DxMAyfGMFTY ultrarealistic picture of a " + gender + " as a popstar on the concert stage, microphone in their hand --iw 2 "
		_ = flow.GoTo("ShowPic")
		_ = flow.UseStateStr("prompt", promptString).Set(promptString)
	})
	button6 := widget.NewButton("Fußballer:in", func() {
		promptString = " ::4 https://s.mj.run/J0aQ9gVwN6k ::1 photographic of a " + gender + "as a football player in a press conference, logos in the background"
		_ = flow.GoTo("ShowPic")
		_ = flow.UseStateStr("prompt", promptString).Set(promptString)
	})
	button7 := widget.NewButton("Auf Abenteuer", func() {
		promptString = " ::5 https://s.mj.run/_zpWilkBFyQ ::1 ultrarealistic picture of a " + gender + " in a jungle sitting in the trees, beige clothes and a hat "
		_ = flow.GoTo("ShowPic")
		_ = flow.UseStateStr("prompt", promptString).Set(promptString)
	})
	button8 := widget.NewButton("Model", func() {
		promptString = " ::4 https://s.mj.run/BzHMDLF1RhE ::1 photographic picture of a " + gender + " as a model wearing haute couture on the catwalk "
		_ = flow.GoTo("ShowPic")
		_ = flow.UseStateStr("prompt", promptString).Set(promptString)
	})
	button9 := widget.NewButton("Im TED-Talk", func() {
		promptString = " ::4 https://s.mj.run/qcnxNbyX9hE ::1 Create a photograph of a " + gender + " holding a motivational speech at a TEDtalk. They are standing on a red, round carpet. There are people sitting in the crowd. 'TED' --v 6.0 "
		_ = flow.GoTo("ShowPic")
		_ = flow.UseStateStr("prompt", promptString).Set(promptString)
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

	rsc, err := fyne.LoadResourceFromURLString("https://raw.githubusercontent.com/Frank-Mayer/inno-lab/main/logo.png")
	if err != nil {
		panic(err)
	}

	beispielPic := canvas.NewImageFromResource(rsc)

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
		firebase.GetWebcamUrl()
		//TODO withFrank: Foto aufnehmen und auf firebase spreichern
	})

	stack := container.NewStack(
		container.NewHBox(container.NewVBox(buttonBeenden)),
		container.NewBorder(nil, bottom, nil, nil,
			container.NewCenter(
				container.NewVBox(container.NewCenter(container.NewGridWrap(fyne.NewSize(300, 300), beispielPic)),
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

	//TODO withFrank: Get and show Picture
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
