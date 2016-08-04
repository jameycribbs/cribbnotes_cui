package layouts

import (
	"errors"

	"github.com/jameycribbs/cribbnotes_cui/kbFunctions"
	"github.com/jroimartin/gocui"
)

func tocLayout(g *gocui.Gui) error {
	var maxX, maxY int
	var err error
	var v *gocui.View

	maxX, maxY = g.Size()

	if v, err = g.SetView("toc", -1, 2, maxX-34, maxY-25); err != nil {
		if err != gocui.ErrUnknownView {
			return errors.New("(tocLayout) error setting view: " + err.Error())
		}

		v.Highlight = true
		v.Wrap = true
		v.Title = "[ Table of Contents ]"

		if err = kbFunctions.PopulateToc(g, ""); err != nil {
			return err
		}

		if err = g.SetCurrentView("toc"); err != nil {
			return errors.New("(tocLayout) error setting current view to toc: " + err.Error())
		}
	}

	return nil
}
