package kbFunctions

import (
	"errors"

	"github.com/jameycribbs/cribbnotes_cui/config"
	"github.com/jameycribbs/cribbnotes_cui/db"
	"github.com/jroimartin/gocui"
)

// ShowNote shows the current note in the note view.
func ShowNote(g *gocui.Gui) error {
	var fileID string
	var note *db.Record
	var err error

	if err = clearNoteViews(g); err != nil {
		return err
	}

	if fileID, err = getFileID(g, "toc"); err != nil {
		return err
	}

	if fileID == "" {
		return nil
	}

	if note, err = db.Find(config.DataDir, fileID); err != nil {
		return errors.New("error on db.Find: " + err.Error())
	}

	if err = showNoteInNoteView(g, note); err != nil {
		return err
	}

	if err = g.SetCurrentView("toc"); err != nil {
		return errors.New("(saveNote) error setting current view to toc: " + err.Error())
	}

	return nil
}
