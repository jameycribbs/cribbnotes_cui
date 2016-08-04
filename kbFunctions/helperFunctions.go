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
	var noteTitleView, noteNumberView, noteDetailView *gocui.View
	var err error

	if noteTitleView, noteNumberView, noteDetailView, err = getNoteViews(g); err != nil {
		return err
	}

	noteTitleView.Clear()
	noteNumberView.Clear()
	noteDetailView.Clear()

	return nil
}

func createMessageView(g *gocui.Gui, vName string, vTitle string, msg string) error {
	var msgLine string
	var msgLines []string
	var lineLength, maxMsgLineLength, maxX, maxY int
	var err error
	var v *gocui.View

	msgLines = strings.Split(msg, "\n")

	for _, msgLine = range msgLines {
		lineLength = len(msgLine)

		if lineLength > maxMsgLineLength {
			maxMsgLineLength = lineLength
		}
	}

	maxX, maxY = g.Size()

	if v, err = g.SetView(vName, maxX/2-maxMsgLineLength/2-1, maxY/2, maxX/2+maxMsgLineLength/2+1, maxY/2+1+len(msgLines)); err != nil {
		if err != gocui.ErrUnknownView {
			return errors.New("(createMessageView) error creating view: " + err.Error())
		}

		v.Title = vTitle

		for _, msgLine = range msgLines {
			fmt.Fprintln(v, msgLine)
		}
	}
	return nil
}

func createInputView(g *gocui.Gui, vName string, vTitle string, editable bool) error {
	var maxX, maxY int
	var v *gocui.View
	var err error

	maxX, maxY = g.Size()

	if v, err = g.SetView(vName, maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return errors.New("(createInputView) error creating view: " + err.Error())
		}

		v.Title = vTitle
		v.Editable = editable

		if err = g.SetCurrentView(vName); err != nil {
			return errors.New("(createInputView) error setting current view: " + err.Error())
		}
	}
	return nil
}

func getFileID(g *gocui.Gui, vName string) (string, error) {
	var v *gocui.View
	var err error
	var cy int
	var lineParts []string
	var line string

	if v, err = g.View(vName); err != nil {
		return "", errors.New("(getFileID) error grabbing handle to view in getFileID: " + err.Error())
	}

	_, cy = v.Cursor()

	if line, err = v.Line(cy); err != nil {
		if err.Error() == "invalid point" {
			return "", nil
		}
		return "", errors.New("(getFileID) error on Line method for view: " + err.Error())
	}

	lineParts = strings.Split(line, " - ")

	return strings.TrimSpace(lineParts[0]), nil
}

func getNoteViews(g *gocui.Gui) (*gocui.View, *gocui.View, *gocui.View, error) {
	var noteTitleView, noteNumberView, noteDetailView *gocui.View
	var err error

	if noteTitleView, err = g.View("noteTitle"); err != nil {
		return nil, nil, nil, errors.New("(getNoteViews) error grabbing handle to noteTitleView: " + err.Error())
	}

	if noteNumberView, err = g.View("noteNumber"); err != nil {
		return nil, nil, nil, errors.New("(getNoteViews) error grabbing handle to noteNumberView: " + err.Error())
	}

	if noteDetailView, err = g.View("noteDetail"); err != nil {
		return nil, nil, nil, errors.New("(getNoteViews) error grabbing handle to noteView: " + err.Error())
	}

	return noteTitleView, noteNumberView, noteDetailView, nil
}

// PopulateToc searches for notes and populates the table of contents.
func PopulateToc(g *gocui.Gui, searchStr string) error {
	var tocView *gocui.View
	var err error
	var notes []db.Record
	var note db.Record
	var fileID int

	if tocView, err = g.View("toc"); err != nil {
		return errors.New("(PopulateToc) error grabbing handle to toc view: " + err.Error())
	}

	tocView.Clear()

	if notes, err = db.Search(config.DataDir, searchStr); err != nil {
		return errors.New("(PopulateToc) error doing db.Search: " + err.Error())
	}

	for _, note = range notes {
		if fileID, err = strconv.Atoi(note.FileID); err != nil {
			return errors.New("(PopulateToc) error converting fileID to int: " + err.Error())
		}

		fmt.Fprintf(tocView, "%4d - %v\n", fileID, note.Title)
	}

	if err = tocView.SetCursor(0, 0); err != nil {
		return errors.New("(PopulateToc) error on SetCursor on toc view: " + err.Error())
	}

	if err = tocView.SetOrigin(0, 0); err != nil {
		return errors.New("(PopulateToc) error on SetOrigin on toc view: " + err.Error())
	}

	return nil
}

func showNoteInNoteView(g *gocui.Gui, rec *db.Record) error {
	var noteTitleView, noteNumberView, noteDetailView *gocui.View
	var err error

	if err = clearNoteViews(g); err != nil {
		return err
	}

	if noteTitleView, noteNumberView, noteDetailView, err = getNoteViews(g); err != nil {
		return err
	}

	fmt.Fprintf(noteTitleView, " \x1b[0;32mNote Title: \x1b[0;37m%s", rec.Title)
	fmt.Fprintf(noteNumberView, " \x1b[0;32mNote #: \x1b[0;37m%s", rec.FileID)
	fmt.Fprintf(noteDetailView, "%s", rec.Text)

	if err = noteDetailView.SetCursor(0, 0); err != nil {
		return errors.New("(showNoteInNoteView) error setting cursor for noteView: " + err.Error())
	}

	return nil
}

func scrollToFileID(g *gocui.Gui, fileID string) error {
	var tocView *gocui.View
	var err error
	var cx, cy, ox, oy int
	var nextFileID string

	if tocView, err = g.View("toc"); err != nil {
		return errors.New("(scrollToFileID) error getting to view: " + err.Error())
	}

	for {
		cx, cy = tocView.Cursor()
		if err = tocView.SetCursor(cx, cy+1); err != nil {
			ox, oy = tocView.Origin()
			if err = tocView.SetOrigin(ox, oy+1); err != nil {
				return errors.New("(scrollToFileID) error setting origin: " + err.Error())
			}
		}

		if nextFileID, err = getFileID(g, "toc"); err != nil {
			return err
		}

		if nextFileID == fileID {
			break
		}
	}
	return nil
}

func updateStatus(g *gocui.Gui, msg string) error {
	var statusView *gocui.View
	var err error

	if statusView, err = g.View("status"); err != nil {
		return err
	}

	statusView.Clear()

	if err = statusView.SetOrigin(0, 0); err != nil {
		return errors.New("(updateStatus) error setting origin: " + err.Error())
	}

	if err = statusView.SetCursor(0, 0); err != nil {
		return errors.New("(updateStatus) error setting cursor: " + err.Error())
	}

	fmt.Fprint(statusView, msg)

	return nil
}
