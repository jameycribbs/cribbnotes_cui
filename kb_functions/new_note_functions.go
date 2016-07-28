package kb_functions

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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

func createRec(g *gocui.Gui, v *gocui.View) error {
	newTitleView, err := g.View("newTitle")
	if err != nil {
		panic(err)
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

	fileId, err := db.Create("data", &rec)
	if err != nil {
		return err
	}

	tocView, err := g.View("toc")
	if err != nil {
		return err
	}

	intFileId, err := strconv.Atoi(fileId)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(tocView, "%4d - %v\n", intFileId, rec.Title)

	for {
		cx, cy := tocView.Cursor()
		if err := tocView.SetCursor(cx, cy+1); err != nil {
			ox, oy := tocView.Origin()
			if err := tocView.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}

		nextFileId := getFileId(g, tocView)

		if nextFileId == fileId {
			break
		}
	}

	showRecInMainView(g, v, &rec)
	if err != nil {
		return err
	}

	if err := g.SetCurrentView("main"); err != nil {
		return err
	}

	updateStatus(g, "Ctrl+s - save note")

	return nil
}

func newRec(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := g.Size()

	if newTitleView, err := g.SetView("newTitle", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		newTitleView.Title = "Title"
		newTitleView.Editable = true

		updateStatus(g, "Enter a title for the new note.  Press [Enter] when done.")

		if err := g.SetCurrentView("newTitle"); err != nil {
			return err
		}
	}
	return nil
}
