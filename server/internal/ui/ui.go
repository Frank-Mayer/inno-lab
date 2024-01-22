package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
	"time"
)

var (
	gender   string
	logo     = canvas.NewImageFromFile("C:\\Users\\yagmu\\GolandProjects\\inno-lab\\server\\internal\\ui\\240121_Veritas_logo_n-02-02.png")
	logoGrid = container.New(layout.NewGridWrapLayout(fyne.NewSize(75, 25)), logo)
	left     = container.New(layout.NewGridWrapLayout(fyne.NewSize(200, 100)), layout.NewSpacer())
	right    = container.New(layout.NewGridWrapLayout(fyne.NewSize(200, 100)), layout.NewSpacer())
	bottom   = container.New(layout.NewCenterLayout(), logoGrid, layout.NewSpacer())
)

func Init() fyne.Window {
	myApp := app.New()
	myWindow := myApp.NewWindow("UIUI")

	return myWindow
}
func CreateUI(window fyne.Window) {
	text1 := canvas.NewText("Bevor es losgeht...", color.White)
	text1.TextSize = 30
	text1.Alignment = fyne.TextAlignCenter

	text2 := canvas.NewText("Wie möchten Sie angesprochen werden?", color.White)
	text2.TextSize = 30
	text2.Alignment = fyne.TextAlignCenter

	button1 := widget.NewButton("Frau", func() {
		gender = "Woman"

	})
	button2 := widget.NewButton("Herr", func() {
		gender = "Man"

	})
	//Non-Binary Button?

	grid := container.New(layout.NewGridLayout(2), button1, button2)
	grid1 := container.New(layout.NewGridLayout(1), text1, text2, layout.NewSpacer())

	top := container.New(layout.NewGridLayout(1), layout.NewSpacer(), grid1)
	middle := container.New(layout.NewGridLayout(1), grid, layout.NewSpacer())
	content := container.NewBorder(top, bottom, left, right, middle)

	window.SetContent(content)
	window.ShowAndRun()
}

func CreateScenario() {
	myApp := app.New()
	myWindow := myApp.NewWindow("UIUI")

	text1 := canvas.NewText("In welchem Szenario möchten Sie sich sehen?", color.White)
	text1.TextSize = 40
	text1.Alignment = fyne.TextAlignCenter

	button1 := widget.NewButton("Politiker:in", func() {

	})
	button2 := widget.NewButton("Astronaut:in", func() {
		log.Println("uff")
	})
	button3 := widget.NewButton("Im Urlaub", func() {
		log.Println("tapped")
	})
	button4 := widget.NewButton("Imagewechsel", func() {
		log.Println("uff")
	})
	button5 := widget.NewButton("Pop-Star", func() {
		log.Println("tapped")
	})
	button6 := widget.NewButton("Fußballer:in", func() {
		log.Println("uff")
	})
	button7 := widget.NewButton("Auf Abenteuer", func() {
		log.Println("tapped")
	})
	button8 := widget.NewButton("Model", func() {
		log.Println("uff")
	})
	button9 := widget.NewButton("Im TED-Talk", func() {
		log.Println("uff")
	})

	buttonBeenden := widget.NewButton("Beenden und Bilder löschen.", func() {
		log.Println("Beenden Button gedrückt.")
	})

	grid := container.New(layout.NewGridLayout(3), button1, button2, button3, button4, button5, button6, button7, button8, button9)

	toptop := container.New(layout.NewGridWrapLayout(fyne.NewSize(250, 50)), buttonBeenden)

	top := container.New(layout.NewGridLayout(1), toptop, layout.NewSpacer(), text1)
	middle := container.New(layout.NewGridLayout(1), layout.NewSpacer(), grid, layout.NewSpacer())

	content := container.NewBorder(top, bottom, left, right, middle)
	myWindow.SetContent(content)
	//myWindow.SetFullScreen(true)
	myWindow.ShowAndRun()
}

func CameraLook() {
	myApp := app.New()
	myWindow := myApp.NewWindow("UIUI")

	text1 := canvas.NewText("Um Ihr Szenario zu erstellen wird ein Bild von Ihnen aufgenommen.", color.White)
	text1.TextSize = 30
	text1.Alignment = fyne.TextAlignCenter

	text2 := canvas.NewText("Schauen Sie gerade in die Kamera über dem Bildschirm.", color.White)
	text2.TextSize = 30
	text2.Alignment = fyne.TextAlignCenter

	text3 := canvas.NewText("Drücken Sie die Taste, wenn Sie bereit sind!", color.White)
	text3.TextSize = 20
	text3.Alignment = fyne.TextAlignCenter

	button1 := widget.NewButton("Bereit?", func() {
		log.Println("tapped")
	})

	grid := container.New(layout.NewGridLayout(1), text1, text2, layout.NewSpacer(), text3, button1)
	gridC := container.New(layout.NewCenterLayout(), grid)

	middle := container.New(layout.NewGridLayout(1), gridC)
	content := container.NewBorder(nil, bottom, left, right, middle)

	myWindow.SetContent(content)
	//myWindow.SetFullScreen(true)
	myWindow.ShowAndRun()

}

func Countdown() {
	myApp := app.New()
	myWindow := myApp.NewWindow("UIUI")

	buttonBeenden := widget.NewButton("Beenden und Bilder löschen.", func() {
		log.Println("Beenden Button gedrückt.")
	})

	circle := canvas.NewCircle(color.White)
	toptop := container.New(layout.NewGridWrapLayout(fyne.NewSize(250, 50)), buttonBeenden)

	top := container.New(layout.NewGridLayout(1), toptop, layout.NewSpacer())
	middle := container.New(layout.NewGridLayout(1), layout.NewSpacer(), circle, layout.NewSpacer())

	content := container.NewBorder(top, bottom, left, right, middle)
	myWindow.SetContent(content)
	myWindow.SetFullScreen(true)
	myWindow.ShowAndRun()
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

func ShowPic() {
	myApp := app.New()
	myWindow := myApp.NewWindow("UIUI")

	buttonBeenden := widget.NewButton("Beenden und Bilder löschen.", func() {
		log.Println("Beenden Button gedrückt.")
	})
	buttonNeuesBild := widget.NewButton("Weitere Szenarien auswählen.", func() {
		log.Println("Szenario gedrückt.")
	})

	pic := canvas.NewImageFromFile("C:\\Users\\yagmu\\GolandProjects\\inno-lab\\server\\internal\\ui\\blackstardustxx_ultrarealistic_picture_of_a_woman_heavily_tatto_251705d7-9e41-4c3c-b6ff-5f248ae53ba0.png")
	topleft := container.New(layout.NewGridWrapLayout(fyne.NewSize(250, 50)), buttonBeenden)
	gridB := container.New(layout.NewGridWrapLayout(fyne.NewSize(850, 50)), buttonNeuesBild)

	top := container.New(layout.NewGridLayout(1), topleft)
	imageGrid := container.New(layout.NewGridWrapLayout(fyne.NewSize(850, 850)), pic)
	gridC := container.New(layout.NewGridLayout(1), layout.NewSpacer(), imageGrid, gridB)
	middle := container.New(layout.NewCenterLayout(), gridC)

	content := container.NewBorder(top, bottom, left, right, middle)

	myWindow.SetContent(content)

	myWindow.ShowAndRun()
}
