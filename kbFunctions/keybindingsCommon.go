package kbFunctions

import "github.com/jroimartin/gocui"

// KeybindingsCommon contains all the commong bindings.
func KeybindingsCommon(g *gocui.Gui) error {
	// Navigate between toc and note views.
	if err := g.SetKeybinding("toc", gocui.KeyEnter, gocui.ModNone, nextView); err != nil {
		return err
	}

	// Quit application.
	if err := g.SetKeybinding("", gocui.KeyCtrlQ, gocui.ModNone, quit); err != nil {
		return err
	}

	// Create record.
	if err := g.SetKeybinding("newTitle", gocui.KeyEnter, gocui.ModNone, createNote); err != nil {
		return err
	}

	// Delete record.
	if err := g.SetKeybinding("toc", gocui.KeyCtrlD, gocui.ModNone, deleteRec); err != nil {
		return err
	}
	if err := g.SetKeybinding("note", gocui.KeyCtrlD, gocui.ModNone, deleteRec); err != nil {
		return err
	}

	// Find notes.
	if err := g.SetKeybinding("search", gocui.KeyEnter, gocui.ModNone, findNotes); err != nil {
		return err
	}
	return nil
}
