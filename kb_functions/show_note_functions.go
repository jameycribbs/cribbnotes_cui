package kb_functions

import (
	"github.com/jameycribbs/cribbnotes_cui/db"
	"github.com/jroimartin/gocui"
)

func ShowRec(g *gocui.Gui, v *gocui.View) error {
	var err error

	fileId := getFileId(g, v)

	rec, err := db.Find("data", fileId)
	if err != nil {
		return err
	}

	showRecInMainView(g, v, rec)
	if err != nil {
		return err
	}

	if err := g.SetCurrentView("toc"); err != nil {
		return err
	}

	return nil
}
