package kbFunctions

import (
	"errors"

	"github.com/jameycribbs/cribbnotes_cui/config"
	"github.com/jameycribbs/cribbnotes_cui/db"
	"github.com/jroimartin/gocui"
)

func abortDeleteNote(g *gocui.Gui, v *gocui.View) error {
	var err error

	if err = g.DeleteView("deleteNoteConfirm"); err != nil {
		return errors.New("(abortDeleteNote) error deleting deleteNoteConfirm view " + err.Error())
	}

	updateStatus(g, "")

	if err = g.SetCurrentView("toc"); err != nil {
		return errors.New("(abortDeleteNote) error setting current view to toc " + err.Error())
	}

	return nil
}

func deleteNote(g *gocui.Gui, v *gocui.View) error {
	var err error
	var fileID string

	if err = g.DeleteView("deleteNoteConfirm"); err != nil {
		return errors.New("(deleteNote) error deleting deleteNoteConfirm view " + err.Error())
	}

	if fileID, err = getFileID(g, "toc"); err != nil {
		return err
	}

	if err = db.Delete(config.DataDir, fileID); err != nil {
		if err = g.SetCurrentView("toc"); err != nil {
			return errors.New("(deleteNote) error setting current view to toc " + err.Error())
		}
		return err
	}

	if err = PopulateToc(g, ""); err != nil {
		return err
	}

	if err = ShowNote(g); err != nil {
		return err
	}

	updateStatus(g, "Note # "+fileID+" deleted.")

	return nil
}

func showDeleteNoteConfirm(g *gocui.Gui, v *gocui.View) error {
	var err error
	var fileID string

	if fileID, err = getFileID(g, "toc"); err != nil {
		return err
	}

	if err = createMessageView(g, "deleteNoteConfirm", "Delete Confirmation", "Press 'y' to confirm deletion of note # "+fileID+".\nPress 'n' to abort deletion."); err != nil {
		return err
	}

	if err = g.SetCurrentView("deleteNoteConfirm"); err != nil {
		return errors.New("(showDeleteNoteConfirm) error setting current view to deleteNoteConfirm: " + err.Error())
	}

	return nil
}
