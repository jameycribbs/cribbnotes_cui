package layouts

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

var currentVersion = "0.10"

func headlineLayout(g *gocui.Gui) error {
	maxX, _ := g.Size()

	if v, err := g.SetView("headline", int(0.4*float32(maxX)), -1, int(0.4*float32(maxX))+19, 1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintf(v, " CribbNotes v%v", currentVersion)
		v.Frame = true
	}

	return nil
}