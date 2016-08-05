package main

import (
	"errors"
	"os"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/jameycribbs/cribbnotes_cui/config"
	"github.com/jameycribbs/cribbnotes_cui/db"
	"github.com/jameycribbs/cribbnotes_cui/kbFunctions"
	"github.com/jameycribbs/cribbnotes_cui/layouts"
	"github.com/jroimartin/gocui"
)

func main() {
	var err error
	var g *gocui.Gui
	var dataDir *string
	var vimMode *bool

	dataDir = kingpin.Arg("datadir", "Data directory.").Required().String()
	vimMode = kingpin.Flag("vim-mode", "Vim mode.").Bool()

	kingpin.Parse()

	config.DataDir = *dataDir
	config.VimMode = *vimMode

	if err = os.MkdirAll(config.DataDir, os.ModePerm); err != nil {
		panic(errors.New("(main) error creating db directory: " + err.Error()))
	}

	if db.Count(config.DataDir) == 0 {
		if err = createHelpNote(); err != nil {
			panic(errors.New("(main) error creating help note: " + err.Error()))
		}
	}

	g = gocui.NewGui()
	if err = g.Init(); err != nil {
		panic(errors.New("(main) error initiating gui: " + err.Error()))
	}
	defer g.Close()

	g.SetLayout(layouts.Layout)

	if err = kbFunctions.KeybindingsCommon(g); err != nil {
		panic(err)
	}

	if config.VimMode {
		if err = kbFunctions.KeybindingsVim(g); err != nil {
			panic(err)
		}
	} else {
		if err = kbFunctions.KeybindingsNonVim(g); err != nil {
			panic(err)
		}
	}

	g.SelBgColor = gocui.ColorGreen
	g.SelFgColor = gocui.ColorBlack
	g.Cursor = true

	if err = g.MainLoop(); err != nil && err != gocui.ErrQuit {
		panic(err)
	}
}

func createHelpNote() error {
	var err error
	var rec db.Record

	rec = db.Record{}
	rec.Title = "CribbNotes Help"

	rec.Text = "CribbNotes is a simple, console-user-interface note taking application that stores it's data in json files.\n\n"

	rec.CreatedAt = time.Now()
	rec.UpdatedAt = time.Now()

	_, err = db.Create(config.DataDir, &rec)
	if err != nil {
		return err
	}
	return nil
}
