package main

import (
	"fmt"
	"log"

	"github.com/jameycribbs/cribbnotes_cui/kb_functions"
	"github.com/jroimartin/gocui"
)

var currentVersion = "0.10"

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("headline", int(0.4*float32(maxX)), -1, int(0.4*float32(maxX))+19, 1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintf(v, " CribbNotes v%v", currentVersion)
		v.Frame = true
	}

	if v, err := g.SetView("toc", -1, 2, int(0.8*float32(maxX)), maxY-25); err != nil {
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

	if v, err := g.SetView("help", int(0.8*float32(maxX)), 2, maxX, maxY-25); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "[ Help ]"

		fmt.Fprintln(v, "[k/j]      - scroll")
		fmt.Fprintln(v, "[Enter]    - view note")
		fmt.Fprintln(v, "[Ctrl+k/j] - switch views")
		fmt.Fprintln(v, "[Ctrl+f]   - find notes")
		fmt.Fprintln(v, "[Ctrl+n]   - new note")
		fmt.Fprintln(v, "[Ctrl+s]   - save note")
		fmt.Fprintln(v, "[Ctrl+d]   - delete note")
		fmt.Fprintln(v, "[Ctrl+q]   - quit")
	}

	if v, err := g.SetView("main", -1, maxY-23, maxX, maxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		tocView, err := g.View("toc")
		if err != nil {
			return err
		}

		if err := kb_functions.ShowRec(g, tocView); err != nil {
			log.Panicln(err)
		}

		v.Editable = true
		v.Wrap = true
	}

	if _, err := g.SetView("status", -1, maxY-2, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	return nil
}

func main() {
	g := gocui.NewGui()
	if err := g.Init(); err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetLayout(layout)
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
