package layouts

import (
	"log"

	"github.com/jameycribbs/cribbnotes_cui/kb_functions"
	"github.com/jroimartin/gocui"
)

func noteLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("note", -1, maxY-23, maxX, maxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		if err := kb_functions.ShowNote(g); err != nil {
			log.Panicln(err)
		}

		v.Editable = true
		v.Wrap = true
	}

	return nil
}
