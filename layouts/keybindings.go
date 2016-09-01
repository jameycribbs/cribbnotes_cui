package layouts

import (
	"errors"
	"fmt"

	"github.com/jameycribbs/cribbnotes_cui/config"
	"github.com/jroimartin/gocui"
)

func keybindingsLayout(g *gocui.Gui) error {
	var err error
	var maxX, maxY int
	var v *gocui.View

	maxX, maxY = g.Size()

	if v, err = g.SetView("help", maxX-34, 2, maxX, maxY-25); err != nil {
		if err != gocui.ErrUnknownView {
			return errors.New("(keybindingsLayout) error setting view: " + err.Error())
		}

		v.Title = "[ Keybindings ]"

		if config.VimMode {
			fmt.Fprintln(v, "[j/k]           - scroll line")
			fmt.Fprintln(v, "[Ctrl+f/Ctrl+b] - scroll page")
			fmt.Fprintln(v, "[Ctrl+j/Ctrl+k] - switch views")
			fmt.Fprintln(v, "[i]             - Insert mode")
			fmt.Fprintln(v, "[Esc]           - Command mode")
			fmt.Fprintln(v, "[:]             - Ex mode")
			fmt.Fprintln(v, "[x]             - Delete Char")
			fmt.Fprintln(v, "[F3]            - find notes")
			fmt.Fprintln(v, "[F4]            - clear filter")
			fmt.Fprintln(v, "[Ctrl+i]        - new note")
			fmt.Fprintln(v, "[Ctrl+d]        - delete note")
			fmt.Fprintln(v, "[Ctrl+t]        - edit title")
		} else {
			fmt.Fprintln(v, "[Down/Up]       - scroll line")
			fmt.Fprintln(v, "[PgDown/PgUp]   - scroll page")
			fmt.Fprintln(v, "[F3]            - find notes")
			fmt.Fprintln(v, "[F4]            - clear filter")
			fmt.Fprintln(v, "[F6]            - switch views")
			fmt.Fprintln(v, "[Ctrl+n]        - new note")
			fmt.Fprintln(v, "[Ctrl+s]        - save note")
			fmt.Fprintln(v, "[Ctrl+d]        - delete note")
			fmt.Fprintln(v, "[Ctrl+t]        - edit title")
			fmt.Fprintln(v, "[Ctrl+q]        - quit")
		}
	}

	return nil
}
