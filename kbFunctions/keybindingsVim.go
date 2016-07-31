package kbFunctions

import "github.com/jroimartin/gocui"

// KeybindingsVim contains all the vim bindings.
func KeybindingsVim(g *gocui.Gui) error {
	var OKey, aKey, hKey, iKey, jKey, kKey, lKey, oKey, xKey rune = 79, 97, 104, 105, 106, 107, 108, 111, 120

	// Navigate between toc and note views.
	if err := g.SetKeybinding("toc", gocui.KeyCtrlJ, gocui.ModNone, nextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("note", gocui.KeyCtrlK, gocui.ModNone, nextView); err != nil {
		return err
	}

	// Navigate inside toc view.
	if err := g.SetKeybinding("toc", jKey, gocui.ModNone, tocCursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("toc", kKey, gocui.ModNone, tocCursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("toc", gocui.KeyCtrlF, gocui.ModNone, tocPageDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("toc", gocui.KeyCtrlB, gocui.ModNone, tocPageUp); err != nil {
		return err
	}

	// Save a record.
	if err := g.SetKeybinding("note", gocui.KeyCtrlW, gocui.ModNone, saveNote); err != nil {
		return err
	}

	// New record.
	if err := g.SetKeybinding("toc", gocui.KeyCtrlI, gocui.ModNone, newRec); err != nil {
		return err
	}
	if err := g.SetKeybinding("note", gocui.KeyCtrlI, gocui.ModNone, newRec); err != nil {
		return err
	}

	// Abort new title.
	if err := g.SetKeybinding("newTitle", gocui.KeyCtrlI, gocui.ModNone, AbortNewTitle); err != nil {
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

	// Navigation inside note view
	if err := g.SetKeybinding("note", gocui.KeyCtrlF, gocui.ModNone, notePageDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("note", gocui.KeyCtrlB, gocui.ModNone, notePageUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("note", gocui.KeyCtrlC, gocui.ModNone, noteDisableEditable); err != nil {
		return err
	}
	if err := g.SetKeybinding("note", hKey, gocui.ModNone, noteCursorLeft); err != nil {
		return err
	}
	if err := g.SetKeybinding("note", iKey, gocui.ModNone, noteEnableEditable); err != nil {
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
	if err := g.SetKeybinding("note", aKey, gocui.ModNone, noteEnableEditableNextChar); err != nil {
		return err
	}
	if err := g.SetKeybinding("note", OKey, gocui.ModNone, noteEnableEditableInsertAbove); err != nil {
		return err
	}
	if err := g.SetKeybinding("note", oKey, gocui.ModNone, noteEnableEditableInsertBelow); err != nil {
		return err
	}
	if err := g.SetKeybinding("note", xKey, gocui.ModNone, noteDeleteChar); err != nil {
		return err
	}
	return nil
}
