package figure

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"

	ct "github.com/daviddengcn/go-colortext"
	_ "github.com/godoes/winseq" // Use Unix like Sequences in Windows
)

// IsOldWindows is older Windows OS below Windows 10
var IsOldWindows bool

// Print stdout
func (fig Figure) Print() {
	colorful := fig.color != ""
	for _, printRow := range fig.Slicify() {
		if colorful {
			if IsOldWindows {
				ct.Foreground(ctColors[fig.color], true)
			} else {
				printRow = colors[fig.color] + printRow + colors["reset"]
			}
		}
		fmt.Println(printRow)
		if colorful && IsOldWindows {
			ct.ResetColor()
		}
	}
}

// ColorString returns a colored string
func (fig Figure) ColorString() string {
	s := ""
	for _, printRow := range fig.Slicify() {
		if fig.color != "" {
			printRow = colors[fig.color] + printRow + colors["reset"]
		}
		s += fmt.Sprintf("%s\n", printRow)
	}
	return s
}

func (fig Figure) String() string {
	s := ""
	for _, printRow := range fig.Slicify() {
		s += fmt.Sprintf("%s\n", printRow)
	}
	return s
}

// Scroll A figure responds to the func Scroll, taking three arguments.
// duration is the total time the banner will display, in milliseconds.
// stillness is the length of time the text will not move (also in ms).
// Therefore, the lower the stillness the faster the scroll speed.
// direction can be either "right" or "left" (case-insensitive).
// The direction will be left if an invalid option (e.g. "foo") is passed.
// There is no return value.
func (fig Figure) Scroll(duration, stillness int, direction string) {
	endTime := time.Now().Add(time.Duration(duration) * time.Millisecond)
	fig.phrase = fig.phrase + "   "
	clearScreen()
	for time.Now().Before(endTime) {
		var shiftedPhrase string
		chars := []byte(fig.phrase)
		if strings.HasPrefix(strings.ToLower(direction), "r") {
			shiftedPhrase = string(append(chars[len(chars)-1:], chars[0:len(chars)-1]...))
		} else {
			shiftedPhrase = string(append(chars[1:], chars[0]))
		}
		fig.phrase = shiftedPhrase
		fig.Print()
		sleep(stillness)
		clearScreen()
	}
}

// Blink A figure responds to the func Blink, taking three arguments.
// duration is the total time the banner will display, in milliseconds.
// timeOn is the length of time the text will blink on (also in ms).
// timeOff is the length of time the text will blink off (ms).
// For an even blink, set timeOff to -1 (same as setting timeOff to the value of timeOn).
// There is no return value.
func (fig Figure) Blink(duration, timeOn, timeOff int) {
	if timeOff < 0 {
		timeOff = timeOn
	}
	endTime := time.Now().Add(time.Duration(duration) * time.Millisecond)
	clearScreen()
	for time.Now().Before(endTime) {
		fig.Print()
		sleep(timeOn)
		clearScreen()
		sleep(timeOff)
	}
}

// Dance A figure responds to the func Dance, taking two arguments.
// duration is the total time the banner will display, in milliseconds.
// freeze is the length of time between dance moves (also in ms).
// Therefore, the lower the freeze the faster the dancing.
// There is no return value.
func (fig Figure) Dance(duration, freeze int) {
	endTime := time.Now().Add(time.Duration(duration) * time.Millisecond)
	font := fig.font //TODO: change to deep copy
	font.evenLetters()
	figures := []Figure{{font: font}, {font: font}}
	clearScreen()
	for i, c := range fig.phrase {
		appends := []string{" ", " "}
		appends[i%2] = string(c)
		for f := range figures {
			figures[f].phrase = figures[f].phrase + appends[f]
		}
	}
	for p := 0; time.Now().Before(endTime); p ^= 1 {
		figures[p].Print()
		figures[1-p].Print()
		sleep(freeze)
		clearScreen()
	}
}

// Write writers
func Write(w io.Writer, fig Figure) {
	for _, printRow := range fig.Slicify() {
		_, _ = fmt.Fprintf(w, "%v\n", printRow)
	}
}

// helpers
func clearScreen() {
	if IsOldWindows {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	} else {
		fmt.Print("\033[H\033[2J")
	}
}

func sleep(milliseconds int) {
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)
}
