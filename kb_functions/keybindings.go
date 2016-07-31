package kb_functions

import "github.com/jroimartin/gocui"

func Keybindings(g *gocui.Gui) error {
	var hKey, iKey, jKey, kKey, lKey rune = 104, 105, 106, 107, 108

	// Navigate between toc and note views.
	if err := g.SetKeybinding("toc", gocui.KeyEnter, gocui.ModNone, nextView); err != nil {
		return err
	}
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
	if err := g.SetKeybinding("toc", gocui.KeyCtrlSlash, gocui.ModNone, searchString); err != nil {
		return err
	}

	// Abort Search.
	if err := g.SetKeybinding("search", gocui.KeyCtrlSlash, gocui.ModNone, AbortSearch); err != nil {
		return err
	}

	// Find notes.
	if err := g.SetKeybinding("search", gocui.KeyEnter, gocui.ModNone, findNotes); err != nil {
		return err
	}

	// Navigation inside note view
	if err := g.SetKeybinding("note", hKey, gocui.ModNone, noteCursorLeft); err != nil {
		return err
	}
	if err := g.SetKeybinding("note", iKey, gocui.ModNone, noteEnableEditable); err != nil {
		return err
	}
	if err := g.SetKeybinding("note", gocui.KeyCtrlC, gocui.ModNone, noteDisableEditable); err != nil {
		return err
	}
	if err := g.SetKeybinding("note", jKey, gocui.ModNone, noteCursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("note", kKey, gocui.ModNone, noteCursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("note", lKey, gocui.ModNone, noteCursorRight); err != nil {
		return err
	}

	return nil
}
