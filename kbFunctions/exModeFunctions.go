package kbFunctions

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
)

func abortExMode(g *gocui.Gui, v *gocui.View) error {
	var err error

	updateStatus(g, "")

	if err = g.SetCurrentView("toc"); err != nil {
		return errors.New("(abortExMode) error setting current view to toc " + err.Error())
	}

	return nil
}

func exMode(g *gocui.Gui, v *gocui.View) error {
	var err error
	var statusView *gocui.View

	if v.Editable {
		return nil
	}

	if statusView, err = g.View("status"); err != nil {
		return err
	}

	statusView.Clear()

	fmt.Fprint(statusView, ":")

	if err = statusView.SetOrigin(0, 0); err != nil {
		return errors.New("(exMode) error setting origin: " + err.Error())
	}

	if err = statusView.SetCursor(1, 0); err != nil {
		return errors.New("(exMode) error setting cursor: " + err.Error())
	}

	if err = g.SetCurrentView("status"); err != nil {
		return errors.New("(noteExMode) error setting current view to status: " + err.Error())
	}

	statusView.Editable = true

	return nil
}

func exExecuteCommand(g *gocui.Gui, v *gocui.View) error {
	var err error
	var exCommand string

	exCommand = strings.TrimSuffix(v.ViewBuffer(), "\n")

	v.Clear()
	v.Editable = false

	if err = g.SetCurrentView("toc"); err != nil {
		return errors.New("(exExecuteCommand) error setting current view to toc: " + err.Error())
	}

	switch exCommand {
	case ":w":
		return saveNote(g, v)
	case ":q":
		return quit(g, v)
	}

	return nil
}
