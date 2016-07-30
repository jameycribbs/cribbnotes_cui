package kb_functions

import "github.com/jroimartin/gocui"

func Keybindings(g *gocui.Gui) error {
	var jKey, kKey, slashKey rune = 106, 107, 47

	// Navigate between toc and note views.
	if err := g.SetKeybinding("toc", gocui.KeyCtrlSpace, gocui.ModNone, nextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("note", gocui.KeyCtrlSpace, gocui.ModNone, nextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("toc", gocui.KeyCtrlJ, gocui.ModNone, nextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("note", gocui.KeyCtrlK, gocui.ModNone, nextView); err != nil {
		return err
	}

	// Navigate inside toc view.
	if err := g.SetKeybinding("toc", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("toc", jKey, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("toc", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("toc", kKey, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("toc", gocui.KeyPgdn, gocui.ModNone, tocPageDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("toc", gocui.KeyCtrlF, gocui.ModNone, tocPageDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("toc", gocui.KeyPgup, gocui.ModNone, tocPageUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("toc", gocui.KeyCtrlB, gocui.ModNone, tocPageUp); err != nil {
		return err
	}

	// Quit application.
	if err := g.SetKeybinding("", gocui.KeyCtrlQ, gocui.ModNone, quit); err != nil {
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

	// Search string.
	if err := g.SetKeybinding("toc", slashKey, gocui.ModNone, searchString); err != nil {
		return err
	}

	// Abort Search.
	if err := g.SetKeybinding("search", slashKey, gocui.ModNone, AbortSearch); err != nil {
		return err
	}

	// Find notes.
	if err := g.SetKeybinding("search", gocui.KeyEnter, gocui.ModNone, findNotes); err != nil {
		return err
	}

	return nil
}
