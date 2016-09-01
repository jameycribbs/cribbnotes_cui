package kbFunctions

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jameycribbs/cribbnotes_cui/config"
	"github.com/jameycribbs/cribbnotes_cui/db"
	"github.com/jroimartin/gocui"
)

func abortEditNoteTitle(g *gocui.Gui, v *gocui.View) error {
	var err error

	updateStatus(g, "")

	if err = g.DeleteView("editNoteTitle"); err != nil {
		return errors.New("(abortEditNoteTitle) error deleting editNoteTitle view " + err.Error())
	}

	if err = g.SetCurrentView("toc"); err != nil {
		return errors.New("(abortEditNoteTitle) error setting current view to toc " + err.Error())
	}

	return nil
}

func saveNoteTitle(g *gocui.Gui, v *gocui.View) error {
	var err error
	var fileID string
	var note *db.Record

	if fileID, err = getFileID(g, "toc"); err != nil {
		return err
	}

	if note, err = db.Find(config.DataDir, fileID); err != nil {
		return err
	}

	note.Title = strings.TrimSuffix(v.ViewBuffer(), "\n")
	note.UpdatedAt = time.Now()

	if err = db.Update(config.DataDir, note, fileID); err != nil {
		return err
	}

	updateStatus(g, "Note title saved!")

	if err = PopulateToc(g, ""); err != nil {
		return err
	}

	if err = ShowNote(g); err != nil {
		return err
	}

	if err = g.SetCurrentView("toc"); err != nil {
		return errors.New("(saveNoteTitle) error setting current view to toc " + err.Error())
	}

	if err = g.DeleteView("editNoteTitle"); err != nil {
		return errors.New("(saveNoteTitle) error deleting editNoteTitle view " + err.Error())
	}

	return nil
}

func showEditNoteTitle(g *gocui.Gui, v *gocui.View) error {
	var editNoteTitleView *gocui.View
	var err error
	var fileID string
	var note *db.Record

	if err = createInputView(g, "editNoteTitle", "Note Title", !config.VimMode); err != nil {
		return err
	}

	if fileID, err = getFileID(g, "toc"); err != nil {
		return err
	}

	if note, err = db.Find(config.DataDir, fileID); err != nil {
		return err
	}

	if editNoteTitleView, err = g.View("editNoteTitle"); err != nil {
		return errors.New("(showEditNoteTitle) error getting editNoteTitle view " + err.Error())
	}

	fmt.Fprint(editNoteTitleView, note.Title)

	updateStatus(g, "Update note title and press [Enter] to save.  Press [Esc] to abort.")

	return nil
}
