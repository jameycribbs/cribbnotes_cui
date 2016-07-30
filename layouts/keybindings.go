package layouts

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func keybindingsLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("help", maxX-32, 2, maxX, maxY-25); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "[ Keybindings ]"

		fmt.Fprintln(v, "[Down/Up]       - scroll line")
		fmt.Fprintln(v, "[j/k]           - scroll line")
		fmt.Fprintln(v, "")

		fmt.Fprintln(v, "[PgDown/PgUp]   - scroll page")
		fmt.Fprintln(v, "[Ctrl+f/Ctrl+b] - scroll page")
		fmt.Fprintln(v, "")

		fmt.Fprintln(v, "[Ctrl+Spacebar] - switch views")
		fmt.Fprintln(v, "[Ctrl+j/Ctrl+k] - switch views")
		fmt.Fprintln(v, "")

		fmt.Fprintln(v, "[/]             - find notes")
		fmt.Fprintln(v, "[Ctrl+n]        - new note")
		fmt.Fprintln(v, "[Ctrl+s]        - save note")
		fmt.Fprintln(v, "[Ctrl+d]        - delete note")
		fmt.Fprintln(v, "[Ctrl+q]        - quit")
	}

	return nil
}
