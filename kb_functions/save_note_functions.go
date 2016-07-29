package kb_functions

import (
	"time"

	"github.com/jameycribbs/cribbnotes_cui/db"
	"github.com/jroimartin/gocui"
)

func saveNote(g *gocui.Gui, v *gocui.View) error {
	fileId, err := getFileId(g, "toc")
	if err != nil {
		return err
	}

	rec, err := db.Find("data", fileId)
	if err != nil {
		return err
	}

	rec.Text = v.Buffer()

	rec.UpdatedAt = time.Now()

	if err := db.Update("data", rec, fileId); err != nil {
		return err
	}

	updateStatus(g, "Note saved!")

	return nil
}
