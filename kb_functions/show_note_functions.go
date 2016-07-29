package kb_functions

import (
	"github.com/jameycribbs/cribbnotes_cui/db"
	"github.com/jroimartin/gocui"
)

func ShowNote(g *gocui.Gui) error {
	var err error

	fileId, err := getFileId(g, "toc")
	if err != nil {
		return err
	}

	note, err := db.Find("data", fileId)
	if err != nil {
		return err
	}

	if err := showNoteInNoteView(g, note); err != nil {
		return err
	}

	if err := g.SetCurrentView("toc"); err != nil {
		return err
	}

	return nil
}
