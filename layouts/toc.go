package layouts

import (
	"github.com/jameycribbs/cribbnotes_cui/kb_functions"
	"github.com/jroimartin/gocui"
)

func tocLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("toc", -1, 2, maxX-32, maxY-25); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Highlight = true
		v.Wrap = true
		v.Title = "[ Table of Contents ]"

		if err := kb_functions.PopulateToc(g, ""); err != nil {
			return err
		}

		if err := g.SetCurrentView("toc"); err != nil {
			return err
		}
	}

	return nil
}
