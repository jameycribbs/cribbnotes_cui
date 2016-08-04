package layouts

import (
	"errors"

	"github.com/jameycribbs/cribbnotes_cui/config"
	"github.com/jameycribbs/cribbnotes_cui/kbFunctions"
	"github.com/jroimartin/gocui"
)

func noteTitleLayout(g *gocui.Gui) error {
	var maxX, maxY int
	var v *gocui.View
	var err error

	maxX, maxY = g.Size()

	if v, err = g.SetView("noteTitle", -1, maxY-25, maxX-17, maxY-23); err != nil {
		if err != gocui.ErrUnknownView {
			return errors.New("(noteTitleLayout) error setting view: " + err.Error())
		}
		v.Frame = false
	}

	return nil
}

func noteNumberLayout(g *gocui.Gui) error {
	var maxX, maxY int
	var v *gocui.View
	var err error

	maxX, maxY = g.Size()

	if v, err = g.SetView("noteNumber", maxX-17, maxY-25, maxX, maxY-23); err != nil {
		if err != gocui.ErrUnknownView {
			return errors.New("(noteNumberLayout) error setting view: " + err.Error())
		}
		v.Frame = false
	}

	return nil
}

func noteDetailLayout(g *gocui.Gui) error {
	var maxX, maxY int
	var v *gocui.View
	var err error

	maxX, maxY = g.Size()

	if v, err = g.SetView("noteDetail", 0, maxY-22, maxX-1, maxY-3); err != nil {
		if err != gocui.ErrUnknownView {
			return errors.New("(noteDetailLayout) error setting view: " + err.Error())
		}

		if err = kbFunctions.ShowNote(g); err != nil {
			return err
		}

		v.Editable = !config.VimMode
		v.Wrap = true
		v.Frame = false
	}

	return nil
}
