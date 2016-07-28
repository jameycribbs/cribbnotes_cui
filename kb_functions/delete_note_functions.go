package kb_functions

import (
	"fmt"
	"strconv"
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
		panic(err)
	}

	fileId := strings.TrimSuffix(fileIdToDeleteView.ViewBuffer(), "\n")

	if err := g.DeleteView("fileIdToDelete"); err != nil {
		return err
	}

	err = db.Delete("data", fileId)
	if err != nil {
		updateStatus(g, "Error trying to delete note # "+fileId+" - "+err.Error())

		if err := g.SetCurrentView("toc"); err != nil {
			return err
		}

		return nil
	}

	tocView, err := g.View("toc")
	if err != nil {
		return err
	}

	tocView.Clear()

	if err := tocView.SetCursor(0, 0); err != nil {
		return err
	}

	if err := tocView.SetOrigin(0, 0); err != nil {
		return err
	}

	recs, err := db.Search("data", "")
	if err != nil {
		return err
	}

	for _, rec := range recs {
		fileid, err := strconv.Atoi(rec.FileId)
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(tocView, "%4d - %v\n", fileid, rec.Title)
	}

	if err := g.SetCurrentView("toc"); err != nil {
		return err
	}

	updateStatus(g, "Note # "+fileId+" deleted.")

	return nil
}

func getFileIdToDelete(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := g.Size()

	if fileIdToDeleteView, err := g.SetView("fileIdToDelete", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		fileIdToDeleteView.Title = "Enter note number to delete:"
		fileIdToDeleteView.Editable = true

		updateStatus(g, "Enter a note number to delete.  Press [Enter] to delete the note.")

		if err := g.SetCurrentView("fileIdToDelete"); err != nil {
			return err
		}
	}
	return nil
}
