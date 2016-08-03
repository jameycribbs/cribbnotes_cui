package kbFunctions

import (
	"fmt"
	"strings"
	"time"

	"github.com/jameycribbs/cribbnotes_cui/config"
	"github.com/jameycribbs/cribbnotes_cui/db"
	"github.com/jroimartin/gocui"
)

// AbortEditNoteTitle deletes the edit note title dialog.
func AbortEditNoteTitle(g *gocui.Gui, v *gocui.View) error {
	if err := g.DeleteView("editNoteTitle"); err != nil {
		return err
	}

	if err := g.SetCurrentView("toc"); err != nil {
		return err
	}

	return nil
}

func saveNoteTitle(g *gocui.Gui, v *gocui.View) error {
	fileID, err := getFileID(g, "toc")
	if err != nil {
		return err
	}

	rec, err := db.Find(config.DataDir, fileID)
	if err != nil {
		return err
	}

	title := strings.TrimSuffix(v.ViewBuffer(), "\n")

	rec.Title = title

	rec.UpdatedAt = time.Now()

	if err := db.Update(config.DataDir, rec, fileID); err != nil {
		return err
	}

	updateStatus(g, "Note title saved!")

	if err := PopulateToc(g, ""); err != nil {
		return err
	}

	if err := ShowNote(g); err != nil {
		return err
	}

	if err := g.SetCurrentView("toc"); err != nil {
		return err
	}

	if err := g.DeleteView("editNoteTitle"); err != nil {
		return err
	}

	return nil
}

func showEditNoteTitle(g *gocui.Gui, v *gocui.View) error {
	if err := createInputView(g, "editNoteTitle", "Note Title", !config.VimMode); err != nil {
		return err
	}

	fileID, err := getFileID(g, "toc")
	if err != nil {
		return err
	}

	rec, err := db.Find(config.DataDir, fileID)
	if err != nil {
		return err
	}

	editNoteTitleView, err := g.View("editNoteTitle")
	if err != nil {
		return err
	}

	fmt.Fprint(editNoteTitleView, rec.Title)

	updateStatus(g, "Update note title and press [Enter] to save.  Press [Ctrl-C] to abort.")

	return nil
}
