package kb_functions

import (
	"strings"
	"time"

	"github.com/jameycribbs/cribbnotes_cui/config"
	"github.com/jameycribbs/cribbnotes_cui/db"
	"github.com/jroimartin/gocui"
)

func AbortNewTitle(g *gocui.Gui, v *gocui.View) error {
	if err := g.DeleteView("newTitle"); err != nil {
		return err
	}

	if err := g.SetCurrentView("toc"); err != nil {
		return err
	}

	return nil
}

func createNote(g *gocui.Gui, v *gocui.View) error {
	newTitleView, err := g.View("newTitle")
	if err != nil {
		return err
	}

	title := strings.TrimSuffix(newTitleView.ViewBuffer(), "\n")

	if err := g.DeleteView("newTitle"); err != nil {
		return err
	}

	if title == "" {
		updateStatus(g, "Error adding a new note - Title is blank!")

		if err := g.SetCurrentView("toc"); err != nil {
			return err
		}
		return nil
	}

	rec := db.Record{Title: title, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	fileId, err := db.Create(config.DataDir, &rec)
	if err != nil {
		return err
	}

	if err := PopulateToc(g, ""); err != nil {
		return err
	}

	if err := scrollToFileId(g, fileId); err != nil {
		return err
	}

	showNoteInNoteView(g, &rec)
	if err != nil {
		return err
	}

	if err := g.SetCurrentView("note"); err != nil {
		return err
	}

	updateStatus(g, "Ctrl+s - save note")

	return nil
}

func newRec(g *gocui.Gui, v *gocui.View) error {
	if err := createInputView(g, "newTitle", "Enter title for new note:"); err != nil {
		return err
	}

	updateStatus(g, "Enter a title for the new note.  Press [Enter] when done.")

	return nil
}
