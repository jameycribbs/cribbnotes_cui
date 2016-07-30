package main

import (
	"errors"
	"os"

	"github.com/jameycribbs/cribbnotes_cui/config"
	"github.com/jameycribbs/cribbnotes_cui/kb_functions"
	"github.com/jameycribbs/cribbnotes_cui/layouts"
	"github.com/jroimartin/gocui"
)

func main() {
	if len(os.Args) < 2 {
		panic(errors.New("You must supply a command line argument for the location of the data directory!"))
	}

	config.DataDir = os.Args[1]

	g := gocui.NewGui()
	if err := g.Init(); err != nil {
		panic(err)
	}
	defer g.Close()

	g.SetLayout(layouts.Layout)

	if err := kb_functions.Keybindings(g); err != nil {
		panic(err)
	}

	g.SelBgColor = gocui.ColorGreen
	g.SelFgColor = gocui.ColorBlack
	g.Cursor = true

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		panic(err)
	}
}
