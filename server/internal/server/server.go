package server

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Frank-Mayer/inno-lab/internal/schema"
	"github.com/charmbracelet/log"
	"github.com/go-vgo/robotgo"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	inputPosX = -1
	inputPosY = -1
)

var (
	promptCh     = make(chan string, 1)
	urlCh        = make(chan string, 1)
	expectsImage = atomic.Bool{}
)

func focusBack() {
	robotgo.MoveRelative(1791, -400)
}

func Reset() {
	inputPosX = -1
	inputPosY = -1
	expectsImage.Store(false)
	go func() {
		urlCh <- ""
	}()
	for i := 0; i != 2; i++ {
		select {
		case <-promptCh:
		default:
		}
		select {
		case <-urlCh:
		default:
		}
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// redirect to discord
	http.Redirect(w, r, "https://discord.com/channels/@me/1019912799977230416", http.StatusTemporaryRedirect)
}

func setInputPos(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("Error upgrading websocket", "error", err)
		return
	}
	defer c.Close()

	for {
		_, b, err := c.ReadMessage()
		if err != nil {
			log.Error("Error reading message", "error", err)
			break
		}
		pos := schema.Pos{}
		err = proto.Unmarshal(b, &pos)
		if err != nil {
			log.Error("Error unmarshalling message", "error", err)
			break
		}
		log.Debug("Received input position", "x", pos.X, "y", pos.Y)
		inputPosX = int(pos.X)
		inputPosY = int(pos.Y)
	}
}

func sendImage(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("Error upgrading websocket", "error", err)
		return
	}
	defer c.Close()

	for {
		_, b, err := c.ReadMessage()
		if err != nil {
			log.Error("Error reading message", "error", err)
			break
		}

		// ping?
		if string(b) == "ping" {
			log.Debug("send_image ping")
			continue
		}

		res := schema.Res{}
		err = proto.Unmarshal(b, &res)
		if err != nil {
			log.Error("Error unmarshalling message", "error", err)
			break
		}
		tm := time.Unix(res.Time, 0)
		log.Info("Received image",
			"time", tm.String(),
			"prompt", res.Prompt,
			"image", res.Src,
		)
		if strings.Contains(res.Prompt, "π") {
			log.Info("Received image", "exponat", "2", "image", res.Src)
			storeImage(res.Src)
		} else if expectsImage.Load() {
			log.Info("Received image", "exponat", "1", "image", res.Src)
			urlCh <- res.Src
			log.Debug("Wrote to result channel")
			expectsImage.Store(false)
		} else {
			log.Error("Received image (throwing away)", "exponat", "N/A", "image", res.Src)
		}
	}
}

var chromeWriteMut = sync.Mutex{}

func processPrompt(prompt string) {
	_ = chromeWriteMut.TryLock()
	inputPosX := inputPosX + 50
	inputPosY := inputPosY + 20
	log.Info("Processing prompt", "prompt", prompt)
	robotgo.KeySleep = 300
	robotgo.MouseSleep = 200
	log.Debug("move click", "x", inputPosX, "y", inputPosY)
	robotgo.Move(inputPosX, inputPosY)
	robotgo.MoveClick(inputPosX, inputPosY)
	robotgo.MoveClick(inputPosX, inputPosY)
	log.Debug("sleep")
	time.Sleep(1 * time.Second)
	log.Debug("type")
	TypeStr("/imagine "+prompt, 0, 25)
	err := robotgo.KeyTap("enter")
	if err != nil {
		log.Error("Error sending enter", "error", err)
	}
	focusBack()
}

func TypeStr(str string, args ...int) {
	pid := 0
	tm := 10
	if len(args) > 0 {
		pid = args[0]
	}
	if len(args) > 1 {
		tm = args[1]
	}
	for i := 0; i < len([]rune(str)); i++ {
		ustr := uint32(robotgo.CharCodeAt(str, i))
		robotgo.UnicodeType(ustr, pid)
		<-time.After(time.Duration(tm) * time.Millisecond)
	}
}

var (
	images = make(map[int]string)
	count  = 3
	index  = 0
)

func htmlImage(index int) string {
	return `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Veritas</title>
</head>
<body>
    <div id="image" style="background-image: url('https://raw.githubusercontent.com/Frank-Mayer/inno-lab/main/logo.png'); background-size: contain; background-repeat: no-repeat; background-position: center center; width: 100vw; height: 100vh;"></div>
    <style>
        :root {
            background-color: black;
        }
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
            cursor: none;
        } 
    </style>
    <script>
    window.addEventListener("click", function() {
        document.documentElement.requestFullscreen({navigationUI: "hide"});
    })
    async function updateImage() {
        const res = await fetch("/img_src/` + fmt.Sprintf("%d", index) + `");
        const src = await res.text();
        console.log(src);
        const divEl = document.getElementById("image");
        divEl.style.backgroundImage = "url('" + src + "')";
    }
    window.setInterval(updateImage, 1000);
    </script>
</body>
</html>

    `

}

func createImageHtmlHandler(i int) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		if _, err := w.Write([]byte(htmlImage(i))); err != nil {
			log.Error("Error writing image", "error", err)
		}
	}
}

func createImageSrcHandler(i int) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		src, ok := images[i]
		if !ok || src == "" {
			src = "https://raw.githubusercontent.com/Frank-Mayer/inno-lab/main/logo.png"
		}
		if _, err := w.Write([]byte(src)); err != nil {
			log.Error("Error writing image", "error", err)
		}
	}
}

func storeImage(src string) {
	log.Debug("store image", "src", src)
	images[index] = src
	index = (index + 1) % count
}

func Init() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/set_input_pos", setInputPos)
	http.HandleFunc("/send_image", sendImage)
	for i := 0; i < count; i++ {
		http.HandleFunc(fmt.Sprintf("/image/%d", i), createImageHtmlHandler(i))
		http.HandleFunc(fmt.Sprintf("/img_src/%d", i), createImageSrcHandler(i))
	}

	// Start the server on port 8080
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Error("Server error", "error", err)
			return
		}
	}()

	go func() {
		for {
			prompt := <-promptCh
			if inputPosX != -1 && inputPosY != -1 {
				log.Info("Input position set", "x", inputPosX, "y", inputPosY)
				processPrompt(prompt)
				time.Sleep(1 * time.Second)
				chromeWriteMut.Unlock()
			} else {
				log.Error("Input position not set")
			}
		}
	}()
}

func SendBackgroundPrompt(prompt string) {
	chromeWriteMut.Lock()
	promptCh <- prompt
	chromeWriteMut.Lock()
	chromeWriteMut.Unlock()
}

func SentPrompt(prompt string) chan string {
	chromeWriteMut.Lock()
	promptCh <- prompt
	expectsImage.Store(true)
	chromeWriteMut.Lock()
	chromeWriteMut.Unlock()
	return urlCh
}
