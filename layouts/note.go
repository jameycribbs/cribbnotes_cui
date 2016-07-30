package layouts

import (
	"github.com/jameycribbs/cribbnotes_cui/kb_functions"
	"github.com/jroimartin/gocui"
)

func noteTitleLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("noteTitle", -1, maxY-25, maxX-17, maxY-23); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = false
	}

	return nil
}

func noteNumberLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("noteNumber", maxX-17, maxY-25, maxX, maxY-23); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = false
	}

	return nil
}

func noteLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("note", 0, maxY-22, maxX-1, maxY-3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		if err := kb_functions.ShowNote(g); err != nil {
			return err
		}

		v.Editable = true
		v.Wrap = true
		v.Frame = false
	}

	return nil
}
