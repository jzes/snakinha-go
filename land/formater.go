package land

import (
	"bytes"
	"io"
	"strings"

	"github.com/gdamore/tcell"
)

type TCellFmt struct {
	Screen tcell.Screen
	Style  tcell.Style
}

func (f TCellFmt) Print(w io.Writer) error {
	land := strings.Split(w.(*bytes.Buffer).String(), "\n")
	for y, line := range land {
		for x, char := range line {
			f.Screen.SetContent(x, y, rune(char), nil, f.Style)
		}
	}

	return nil
}
