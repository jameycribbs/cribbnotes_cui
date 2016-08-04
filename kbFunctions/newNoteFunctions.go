package kbFunctions

import (
	"errors"
	"strings"
	"time"

	"github.com/jameycribbs/cribbnotes_cui/config"
	"github.com/jameycribbs/cribbnotes_cui/db"
	"github.com/jroimartin/gocui"
)

func abortNewTitle(g *gocui.Gui, v *gocui.View) error {
	var err error

	if err = g.DeleteView("newTitle"); err != nil {
		return errors.New("(abortNewTitle) error deleting newTitle view: " + err.Error())
	}

	if err = g.SetCurrentView("toc"); err != nil {
		return errors.New("(abortNewTitle) error setting current view to toc: " + err.Error())
	}

	return nil
}

func createNote(g *gocui.Gui, v *gocui.View) error {
	var newTitleView *gocui.View
	var err error
	var title string
	var rec db.Record
	var fileID string

	if newTitleView, err = g.View("newTitle"); err != nil {
		return errors.New("(createNote) error grabbing newTitle view: " + err.Error())
	}

	title = strings.TrimSuffix(newTitleView.ViewBuffer(), "\n")

	if err = g.DeleteView("newTitle"); err != nil {
		return errors.New("(createNote) error deleting newTitle view: " + err.Error())
	}

	if title == "" {
		updateStatus(g, "Error adding a new note - Title is blank!")

		if err = g.SetCurrentView("toc"); err != nil {
			return errors.New("(createNote) error setting view to toc: " + err.Error())
		}
		return nil
	}

	rec = db.Record{Title: title, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	if fileID, err = db.Create(config.DataDir, &rec); err != nil {
		return err
	}

	if err = PopulateToc(g, ""); err != nil {
		return err
	}

	if err = scrollToFileID(g, fileID); err != nil {
		return err
	}

	if err = showNoteInNoteView(g, &rec); err != nil {
		return err
	}

	if err = g.SetCurrentView("noteDetail"); err != nil {
		return errors.New("(createNote) error setting view to noteDetail: " + err.Error())
	}

	return nil
}

func newRec(g *gocui.Gui, v *gocui.View) error {
	var err error

	if err = createInputView(g, "newTitle", "Enter title for new note:", true); err != nil {
		return err
	}

	updateStatus(g, "Enter a title for the new note and press [Enter] to save.  Press [Ctrl-X] to abort.")

	return nil
}
