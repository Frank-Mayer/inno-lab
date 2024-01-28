package keyboard

import (
	"errors"
	"runtime"
	"sync"
	"time"

	"github.com/micmonay/keybd_event"
	"golang.design/x/clipboard"
)

func SendKeys(text string) error {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		return errors.Join(errors.New("could not create keybd_event"), err)
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		// For linux, it is very important to wait 2 seconds
		if runtime.GOOS == "linux" {
			<-time.After(2 * time.Second)
		}
		wg.Done()
	}()

	go func() {
		<-clipboard.Write(clipboard.FmtText, []byte(text))
		wg.Done()
	}()

	wg.Wait()

	// set keys
	kb.HasCTRL(true)
	kb.SetKeys(keybd_event.VK_V)
	if err := kb.Launching(); err != nil {
		return errors.Join(errors.New("could not launch keybd_event"), err)
	}
	return nil
}

// Insert simulates Ctrl+V
func Insert() error {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	// For linux, it is very important to wait 2 seconds
	if runtime.GOOS == "linux" {
		<-time.After(2 * time.Second)
	}

	// set keys
	kb.HasCTRL(true)
	kb.SetKeys(keybd_event.VK_V)
	err = kb.Launching()
	if err != nil {
		return err
	}
	kb.HasCTRL(false)
	return nil
}

func runeToVk(r rune) (int, bool) {
	switch r {
	case 'a':
		return keybd_event.VK_A, false
	case 'b':
		return keybd_event.VK_B, false
	case 'c':
		return keybd_event.VK_C, false
	case 'd':
		return keybd_event.VK_D, false
	case 'e':
		return keybd_event.VK_E, false
	case 'f':
		return keybd_event.VK_F, false
	case 'g':
		return keybd_event.VK_G, false
	case 'h':
		return keybd_event.VK_H, false
	case 'i':
		return keybd_event.VK_I, false
	case 'j':
		return keybd_event.VK_J, false
	case 'k':
		return keybd_event.VK_K, false
	case 'l':
		return keybd_event.VK_L, false
	case 'm':
		return keybd_event.VK_M, false
	case 'n':
		return keybd_event.VK_N, false
	case 'o':
		return keybd_event.VK_O, false
	case 'p':
		return keybd_event.VK_P, false
	case 'q':
		return keybd_event.VK_Q, false
	case 'r':
		return keybd_event.VK_R, false
	case 's':
		return keybd_event.VK_S, false
	case 't':
		return keybd_event.VK_T, false
	case 'u':
		return keybd_event.VK_U, false
	case 'v':
		return keybd_event.VK_V, false
	case 'w':
		return keybd_event.VK_W, false
	case 'x':
		return keybd_event.VK_X, false
	case 'y':
		return keybd_event.VK_Y, false
	case 'z':
		return keybd_event.VK_Z, false
	case 'A':
		return keybd_event.VK_A, true
	case 'B':
		return keybd_event.VK_B, true
	case 'C':
		return keybd_event.VK_C, true
	case 'D':
		return keybd_event.VK_D, true
	case 'E':
		return keybd_event.VK_E, true
	case 'F':
		return keybd_event.VK_F, true
	case 'G':
		return keybd_event.VK_G, true
	case 'H':
		return keybd_event.VK_H, true
	case 'I':
		return keybd_event.VK_I, true
	case 'J':
		return keybd_event.VK_J, true
	case 'K':
		return keybd_event.VK_K, true
	case 'L':
		return keybd_event.VK_L, true
	case 'M':
		return keybd_event.VK_M, true
	case 'N':
		return keybd_event.VK_N, true
	case 'O':
		return keybd_event.VK_O, true
	case 'P':
		return keybd_event.VK_P, true
	case 'Q':
		return keybd_event.VK_Q, true
	case 'R':
		return keybd_event.VK_R, true
	case 'S':
		return keybd_event.VK_S, true
	case 'T':
		return keybd_event.VK_T, true
	case 'U':
		return keybd_event.VK_U, true
	case 'V':
		return keybd_event.VK_V, true
	case 'W':
		return keybd_event.VK_W, true
	case 'X':
		return keybd_event.VK_X, true
	case 'Y':
		return keybd_event.VK_Y, true
	case 'Z':
		return keybd_event.VK_Z, true
	case '0':
		return keybd_event.VK_0, false
	case '1':
		return keybd_event.VK_1, false
	case '2':
		return keybd_event.VK_2, false
	case '3':
		return keybd_event.VK_3, false
	case '4':
		return keybd_event.VK_4, false
	case '5':
		return keybd_event.VK_5, false
	case '6':
		return keybd_event.VK_6, false
	case '7':
		return keybd_event.VK_7, false
	case '8':
		return keybd_event.VK_8, false
	case '9':
		return keybd_event.VK_9, false
	case '!':
		return keybd_event.VK_1, true
	case '"':
		return keybd_event.VK_2, true
	case '§':
		return keybd_event.VK_3, true
	case '$':
		return keybd_event.VK_4, true
	case '%':
		return keybd_event.VK_5, true
	case '&':
		return keybd_event.VK_6, true
	case '/':
		return keybd_event.VK_7, true
	case '(':
		return keybd_event.VK_8, true
	case ')':
		return keybd_event.VK_9, true
	case '=':
		return keybd_event.VK_0, true
	case ' ':
		return keybd_event.VK_SPACE, false
	case '-':
		return keybd_event.VK_SP11, false
	case '_':
		return keybd_event.VK_SP11, true
	case '.':
		return keybd_event.VK_SP10, false
	case ':':
		return keybd_event.VK_SP10, true
	case ',':
		return keybd_event.VK_SP9, false
	case ';':
		return keybd_event.VK_SP9, true
	case 'ä':
		return keybd_event.VK_SP7, false
	case 'Ä':
		return keybd_event.VK_SP7, true
	case 'ö':
		return keybd_event.VK_SP6, false
	case 'Ö':
		return keybd_event.VK_SP6, true
	case 'ü':
		return keybd_event.VK_SP4, false
	case 'Ü':
		return keybd_event.VK_SP4, true
	case 'ß':
		return keybd_event.VK_SP2, false
	case '?':
		return keybd_event.VK_SP2, true
	default:
		panic("Unknown rune " + string(r))
	}
}
