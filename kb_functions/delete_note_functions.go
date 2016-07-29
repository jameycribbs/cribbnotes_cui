package kb_functions

import (
	"strings"

	"github.com/jameycribbs/cribbnotes_cui/db"
	"github.com/jroimartin/gocui"
)

func AbortGetFileIdToDelete(g *gocui.Gui, v *gocui.View) error {
	if err := g.DeleteView("fileIdToDelete"); err != nil {
		return err
	}

	if err := g.SetCurrentView("toc"); err != nil {
		return err
	}
	return nil
}

func deleteRec(g *gocui.Gui, v *gocui.View) error {
	fileIdToDeleteView, err := g.View("fileIdToDelete")
	if err != nil {
		return err
	}

	fileId := strings.TrimSuffix(fileIdToDeleteView.ViewBuffer(), "\n")

	if err := g.DeleteView("fileIdToDelete"); err != nil {
		return err
	}

	if err := db.Delete("data", fileId); err != nil {
		if err := g.SetCurrentView("toc"); err != nil {
			return err
		}

		updateStatus(g, "Error trying to delete note # "+fileId+" - "+err.Error())

		return nil
	}

	if err := PopulateToc(g, ""); err != nil {
		return err
	}

	if err := g.SetCurrentView("toc"); err != nil {
		return err
	}

	updateStatus(g, "Note # "+fileId+" deleted.")

	return nil
}

func getFileIdToDelete(g *gocui.Gui, v *gocui.View) error {
	if err := createInputView(g, "fileIdToDelete", "Enter note number to delete:"); err != nil {
		return err
	}

	updateStatus(g, "Enter a note number to delete.  Press [Enter] to delete the note.")

	return nil
}
