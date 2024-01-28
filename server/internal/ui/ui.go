package ui

import (
	"fmt"
	"image/color"
	"math/rand"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Frank-Mayer/fyneflow"
	"github.com/Frank-Mayer/inno-lab/internal/firebase"
	"github.com/Frank-Mayer/inno-lab/internal/server"
	"github.com/charmbracelet/log"
	"github.com/pkg/errors"
)

var (
	logo   *canvas.Image
	bottom *fyne.Container
	flow   *fyneflow.Flow[string]
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
	flow.Set("Loading", Loading)
	flow.Set("error", UiUiError)

	return myWindow
}

func UiUiError() fyne.CanvasObject {
	err, _ := flow.UseStateStr("error", "").Get()
	log.Debug("Displaying error ui", "error", err)
	vbox := container.NewVBox(
		canvas.NewText("Da lief was schief :(", color.RGBA{R: 255, G: 0, B: 0, A: 255}),
		canvas.NewText("Bitte gib dem Personal bescheid", color.White),
	)
	if err != "" {
		// split error message into multiple lines
		for _, line := range strings.Split(err, "\n") {
			vbox.Add(canvas.NewText(line, color.White))
		}
	}
	return container.NewBorder(nil, bottom, nil, nil, vbox)
}

// timeout waits and then displays the error ui.
// It returns a function that can be called to cancel the timeout.
func timeout(action string) func() {
	// timeout after 10 seconds
	t := time.NewTimer(2 * time.Minute)
	go func() {
		<-t.C
		displayError(errors.Errorf("Timeout: %s", action))
	}()
	return func() {
		t.Stop()
	}
}

func renderError(err error) fyne.CanvasObject {
	if err := flow.UseStateStr("error", "").Set(fmt.Sprint(err)); err != nil {
		log.Error("Failed to set error", "error", err)
	}
	return UiUiError()
}

func displayError(err error) {
	if err == nil {
		log.Error("called displayError with nil error")
		return
	}
	log.Error("Error", "error", err)
	if err := flow.UseStateStr("error", "").Set(err.Error()); err != nil {
		log.Error("Failed to set error", "error", err)
	}
	_ = flow.GoTo("error")
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

	gender, err := flow.UseStateStr("gender", "").Get()
	if err != nil {
		return renderError(err)
	}
	promptResult := flow.UseStateStr("prompt", "")
	webcamUrl, err := flow.UseStateStr("webcamUrl", "").Get()
	if err != nil {
		return renderError(err)
	}

	button1 := widget.NewButton("Politiker:in", func() {
		_ = flow.GoTo("Loading")
		cancelTo := timeout("generating politician picture")
		go func() {
			promptString := webcamUrl + " https://s.mj.run/hWqyh5IpNio photographic picture of a " + gender + " as a politican in a press conference --iw 2"
			resUrlChan := server.SentPrompt(promptString)
			expo2Gen(webcamUrl)
			if err := promptResult.Set(<-resUrlChan); err != nil {
				log.Error("Failed to set prompt result", "error", err)
			}
			_ = flow.GoTo("ShowPic")
			cancelTo()
		}()
	})
	button2 := widget.NewButton("Astronaut:in", func() {
		_ = flow.GoTo("Loading")
		cancelTo := timeout("generating astronaut picture")
		go func() {
			promptString := webcamUrl + " ::4 https://s.mj.run/QY2Z4ddoxck ::1 ultrarealistic picture of a " + gender + " as an astronaut in a space suit, stars and planets in the background "
			resChan := server.SentPrompt(promptString)
			expo2Gen(webcamUrl)
			if err := promptResult.Set(<-resChan); err != nil {
				log.Error("Failed to set prompt result", "error", err)
			}
			_ = flow.GoTo("ShowPic")
			cancelTo()
		}()
	})
	button3 := widget.NewButton("Im Urlaub", func() {
		_ = flow.GoTo("Loading")
		cancelTo := timeout("generating holiday picture")
		go func() {
			promptString := webcamUrl + "::4 https://s.mj.run/wb3_lBMHQZw ::1 ultrarealistic picture of a " + gender + " at a holiday resort. warm tones --v 5.2 "
			resChan := server.SentPrompt(promptString)
			expo2Gen(webcamUrl)
			if err := promptResult.Set(<-resChan); err != nil {
				log.Error("Failed to set prompt result", "error", err)
			}
			_ = flow.GoTo("ShowPic")
			cancelTo()
		}()
	})
	button4 := widget.NewButton("Imagewechsel", func() {
		_ = flow.GoTo("Loading")
		cancelTo := timeout("generating image change picture")
		go func() {
			promptString := webcamUrl + " ultrarealistic picture of a " + gender + " heavily tattood with piercings in their face, in the background you can see a tattoo parlor --iw 2 "
			resChan := server.SentPrompt(promptString)
			expo2Gen(webcamUrl)
			if err := promptResult.Set(<-resChan); err != nil {
				log.Error("Failed to set prompt result", "error", err)
			}
			_ = flow.GoTo("ShowPic")
			cancelTo()
		}()
	})
	button5 := widget.NewButton("Pop-Star", func() {
		_ = flow.GoTo("Loading")
		cancelTo := timeout("generating popstar picture")
		go func() {
			promptString := webcamUrl + " https://s.mj.run/DxMAyfGMFTY ultrarealistic picture of a " + gender + " as a popstar on the concert stage, microphone in their hand --iw 2 "
			resChan := server.SentPrompt(promptString)
			expo2Gen(webcamUrl)
			if err := promptResult.Set(<-resChan); err != nil {
				log.Error("Failed to set prompt result", "error", err)
			}
			_ = flow.GoTo("ShowPic")
			cancelTo()
		}()
	})
	button6 := widget.NewButton("Fußballer:in", func() {
		_ = flow.GoTo("Loading")
		cancelTo := timeout("generating football player picture")
		go func() {
			promptString := webcamUrl + " ::4 https://s.mj.run/J0aQ9gVwN6k ::1 photographic of a " + gender + " as a football player in a press conference, logos in the background"
			resChan := server.SentPrompt(promptString)
			expo2Gen(webcamUrl)
			if err := promptResult.Set(<-resChan); err != nil {
				log.Error("Failed to set prompt result", "error", err)
			}
			_ = flow.GoTo("ShowPic")
			cancelTo()
		}()
	})
	button7 := widget.NewButton("Auf Abenteuer", func() {
		_ = flow.GoTo("Loading")
		cancelTo := timeout("generating adventure picture")
		go func() {
			promptString := webcamUrl + " ::5 https://s.mj.run/_zpWilkBFyQ ::1 ultrarealistic picture of a " + gender + " in a jungle sitting in the trees, beige clothes and a hat "
			resChan := server.SentPrompt(promptString)
			expo2Gen(webcamUrl)
			if err := promptResult.Set(<-resChan); err != nil {
				log.Error("Failed to set prompt result", "error", err)
			}
			_ = flow.GoTo("ShowPic")
			cancelTo()
		}()
	})
	button8 := widget.NewButton("Model", func() {
		_ = flow.GoTo("Loading")
		cancelTo := timeout("generating model picture")
		go func() {
			promptString := webcamUrl + " ::4 https://s.mj.run/BzHMDLF1RhE ::1 photographic picture of a " + gender + " as a model wearing haute couture on the catwalk "
			resChan := server.SentPrompt(promptString)
			expo2Gen(webcamUrl)
			if err := promptResult.Set(<-resChan); err != nil {
				log.Error("Failed to set prompt result", "error", err)
			}
			_ = flow.GoTo("ShowPic")
			cancelTo()
		}()
	})
	button9 := widget.NewButton("Im TED-Talk", func() {
		_ = flow.GoTo("Loading")
		cancelTo := timeout("generating ted talk picture")
		go func() {
			promptString := webcamUrl + " ::4 https://s.mj.run/qcnxNbyX9hE ::1 Create a photograph of a " + gender + " holding a motivational speech at a TEDtalk. They are standing on a red, round carpet. There are people sitting in the crowd. \"TED\" --v 6.0 "
			resChan := server.SentPrompt(promptString)
			expo2Gen(webcamUrl)
			if err := promptResult.Set(<-resChan); err != nil {
				log.Error("Failed to set prompt result", "error", err)
			}
			_ = flow.GoTo("ShowPic")
			cancelTo()
		}()
	})

	buttonBeenden := widget.NewButton("Beenden und Bilder löschen.", func() {
		_ = flow.GoTo("CreateUI")
		_ = promptResult.Set("")
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

func Loading() fyne.CanvasObject {
	text1 := canvas.NewText("Bitte warten...", color.White)
	text1.TextSize = 50
	text1.Alignment = fyne.TextAlignCenter
	text1.TextStyle = fyne.TextStyle{Monospace: true}

	text2 := canvas.NewText("Fass mich bitte nicht an, ich bin empfindlich.", color.RGBA{R: 255, G: 0, B: 0, A: 255})
	text2.TextSize = 20
	text2.Alignment = fyne.TextAlignCenter
	text2.TextStyle = fyne.TextStyle{Monospace: true}

	return container.NewBorder(
		nil, bottom, nil, nil,
		container.NewCenter(
			container.NewVBox(text1, text2),
		),
	)
}

var lastPromptPhoto string

func expo2Gen(webcamUrl string) {
	if lastPromptPhoto == webcamUrl {
		return
	}
	lastPromptPhoto = webcamUrl
	genderState := flow.UseStateStr("gender", "")
	gender, err := genderState.Get()
	if err != nil {
		displayError(err)
	}
	switch rand.Intn(4) {
	case 0:
		server.SendBackgroundPrompt(
			webcamUrl +
				" https://s.mj.run/H2OCJUnS3T8 photographic picture of a " + gender + " in front of the building π --iw 2")
	case 1:
		server.SendBackgroundPrompt(
			webcamUrl +
				" Swap the face. Create a picture of a " + gender + " walking through an art exhibition. They photographed from the side. There is colorful art and people in the background. π --s 0 --v 6.0")
	case 2:
		server.SendBackgroundPrompt(
			webcamUrl +
				" ::3 https://s.mj.run/c17w4yFMpsc ::1 Create an photograph of a " + gender + " sitting on the bus who is looking out of the window. The picture has warm tones, because the sun is shining. π --s 0 --v 6.0")
	case 3:
		server.SendBackgroundPrompt(
			webcamUrl +
				" ::3 Create an photograph of a " + gender + " driving a car. They fastened their seatbelt. They are looking out of the window. The picture has warm tones, because the sun is shining. π --s 0 --v 6.0")
	}
}

func CameraLook() fyne.CanvasObject {
	buttonBeenden := widget.NewButton("Beenden und Bilder löschen.", func() {
		_ = flow.GoTo("CreateUI")
		//flow.ClearStates ?
	})

	rsc, err := fyne.LoadResourceFromURLString("https://raw.githubusercontent.com/Frank-Mayer/inno-lab/main/beispielPose.jpg")
	if err != nil {
		return renderError(err)
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
		webcamUrlBind := flow.UseStateStr("webcamUrl", "")
		if webcamUrl, err := firebase.GetWebcamUrl(); err == nil {
			if err := webcamUrlBind.Set(webcamUrl); err != nil {
				displayError(err)
				return
			}
		} else {
			displayError(err)
			return
		}
		_ = flow.GoTo("CreateScenario")
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
		_ = flow.GoTo("CreateUI")
		//flow.ClearStates ?

	})

	text1 := canvas.NewText("Bitte warten...", color.White)
	text1.TextSize = 50
	text1.Alignment = fyne.TextAlignCenter
	text1.TextStyle = fyne.TextStyle{Monospace: true}

	promptResult, err := flow.UseStateStr("prompt", "").Get()
	if err != nil {
		return renderError(err)
	}
	rsc, err := fyne.LoadResourceFromURLString(promptResult)
	if err != nil {
		return renderError(err)
	}

	pic := canvas.NewImageFromResource(rsc)

	stack := container.NewStack(
		container.NewHBox(container.NewVBox(buttonBeenden)),
		container.NewBorder(nil, bottom, nil, nil,
			container.NewCenter(text1,
				container.NewVBox(container.NewGridWrap(fyne.NewSize(620, 620), pic),
					widget.NewButton("Weitere Szenarien auswählen.", func() {
						_ = flow.GoTo("CreateScenario")
					})))))

	return stack
}
