package ui

import (
	"fmt"
	"github.com/Frank-Mayer/inno-lab/internal/firebase"
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
	"github.com/Frank-Mayer/inno-lab/internal/server"
	"github.com/charmbracelet/log"
	"github.com/pkg/errors"
)

var (
	logo    *canvas.Image
	logohhn *canvas.Image
	logohp  *canvas.Image
	bottom  *fyne.Container
	flow    *fyneflow.Flow
)

func Init() fyne.Window {
	rsc, err := fyne.LoadResourceFromURLString("https://raw.githubusercontent.com/Frank-Mayer/inno-lab/main/logo.png")
	if err != nil {
		panic(err)
	}
	logo = canvas.NewImageFromResource(rsc)

	rschhn, err := fyne.LoadResourceFromURLString("https://raw.githubusercontent.com/Frank-Mayer/inno-lab/main/HHN_logo-modified-removebg-preview.png")
	if err != nil {
		panic(err)
	}
	logohhn = canvas.NewImageFromResource(rschhn)

	rschp, err := fyne.LoadResourceFromURLString("https://raw.githubusercontent.com/Frank-Mayer/inno-lab/main/hochschule-pforzheim-hs-pf-logo-vector-01.png")
	if err != nil {
		panic(err)
	}
	logohp = canvas.NewImageFromResource(rschp)

	txtfrk := canvas.NewText("Franziska", color.White)
	txtfrk.TextSize = 20
	txtfrk.Alignment = fyne.TextAlignCenter
	txtfrk.TextStyle = fyne.TextStyle{Monospace: true}

	txtfrz := canvas.NewText("Frank", color.White)
	txtfrz.TextSize = 20
	txtfrz.Alignment = fyne.TextAlignCenter
	txtfrz.TextStyle = fyne.TextStyle{Monospace: true}

	txtmm := canvas.NewText("Maria-Magdalena", color.White)
	txtmm.TextSize = 20
	txtmm.Alignment = fyne.TextAlignCenter
	txtmm.TextStyle = fyne.TextStyle{Monospace: true}

	txtym := canvas.NewText("Yagmur Mine", color.White)
	txtym.TextSize = 20
	txtym.Alignment = fyne.TextAlignCenter
	txtym.TextStyle = fyne.TextStyle{Monospace: true}

	var logoGrid fyne.CanvasObject = container.New(layout.NewCenterLayout(), container.New(layout.NewGridWrapLayout(fyne.NewSize(200, 75)), logo), layout.NewSpacer())
	var hhnGrid fyne.CanvasObject = container.New(layout.NewCenterLayout(), container.New(layout.NewGridWrapLayout(fyne.NewSize(175, 100)), logohhn), layout.NewSpacer())
	var hpGrid fyne.CanvasObject = container.New(layout.NewCenterLayout(), container.New(layout.NewGridWrapLayout(fyne.NewSize(175, 100)), logohp), layout.NewSpacer())
	bottom = container.New(layout.NewGridLayout(9), hpGrid, layout.NewSpacer(), txtfrk, txtfrz, logoGrid, txtmm, txtym, layout.NewSpacer(), hhnGrid)

	myApp := app.New()
	myWindow := myApp.NewWindow("UIUI")

	flow = fyneflow.NewFlow(myWindow)
	flow.Set("CreateUI", CreateUI)
	flow.Set("CreateScenario", CreateScenario)
	flow.Set("CameraLook", CameraLook)
	flow.Set("ShowPic", ShowPic)
	flow.Set("Loading", Loading)
	flow.Set("error", UiUiError)
	flow.Set("LoadingInfo", LoadingInfo)
	flow.Set("c1", C1)
	flow.Set("c2", C2)
	flow.Set("c3", C3)

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
	t := time.NewTimer(5 * time.Minute)
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
	text1 := canvas.NewText("Bevor es losgeht... Wie möchten Sie angesprochen werden?", color.White)
	text1.TextSize = 30
	text1.Alignment = fyne.TextAlignCenter
	text1.TextStyle = fyne.TextStyle{Monospace: true}

	text2 := canvas.NewText(" ", color.White)
	text2.TextSize = 30

	button1 := widget.NewButton("Frau", func() {
		_ = flow.UseStateStr("gender", "Woman").Set("Woman")
		_ = flow.GoTo("CameraLook")
	})

	button2 := widget.NewButton("Herr", func() {
		_ = flow.UseStateStr("gender", "Man").Set("Man")
		_ = flow.GoTo("CameraLook")
	})
	//Non-Binary Button?

	middle := container.New(layout.NewGridLayout(
		1), text1, container.New(layout.NewCenterLayout(),
		container.New(layout.NewGridLayout(2),
			container.New(layout.NewGridWrapLayout(fyne.NewSize(600, 400)), button1),
			container.New(layout.NewGridWrapLayout(fyne.NewSize(600, 400)), button2))))

	textGrid := container.New(layout.NewGridWrapLayout(fyne.NewSize(70, 70)), text2)

	stack := container.NewStack(

		container.NewBorder(textGrid,
			bottom, textGrid, textGrid,
			container.NewVBox(middle)))

	return stack
}

func genderVeriation(male string, female string) string {
	gender, err := flow.UseStateStr("gender", "").Get()
	if err != nil {
		displayError(errors.Wrap(err, "failed to read state for gender"))
		return ""
	}
	switch gender {
	case "Man":
		return male
	case "Woman":
		return female
	}
	displayError(fmt.Errorf("invalid gender '%s'", gender))
	return ""
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

	button1 := widget.NewButton(genderVeriation("Politiker", "Politikerin"), func() {
		_ = flow.GoTo("LoadingInfo")
		cancelTo := timeout("generating politician picture")
		go func() {
			promptString := webcamUrl + "  ::4 https://l.frankmayer.dev/v_politik ::1 photographic picture of a  " + gender + "  as a politican in the Bundestag "
			resUrlChan := server.SentPrompt(promptString)
			expo2Gen(webcamUrl)
			if err := promptResult.Set(<-resUrlChan); err != nil {
				log.Error("Failed to set prompt result", "error", err)
			}
			_ = flow.GoTo("ShowPic")
			cancelTo()
		}()
	})
	button2 := widget.NewButton(genderVeriation("Astronaut", "Astronautin"), func() {
		_ = flow.GoTo("LoadingInfo")
		cancelTo := timeout("generating astronaut picture")
		go func() {
			promptString := webcamUrl + " ::4 https://l.frankmayer.dev/v_astronaut ::1 ultrarealistic picture of a " + gender + " as an astronaut in a space suit, stars and planets in the background "
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
		_ = flow.GoTo("LoadingInfo")
		cancelTo := timeout("generating holiday picture")
		go func() {
			promptString := webcamUrl + "::4 https://l.frankmayer.dev/v_urlaub ::1 ultrarealistic picture of a " + gender + " at a holiday resort. warm tones --v 5.2 "
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
		_ = flow.GoTo("LoadingInfo")
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
		_ = flow.GoTo("LoadingInfo")
		cancelTo := timeout("generating popstar picture")
		go func() {
			promptString := webcamUrl + " ::5 https://l.frankmayer.dev/v_popstar1 ::2 https://l.frankmayer.dev/v_popstar2 ::3 ultrarealistic picture of a " + gender + " as a popstar on the concert stage, microphone in the hand, crowd cheering  "
			resChan := server.SentPrompt(promptString)
			expo2Gen(webcamUrl)
			if err := promptResult.Set(<-resChan); err != nil {
				log.Error("Failed to set prompt result", "error", err)
			}
			_ = flow.GoTo("ShowPic")
			cancelTo()
		}()
	})
	button6 := widget.NewButton(genderVeriation("Fußballer", "Fußballerin"), func() {
		_ = flow.GoTo("LoadingInfo")
		cancelTo := timeout("generating football player picture")
		go func() {
			promptString := webcamUrl + " https://l.frankmayer.dev/v_fußball photographic of a " + gender + " as a football player in a press conference, logos in the background, wearing a jersey --iw 2"
			resChan := server.SentPrompt(promptString)
			expo2Gen(webcamUrl)
			if err := promptResult.Set(<-resChan); err != nil {
				log.Error("Failed to set prompt result", "error", err)
			}
			_ = flow.GoTo("ShowPic")
			cancelTo()
		}()
	})
	button7 := widget.NewButton(genderVeriation("Abenteurer", "Abenteurerin"), func() {
		_ = flow.GoTo("LoadingInfo")
		cancelTo := timeout("generating adventure picture")
		go func() {
			promptString := webcamUrl + " ::5 https://l.frankmayer.dev/v_jungle ::1 ultrarealistic picture of a  " + gender + "  outdoors in a jungle sitting in the trees, beige clothes and a hat "
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
		_ = flow.GoTo("LoadingInfo")
		cancelTo := timeout("generating model picture")
		go func() {
			promptString := webcamUrl + " ::4 https://l.frankmayer.dev/v_model ::1 photographic picture of a " + gender + " as a model wearing haute couture on the catwalk "
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
		_ = flow.GoTo("LoadingInfo")
		cancelTo := timeout("generating ted talk picture")
		go func() {
			promptString := webcamUrl + " ::4 https://l.frankmayer.dev/v_ted ::1 Create a photograph of a " + gender + " holding a motivational speech at a TEDtalk. They are standing on a red, round carpet. There are people sitting in the crowd. \"TED\" --v 6.0 "
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

	text3 := canvas.NewText(" ", color.White)
	text3.TextSize = 30

	top := container.New(layout.NewGridLayout(1), layout.NewSpacer(), text1)

	middle := container.New(layout.NewGridLayout(1), layout.NewSpacer(),
		container.New(layout.NewGridLayout(3), button1, button2, button3, button4, button5, button6, button7, button8, button9), layout.NewSpacer())

	textGrid := container.New(layout.NewGridWrapLayout(fyne.NewSize(70, 70)), text3)

	stack := container.NewStack(
		container.NewHBox(container.NewVBox(buttonBeenden)),
		container.NewBorder(top, bottom, textGrid, textGrid,
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

func LoadingInfo() fyne.CanvasObject {
	text1 := canvas.NewText("Bitte warten...", color.White)
	text1.TextSize = 50
	text1.Alignment = fyne.TextAlignCenter
	text1.TextStyle = fyne.TextStyle{Monospace: true}

	text2 := canvas.NewText("Fass mich bitte nicht an, ich bin empfindlich.", color.RGBA{R: 255, G: 0, B: 0, A: 255})
	text2.TextSize = 20
	text2.Alignment = fyne.TextAlignCenter
	text2.TextStyle = fyne.TextStyle{Monospace: true}

	var textForString string

	switch rand.Intn(9) {
	case 0:
		textForString = "Wussten Sie, dass an der Entwicklung dieses Projektes Studenten der Hochschule Heilbronn & Hochschule Pforzheim beteiligt waren? Unseren Namen:  Frank, Franziska, Maria und Yagmur."
	case 1:
		textForString = "Wussten Sie wie unsere KI-Bilder generiert werden? Von einem aus zufälligen Pixeln erstellten Bild, wird ein Foto konstruiert."
	case 2:
		textForString = "Wussten Sie die Bilder die Sie sehen, sind nach dem aktuellen Stand der Technik die besten Deepfakes die ohne zusätzliche Fachkenntnisse möglich sind ?  Schnell und Einfach also!"
	case 3:
		textForString = "Wussten Sie welche Eingaben wir benötigen um diese Bilder zu generieren? Ein Bild von Ihnen, ein Hintergrundbild (optional) und eine textuelle Beschreibung der Szene."
	case 4:
		textForString = "Wussten Sie worin sich unsere Deepfakes von denen die man im Internet sieht unterscheiden?  Sie stammen von geübten KI Profis und wurden manuell nachbearbeitet. Der gesamte Erstellungsprozess eines Deepfakes dauert Stunden und Tage."
	case 5:
		textForString = "Wussten Sie, dass der Begriff \"Deepfake\" 2017 auf Reddit geprägt wurde und sich aus den Wörtern \"deep\" (tief), wie in der KI-Tiefenlerntechnologie, und \"fake\" (gefälscht) zusammensetzt ?"
	case 6:
		textForString = "Wussten Sie professionelle Deepfakes können für verschiedene Zwecke verwendet werden, z. B. für Unterhaltung, Betrug oder Propaganda ?"
	case 7:
		textForString = "Wussten Sie, dass KI-Modelle, die mit Bildern trainiert werden, Vorurteile übernehmen können, wenn die Trainingsdaten selbst stereotypisch sind?"
	case 8:
		textForString = "Wussten Sie, dass einige Länder bereits Gesetze erlassen haben, die die Erstellung und Verbreitung von Deepfakes, die für böswillige Zwecke verwendet werden, verbieten?"
	case 9:
		textForString = "Wussten Sie, dass Deepfakes verwendet werden können, um Menschen in Situationen zu bringen, in denen sie sich nicht befinden (wollen)?"
	}

	text3 := canvas.NewText(textForString, color.White)
	text3.TextSize = 12
	text3.Alignment = fyne.TextAlignCenter
	text3.TextStyle = fyne.TextStyle{Monospace: true}

	text4 := canvas.NewText(" ", color.White)
	text4.TextSize = 60

	return container.NewBorder(
		nil, bottom, text4, text4,
		container.NewCenter(
			container.NewVBox(text1, text2, text4, text3),
		))
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
				" https://l.frankmayer.dev/v_2_1 photographic picture of a " + gender + " in front of the building π --iw 2")
	case 1:
		server.SendBackgroundPrompt(
			webcamUrl +
				" Swap the face. Create a picture of a " + gender + " walking through an art exhibition. They photographed from the side. There is colorful art and people in the background. π --s 0 --v 6.0")
	case 2:
		server.SendBackgroundPrompt(
			webcamUrl +
				" ::3 https://l.frankmayer.dev/v_2_0 ::1 Create an photograph of a " + gender + " sitting on the bus who is looking out of the window. The picture has warm tones, because the sun is shining. π --s 0 --v 6.0")
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

	text2 := canvas.NewText("Schauen Sie gerade in die Kamera links vom Bildschirm.", color.White)
	text2.TextSize = 20
	text2.Alignment = fyne.TextAlignCenter
	text2.TextStyle = fyne.TextStyle{Monospace: true}

	text3 := canvas.NewText("Drücken Sie die Taste, wenn Sie bereit sind!", color.White)
	text3.TextSize = 15
	text3.Alignment = fyne.TextAlignCenter
	text3.TextStyle = fyne.TextStyle{Monospace: true}

	button1 := widget.NewButton("Bereit?", func() {
		_ = flow.GoTo("c3")
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

func C1() fyne.CanvasObject {
	text := canvas.NewText("1", color.White)
	text.TextSize = 100
	go func() {
		<-time.After(time.Second)
		webcamUrlBind := flow.UseStateStr("webcamUrl", "")
		_ = flow.GoTo("Loading")
		if webcamUrl, err := firebase.GetWebcamUrl(); err == nil {
			if err := webcamUrlBind.Set(webcamUrl); err != nil {
				displayError(err)
				return
			} else {
				log.Debug("Webcam URL", "url", webcamUrl)
			}
		} else {
			displayError(err)
			return
		}
		_ = flow.GoTo("CreateScenario")

	}()
	return container.NewCenter(
		text,
	)
}

func C2() fyne.CanvasObject {
	text := canvas.NewText("2", color.White)
	text.TextSize = 100
	go func() {
		<-time.After(time.Second)
		_ = flow.GoTo("c1")
	}()
	return container.NewCenter(
		text,
	)
}

func C3() fyne.CanvasObject {
	text := canvas.NewText("3", color.White)
	text.TextSize = 100
	go func() {
		<-time.After(time.Second)
		_ = flow.GoTo("c2")
	}()
	return container.NewCenter(
		text,
	)
}
