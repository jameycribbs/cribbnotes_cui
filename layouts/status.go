package layouts

import (
	"errors"

	"github.com/jroimartin/gocui"
)

func statusLayout(g *gocui.Gui) error {
	var maxX, maxY int
	var err error

	maxX, maxY = g.Size()

	if _, err = g.SetView("status", -1, maxY-2, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return errors.New("(statusLayout) error setting view: " + err.Error())
		}
	}

	return nil
}
