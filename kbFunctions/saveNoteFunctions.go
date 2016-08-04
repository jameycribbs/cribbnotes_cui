package kbFunctions

import (
	"errors"
	"time"

	"github.com/jameycribbs/cribbnotes_cui/config"
	"github.com/jameycribbs/cribbnotes_cui/db"
	"github.com/jroimartin/gocui"
)

func saveNote(g *gocui.Gui, v *gocui.View) error {
	var fileID string
	var err error
	var rec *db.Record

	fileID, err = getFileID(g, "toc")
	if err != nil {
		return err
	}

	rec, err = db.Find(config.DataDir, fileID)
	if err != nil {
		return err
	}

	rec.Text = v.Buffer()

	rec.UpdatedAt = time.Now()

	if err = db.Update(config.DataDir, rec, fileID); err != nil {
		return err
	}

	updateStatus(g, "Note saved!")

	if err = g.SetCurrentView("toc"); err != nil {
		return errors.New("(saveNote) error setting current view to toc: " + err.Error())
	}
	return nil
}
