package layouts

import "github.com/jroimartin/gocui"

func statusLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if _, err := g.SetView("status", -1, maxY-2, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	return nil
}
