package kbFunctions

import "github.com/jroimartin/gocui"

// KeybindingsNonVim contains all the non-vim bindings.
func KeybindingsNonVim(g *gocui.Gui) error {
	// Navigate between toc and note views.
	if err := g.SetKeybinding("toc", gocui.KeyCtrlSpace, gocui.ModNone, nextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("noteDetail", gocui.KeyCtrlSpace, gocui.ModNone, nextView); err != nil {
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
	if err := g.SetKeybinding("noteDetail", gocui.KeyPgdn, gocui.ModNone, notePageDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("noteDetail", gocui.KeyPgup, gocui.ModNone, notePageUp); err != nil {
		return err
	}

	// Save a record.
	if err := g.SetKeybinding("noteDetail", gocui.KeyCtrlS, gocui.ModNone, saveNote); err != nil {
		return err
	}

	// Search string.
	if err := g.SetKeybinding("toc", gocui.KeyCtrlF, gocui.ModNone, searchString); err != nil {
		return err
	}
	if err := g.SetKeybinding("search", gocui.KeyCtrlX, gocui.ModNone, AbortSearch); err != nil {
		return err
	}

	// New record.
	if err := g.SetKeybinding("toc", gocui.KeyCtrlN, gocui.ModNone, newRec); err != nil {
		return err
	}
	if err := g.SetKeybinding("noteDetail", gocui.KeyCtrlN, gocui.ModNone, newRec); err != nil {
		return err
	}
	if err := g.SetKeybinding("newTitle", gocui.KeyCtrlX, gocui.ModNone, AbortNewTitle); err != nil {
		return err
	}

	return nil
}
