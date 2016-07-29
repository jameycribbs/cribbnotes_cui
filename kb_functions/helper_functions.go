package kb_functions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jameycribbs/cribbnotes_cui/db"
	"github.com/jroimartin/gocui"
)

func createInputView(g *gocui.Gui, vName string, vTitle string) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView(vName, maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = vTitle
		v.Editable = true

		if err := g.SetCurrentView(vName); err != nil {
			return err
		}
	}
	return nil
}

func getFileId(g *gocui.Gui, vName string) (string, error) {
	v, err := g.View(vName)
	if err != nil {
		return "", err
	}

	_, cy := v.Cursor()

	line, err := v.Line(cy)
	if err != nil {
		return "", err
	}

	lineParts := strings.Split(line, " - ")

	return strings.TrimSpace(lineParts[0]), nil
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

func showNoteInNoteView(g *gocui.Gui, rec *db.Record) error {
	var err error

	noteView, err := g.View("note")
	if err != nil {
		return err
	}

	noteView.Clear()

	fmt.Fprintf(noteView, "%s", rec.Text)

	noteView.Title = "[ Title: " + rec.Title + "    Note #: " + rec.FileId + " ]"

	if err := noteView.SetCursor(0, 0); err != nil {
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
