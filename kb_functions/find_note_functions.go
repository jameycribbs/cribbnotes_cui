package kb_functions

import (
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
		panic(err)
	}

	searchStr := strings.TrimSuffix(searchView.ViewBuffer(), "\n")

	updateStatus(g, "Searched for '"+searchStr+"'.")

	if err := g.DeleteView("search"); err != nil {
		return err
	}

	if err := PopulateToc(g, searchStr); err != nil {
		return err
	}

	if err := g.SetCurrentView("toc"); err != nil {
		return err
	}

	return nil
}

func searchString(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := g.Size()

	if searchView, err := g.SetView("search", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		searchView.Title = "Enter string to search on:"
		searchView.Editable = true

		updateStatus(g, "Enter a search string.  Press [Enter] to search.")

		if err := g.SetCurrentView("search"); err != nil {
			return err
		}
	}
	return nil
}
