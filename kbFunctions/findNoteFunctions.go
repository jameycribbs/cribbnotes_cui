package kbFunctions

import (
	"errors"
	"strings"

	"github.com/jroimartin/gocui"
)

func abortSearch(g *gocui.Gui, v *gocui.View) error {
	var err error

	updateStatus(g, "")

	if err = g.DeleteView("search"); err != nil {
		return errors.New("(AbortSearch) error deleting search view " + err.Error())
	}

	if err = g.SetCurrentView("toc"); err != nil {
		return errors.New("(AbortSearch) error setting current view to toc " + err.Error())
	}
	return nil
}

func clearCurrentFilter(g *gocui.Gui, v *gocui.View) error {
	var err error

	if err = PopulateToc(g, ""); err != nil {
		return err
	}

	if err = g.SetCurrentView("toc"); err != nil {
		return errors.New("(findNotes) error setting current view to toc: " + err.Error())
	}

	if err = ShowNote(g); err != nil {
		return err
	}

	if err = UpdateFilterMsg(g, ""); err != nil {
		return err
	}
	return nil
}

func findNotes(g *gocui.Gui, v *gocui.View) error {
	var err error
	var searchStr string
	var searchView *gocui.View

	if searchView, err = g.View("search"); err != nil {
		return errors.New("(findNotes) error grabbing handle to search view: " + err.Error())
	}

	searchStr = strings.TrimSuffix(searchView.ViewBuffer(), "\n")

	if err = g.DeleteView("search"); err != nil {
		return errors.New("(findNotes) error deleting search view " + err.Error())
	}

	if err = PopulateToc(g, searchStr); err != nil {
		return err
	}

	if err = g.SetCurrentView("toc"); err != nil {
		return errors.New("(findNotes) error setting current view to toc: " + err.Error())
	}

	if err = ShowNote(g); err != nil {
		return err
	}

	updateStatus(g, "")

	if err = UpdateFilterMsg(g, searchStr); err != nil {
		return err
	}

	return nil
}

func searchString(g *gocui.Gui, v *gocui.View) error {
	var err error

	if err = createInputView(g, "search", "Enter word(s) to search on:", true); err != nil {
		return err
	}

	updateStatus(g, "Enter a search string and press [Enter] to search.  Press [Esc] to abort.")

	return nil
}
