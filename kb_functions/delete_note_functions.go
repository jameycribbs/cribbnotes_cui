package kb_functions

import (
	"errors"

	"github.com/jameycribbs/cribbnotes_cui/config"
	"github.com/jameycribbs/cribbnotes_cui/db"
	"github.com/jroimartin/gocui"
)

func deleteRec(g *gocui.Gui, v *gocui.View) error {
	fileId, err := getFileId(g, "toc")
	if err != nil {
		return err
	}

	if err := db.Delete(config.DataDir, fileId); err != nil {
		if err := g.SetCurrentView("toc"); err != nil {
			return errors.New("(deleteRec) error setting current view to toc " + err.Error())
		}
		return err
	}

	if err := PopulateToc(g, ""); err != nil {
		return err
	}

	if err := ShowNote(g); err != nil {
		return err
	}

	updateStatus(g, "Note # "+fileId+" deleted.")

	return nil
}
