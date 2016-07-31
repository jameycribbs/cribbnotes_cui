package kbFunctions

import "github.com/jroimartin/gocui"

// KeybindingsNonVim contains all the non-vim bindings.
func KeybindingsNonVim(g *gocui.Gui) error {
	// Navigate between toc and note views.
	if err := g.SetKeybinding("toc", gocui.KeyCtrlSpace, gocui.ModNone, nextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("note", gocui.KeyCtrlSpace, gocui.ModNone, nextView); err != nil {
		return err
	}

	// Navigate inside toc view.
	if err := g.SetKeybinding("toc", gocui.KeyArrowDown, gocui.ModNone, tocCursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("toc", gocui.KeyArrowUp, gocui.ModNone, tocCursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("toc", gocui.KeyPgdn, gocui.ModNone, tocPageDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("toc", gocui.KeyPgup, gocui.ModNone, tocPageUp); err != nil {
		return err
	}

	// Navigate inside note view.
	if err := g.SetKeybinding("note", gocui.KeyPgdn, gocui.ModNone, notePageDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("note", gocui.KeyPgup, gocui.ModNone, notePageUp); err != nil {
		return err
	}

	// Save a record.
	if err := g.SetKeybinding("note", gocui.KeyCtrlS, gocui.ModNone, saveNote); err != nil {
		return err
	}

	// New record.
	if err := g.SetKeybinding("toc", gocui.KeyCtrlN, gocui.ModNone, newRec); err != nil {
		return err
	}
	if err := g.SetKeybinding("note", gocui.KeyCtrlN, gocui.ModNone, newRec); err != nil {
		return err
	}

	// Abort new title.
	if err := g.SetKeybinding("newTitle", gocui.KeyCtrlN, gocui.ModNone, AbortNewTitle); err != nil {
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
	return nil
}
