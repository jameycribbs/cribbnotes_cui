package kb_functions

import (
	"errors"

	"github.com/jameycribbs/cribbnotes_cui/config"
	"github.com/jameycribbs/cribbnotes_cui/db"
	"github.com/jroimartin/gocui"
)

func ShowNote(g *gocui.Gui) error {
	var err error

	if err := clearNoteViews(g); err != nil {
		return err
	}

	fileId, err := getFileId(g, "toc")
	if err != nil {
		return err
	}

	if fileId == "" {
		return nil
	}

	note, err := db.Find(config.DataDir, fileId)
	if err != nil {
		return errors.New("error on db.Find: " + err.Error())
	}

	if err := showNoteInNoteView(g, note); err != nil {
		return err
	}

	if err := g.SetCurrentView("toc"); err != nil {
		return errors.New("error setting current view to toc: " + err.Error())
	}

	return nil
}
