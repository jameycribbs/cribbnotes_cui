package kb_functions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jameycribbs/cribbnotes_cui/db"
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

	recs, err := db.Search("data", searchStr)
	if err != nil {
		return err
	}

	updateStatus(g, "Search for '"+searchStr+"' found "+strconv.Itoa(len(recs))+" notes.")

	if err := g.DeleteView("search"); err != nil {
		return err
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
