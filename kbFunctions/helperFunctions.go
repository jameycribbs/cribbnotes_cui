package kbFunctions

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jameycribbs/cribbnotes_cui/config"
	"github.com/jameycribbs/cribbnotes_cui/db"
	"github.com/jroimartin/gocui"
)

func clearNoteViews(g *gocui.Gui) error {
	noteTitleView, noteNumberView, noteView, err := getNoteViews(g)
	if err != nil {
		return err
	}

	noteTitleView.Clear()
	noteNumberView.Clear()
	noteView.Clear()

	return nil
}

func createMessageView(g *gocui.Gui, vName string, vTitle string, msg string) error {
	var maxMsgLineLength int

	msgLines := strings.Split(msg, "\n")

	for _, msgLine := range msgLines {
		lineLength := len(msgLine)

		if lineLength > maxMsgLineLength {
			maxMsgLineLength = lineLength
		}
	}

	maxX, maxY := g.Size()

	if v, err := g.SetView(vName, maxX/2-maxMsgLineLength/2-1, maxY/2, maxX/2+maxMsgLineLength/2+1, maxY/2+1+len(msgLines)); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = vTitle

		for _, msgLine := range msgLines {
			fmt.Fprintln(v, msgLine)
		}
	}
	return nil
}

func createInputView(g *gocui.Gui, vName string, vTitle string, editable bool) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView(vName, maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = vTitle
		v.Editable = editable

		if err := g.SetCurrentView(vName); err != nil {
			return err
		}
	}
	return nil
}

func getFileID(g *gocui.Gui, vName string) (string, error) {
	v, err := g.View(vName)
	if err != nil {
		return "", errors.New("error grabbing handle to view in getFileID: " + err.Error())
	}

	_, cy := v.Cursor()

	line, err := v.Line(cy)
	if err != nil {
		if err.Error() == "invalid point" {
			return "", nil
		}
		return "", errors.New("error on Line method for view: " + err.Error())
	}

	lineParts := strings.Split(line, " - ")

	return strings.TrimSpace(lineParts[0]), nil
}

func getNoteViews(g *gocui.Gui) (*gocui.View, *gocui.View, *gocui.View, error) {
	noteTitleView, err := g.View("noteTitle")
	if err != nil {
		return nil, nil, nil, errors.New("(getNoteViews) error grabbing handle to noteTitleView: " + err.Error())
	}

	noteNumberView, err := g.View("noteNumber")
	if err != nil {
		return nil, nil, nil, errors.New("(getNoteViews) error grabbing handle to noteNumberView: " + err.Error())
	}

	noteView, err := g.View("note")
	if err != nil {
		return nil, nil, nil, errors.New("(getNoteViews) error grabbing handle to noteView: " + err.Error())
	}

	return noteTitleView, noteNumberView, noteView, nil
}

// PopulateToc searches for notes and populates the table of contents.
func PopulateToc(g *gocui.Gui, searchStr string) error {
	tocView, err := g.View("toc")
	if err != nil {
		return errors.New("error grabbing handle to toc view: " + err.Error())
	}

	tocView.Clear()

	notes, err := db.Search(config.DataDir, searchStr)
	if err != nil {
		return errors.New("error doing db.Search: " + err.Error())
	}

	for _, note := range notes {
		fileID, err := strconv.Atoi(note.FileID)
		if err != nil {
			return errors.New("error converting fileID to int: " + err.Error())
		}

		fmt.Fprintf(tocView, "%4d - %v\n", fileID, note.Title)
	}

	if err := tocView.SetCursor(0, 0); err != nil {
		return errors.New("error on SetCursor on toc view: " + err.Error())
	}

	if err := tocView.SetOrigin(0, 0); err != nil {
		return errors.New("error on SetOrigin on toc view: " + err.Error())
	}

	return nil
}

func showNoteInNoteView(g *gocui.Gui, rec *db.Record) error {
	var err error

	if err := clearNoteViews(g); err != nil {
		return err
	}

	noteTitleView, noteNumberView, noteView, err := getNoteViews(g)
	if err != nil {
		return err
	}

	fmt.Fprintf(noteTitleView, " \x1b[0;32mNote Title: \x1b[0;37m%s", rec.Title)
	fmt.Fprintf(noteNumberView, " \x1b[0;32mNote #: \x1b[0;37m%s", rec.FileID)
	fmt.Fprintf(noteView, "%s", rec.Text)

	if err := noteView.SetCursor(0, 0); err != nil {
		return errors.New("error setting cursor for noteView: " + err.Error())
	}

	return nil
}

func scrollToFileID(g *gocui.Gui, fileID string) error {
	tocView, err := g.View("toc")
	if err != nil {
		return err
	}

	for {
		cx, cy := tocView.Cursor()
		if err := tocView.SetCursor(cx, cy+1); err != nil {
			ox, oy := tocView.Origin()
			if err := tocView.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}

		nextFileID, err := getFileID(g, "toc")
		if err != nil {
			return err
		}

		if nextFileID == fileID {
			break
		}
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
