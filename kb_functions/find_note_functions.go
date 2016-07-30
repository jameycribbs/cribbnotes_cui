package kb_functions

import (
	"errors"
	"strings"

	"github.com/jroimartin/gocui"
)

func AbortSearch(g *gocui.Gui, v *gocui.View) error {
	if err := g.DeleteView("search"); err != nil {
		return err
	}

	if err := g.SetCurrentView("toc"); err != nil {
		return err
	}
	return nil
}

func findNotes(g *gocui.Gui, v *gocui.View) error {
	searchView, err := g.View("search")
	if err != nil {
		return errors.New("error grabbing handle to search view: " + err.Error())
	}

	searchStr := strings.TrimSuffix(searchView.ViewBuffer(), "\n")

	updateStatus(g, "Searched for '"+searchStr+"'.")

	if err := g.DeleteView("search"); err != nil {
		return errors.New("error deleting search view: " + err.Error())
	}

	if err := PopulateToc(g, searchStr); err != nil {
		return err
	}

	if err := g.SetCurrentView("toc"); err != nil {
		return errors.New("error setting current view to toc: " + err.Error())
	}

	if err := ShowNote(g); err != nil {
		return err
	}

	return nil
}

func searchString(g *gocui.Gui, v *gocui.View) error {
	if err := createInputView(g, "search", "Enter word(s) to search on:"); err != nil {
		return err
	}

	updateStatus(g, "Enter a search string.  Press [Enter] to search.")

	return nil
}
