package kbFunctions

import "github.com/jroimartin/gocui"

func nextView(g *gocui.Gui, v *gocui.View) error {
	if v == nil || v.Name() == "toc" {
		return g.SetCurrentView("note")
	}

	updateStatus(g, "")

	return g.SetCurrentView("toc")
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
