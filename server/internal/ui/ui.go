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
	"time"
)

var (
	gender   string
	logo     = canvas.NewImageFromFile("C:\\Users\\yagmu\\GolandProjects\\inno-lab\\server\\internal\\ui\\240121_Veritas_logo_n-02-02.png")
	logoGrid = container.New(layout.NewGridWrapLayout(fyne.NewSize(75, 25)), logo)
	left     = container.New(layout.NewGridWrapLayout(fyne.NewSize(200, 100)), layout.NewSpacer())
	right    = container.New(layout.NewGridWrapLayout(fyne.NewSize(200, 100)), layout.NewSpacer())
	bottom   = container.New(layout.NewCenterLayout(), logoGrid, layout.NewSpacer())

	flow     *fyneflow.Flow[string]
	myWindow fyne.Window
)

func Init() fyne.Window {
	myApp := app.New()
	myWindow := myApp.NewWindow("UIUI")

	flow = fyneflow.NewFlow[string](myWindow)
	flow.Set("CreateUI", CreateUI)
	flow.Set("CreateScenario", CreateScenario)
	flow.Set("CameraLook", CameraLook)
	flow.Set("Countdown", Countdown)
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
		gender = "Woman"
		flow.UseStateStr("gender", "Woman").Set("Woman")
		flow.GoTo("CameraLook")
	})
	button2 := widget.NewButton("Herr", func() {
		gender = "Man"
		flow.UseStateStr("gender", "Man").Set("Man")
		flow.GoTo("CameraLook")
	})
	//Non-Binary Button?

	grid := container.New(layout.NewGridLayout(2), button1, button2)
	grid1 := container.New(layout.NewGridLayout(1), text1, text2, layout.NewSpacer())

	top := container.New(layout.NewGridLayout(1), layout.NewSpacer(), grid1)
	middle := container.New(layout.NewGridLayout(1), grid, layout.NewSpacer())
	content := container.NewBorder(top, bottom, left, right, middle)

	return content
}

func CreateScenario() fyne.CanvasObject {
	text1 := canvas.NewText("In welchem Szenario möchten Sie sich sehen?", color.White)
	text1.TextSize = 40
	text1.Alignment = fyne.TextAlignCenter
	text1.TextStyle = fyne.TextStyle{Monospace: true}

	button1 := widget.NewButton("Politiker:in", func() {
		flow.GoTo("ShowPic")
	})
	button2 := widget.NewButton("Astronaut:in", func() {
		flow.GoTo("ShowPic")
	})
	button3 := widget.NewButton("Im Urlaub", func() {
		flow.GoTo("ShowPic")
	})
	button4 := widget.NewButton("Imagewechsel", func() {
		flow.GoTo("ShowPic")
	})
	button5 := widget.NewButton("Pop-Star", func() {
		flow.GoTo("ShowPic")
	})
	button6 := widget.NewButton("Fußballer:in", func() {
		flow.GoTo("ShowPic")
	})
	button7 := widget.NewButton("Auf Abenteuer", func() {
		flow.GoTo("ShowPic")
	})
	button8 := widget.NewButton("Model", func() {
		flow.GoTo("ShowPic")
	})
	button9 := widget.NewButton("Im TED-Talk", func() {
		flow.GoTo("ShowPic")
	})

	buttonBeenden := widget.NewButton("Beenden und Bilder löschen.", func() {
		flow.GoTo("CreateUI")
		//flow.ClearStates ?
	})

	grid := container.New(layout.NewGridLayout(3), button1, button2, button3, button4, button5, button6, button7, button8, button9)

	toptop := container.New(layout.NewGridWrapLayout(fyne.NewSize(250, 50)), buttonBeenden)

	top := container.New(layout.NewGridLayout(1), toptop, layout.NewSpacer(), text1)
	middle := container.New(layout.NewGridLayout(1), layout.NewSpacer(), grid, layout.NewSpacer())

	content := container.NewBorder(top, bottom, left, right, middle)

	return content
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

	grid := container.New(layout.NewGridLayout(1), text1, text2, layout.NewSpacer(), text3, button1)
	gridC := container.New(layout.NewCenterLayout(), grid)

	toptop := container.New(layout.NewGridWrapLayout(fyne.NewSize(250, 50)), buttonBeenden)

	top := container.New(layout.NewGridLayout(1), toptop, layout.NewSpacer())
	middle := container.New(layout.NewGridLayout(1), gridC)
	content := container.NewBorder(top, bottom, left, right, middle)

	return content
}

func Countdown() fyne.CanvasObject {
	buttonBeenden := widget.NewButton("Beenden und Bilder löschen.", func() {
		flow.GoTo("CreateUI")
		//flow.ClearStates ?
	})

	circle := canvas.NewCircle(color.White)
	toptop := container.New(layout.NewGridWrapLayout(fyne.NewSize(250, 50)), buttonBeenden)

	top := container.New(layout.NewGridLayout(1), toptop, layout.NewSpacer())
	middle := container.New(layout.NewGridLayout(1), layout.NewSpacer(), circle, layout.NewSpacer())

	content := container.NewBorder(top, bottom, left, right, middle)
	return content
}

func Counto() {
	myApp := app.New()
	myWindow := myApp.NewWindow("ProgressBar Widget")

	progress := widget.NewProgressBar()
	infinite := widget.NewProgressBarInfinite()

	go func() {
		for i := 0.0; i <= 1.0; i += 0.1 {
			time.Sleep(time.Millisecond * 250)
			progress.SetValue(i)
		}
	}()

	myWindow.SetContent(container.NewVBox(progress, infinite))
	myWindow.ShowAndRun()
}

func ShowPic() fyne.CanvasObject {
	buttonBeenden := widget.NewButton("Beenden und Bilder löschen.", func() {
		flow.GoTo("CreateUI")
		//flow.ClearStates ?

	})

	urlString, err := fyne.LoadResourceFromURLString("https://upload.wikimedia.org/wikipedia/commons/thumb/1/15/Cat_August_2010-4.jpg/2560px-Cat_August_2010-4.jpg")
	if err != nil {
		panic(err)
	}

	pic := canvas.NewImageFromResource(urlString)
	stack := container.NewStack(
		container.NewHBox(container.NewVBox(buttonBeenden)),
		container.NewBorder(nil, bottom, nil, nil,
			container.NewCenter(
				container.NewVBox(container.NewGridWrap(fyne.NewSize(620, 620), pic),
					widget.NewButton("Weitere Szenarien auswählen.", func() {
						flow.GoTo("CreateScenario")
					})))))

	return stack
}
