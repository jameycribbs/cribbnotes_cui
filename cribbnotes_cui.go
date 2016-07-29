package main

import (
	"log"

	"github.com/jameycribbs/cribbnotes_cui/kb_functions"
	"github.com/jameycribbs/cribbnotes_cui/layouts"
	"github.com/jroimartin/gocui"
)

func main() {
	g := gocui.NewGui()
	if err := g.Init(); err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetLayout(layouts.Layout)

	if err := kb_functions.Keybindings(g); err != nil {
		log.Panicln(err)
	}

	g.SelBgColor = gocui.ColorGreen
	g.SelFgColor = gocui.ColorBlack
	g.Cursor = true

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
