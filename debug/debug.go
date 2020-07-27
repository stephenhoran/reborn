package debug

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	messageYOffset = 15
)

type Debugger struct {
	messages []string
}

func NewDebugger() *Debugger {
	return &Debugger{
		messages: make([]string, 0),
	}
}

func (d *Debugger) AddMessage(message string) {
	d.messages = append(d.messages, message)
}

func (d *Debugger) Draw(screen *ebiten.Image) {
	var y int

	for _, message := range d.messages {
		ebitenutil.DebugPrintAt(screen, message, 0, y)
		y += messageYOffset
	}

	d.messages = make([]string, 0)
}
