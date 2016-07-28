package kb_functions

import (
	"time"

	"github.com/jameycribbs/cribbnotes_cui/db"
	"github.com/jroimartin/gocui"
)

func saveRec(g *gocui.Gui, v *gocui.View) error {
	tocView, err := g.View("toc")
	if err != nil {
		panic(err)
	}

	fileId := getFileId(g, tocView)

	rec, err := db.Find("data", fileId)
	if err != nil {
		return err
	}

	rec.Text = v.ViewBuffer()

	rec.UpdatedAt = time.Now()

	err = db.Update("data", rec, fileId)
	if err != nil {
		return err
	}

	updateStatus(g, "Note saved!")

	return nil
}
