package kb_functions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jameycribbs/cribbnotes_cui/db"
	"github.com/jroimartin/gocui"
)

func getFileId(g *gocui.Gui, v *gocui.View) string {
	var line string
	var err error

	_, cy := v.Cursor()
	if line, err = v.Line(cy); err != nil {
		panic(err)
	}

	lineParts := strings.Split(line, " - ")

	return strings.TrimSpace(lineParts[0])
}

func PopulateToc(g *gocui.Gui, searchStr string) error {
	tocView, err := g.View("toc")
	if err != nil {
		return err
	}

	tocView.Clear()

	notes, err := db.Search("data", searchStr)
	if err != nil {
		return err
	}

	for _, note := range notes {
		fileid, err := strconv.Atoi(note.FileId)
		if err != nil {
			return err
		}

		fmt.Fprintf(tocView, "%4d - %v\n", fileid, note.Title)
	}

	if err := tocView.SetCursor(0, 0); err != nil {
		return err
	}

	if err := tocView.SetOrigin(0, 0); err != nil {
		return err
	}

	return nil
}

func showRecInMainView(g *gocui.Gui, v *gocui.View, rec *db.Record) error {
	var err error

	mainView, err := g.View("main")
	if err != nil {
		return err
	}

	mainView.Clear()

	fmt.Fprintf(mainView, "%s", rec.Text)

	mainView.Title = "[ Title: " + rec.Title + "    Note #: " + rec.FileId + " ]"

	if err := mainView.SetCursor(0, 0); err != nil {
		return err
	}

	return nil
}

func updateStatus(g *gocui.Gui, msg string) error {
	statusView, err := g.View("status")
	if err != nil {
		return err
	}

	statusView.Clear()

	if err := statusView.SetOrigin(0, 0); err != nil {
		return err
	}

	if err := statusView.SetCursor(0, 0); err != nil {
		return err
	}

	fmt.Fprint(statusView, msg)

	return nil
}
