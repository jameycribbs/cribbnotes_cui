package kbFunctions

import (
	"time"

	"github.com/jameycribbs/cribbnotes_cui/config"
	"github.com/jameycribbs/cribbnotes_cui/db"
	"github.com/jroimartin/gocui"
)

func saveNote(g *gocui.Gui, v *gocui.View) error {
	var fileID string
	var err error
	var noteDetailView *gocui.View
	var rec *db.Record

	fileID, err = getFileID(g, "toc")
	if err != nil {
		return err
	}

	rec, err = db.Find(config.DataDir, fileID)
	if err != nil {
		return err
	}

	if noteDetailView, err = g.View("noteDetail"); err != nil {
		return err
	}

	rec.Text = noteDetailView.Buffer()

	rec.UpdatedAt = time.Now()

	if err = db.Update(config.DataDir, rec, fileID); err != nil {
		return err
	}

	updateStatus(g, "Note saved!")

	return nil
}
