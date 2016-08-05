package layouts

import (
	"errors"
	"fmt"

	"github.com/jroimartin/gocui"
)

var currentVersion = "1.1"

func headlineLayout(g *gocui.Gui) error {
	var err error
	var maxX int
	var v *gocui.View

	maxX, _ = g.Size()

	if v, err = g.SetView("headline", int(0.4*float32(maxX)), -1, int(0.4*float32(maxX))+19, 1); err != nil {
		if err != gocui.ErrUnknownView {
			return errors.New("(headlineLayout) error setting view: " + err.Error())
		}
		fmt.Fprintf(v, " CribbNotes v%v", currentVersion)
	}

	return nil
}
