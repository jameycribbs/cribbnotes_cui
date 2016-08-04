package kbFunctions

import "github.com/jroimartin/gocui"

// KeybindingsVim contains all the vim bindings.
func KeybindingsVim(g *gocui.Gui) error {
	var OKey, aKey, hKey, iKey, jKey, kKey, lKey, oKey, xKey rune = 79, 97, 104, 105, 106, 107, 108, 111, 120

	// Navigate between toc and note views.
	if err := g.SetKeybinding("toc", gocui.KeyCtrlJ, gocui.ModNone, nextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("noteDetail", gocui.KeyCtrlK, gocui.ModNone, nextView); err != nil {
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
	if err := g.SetKeybinding("noteDetail", gocui.KeyCtrlW, gocui.ModNone, saveNote); err != nil {
		return err
	}

	// New record.
	if err := g.SetKeybinding("toc", gocui.KeyCtrlI, gocui.ModNone, newRec); err != nil {
		return err
	}
	if err := g.SetKeybinding("noteDetail", gocui.KeyCtrlI, gocui.ModNone, newRec); err != nil {
		return err
	}
	if err := g.SetKeybinding("newTitle", gocui.KeyCtrlX, gocui.ModNone, AbortNewTitle); err != nil {
		return err
	}

	// Search string.
	if err := g.SetKeybinding("toc", gocui.KeyCtrlSlash, gocui.ModNone, searchString); err != nil {
		return err
	}
	if err := g.SetKeybinding("search", gocui.KeyCtrlX, gocui.ModNone, AbortSearch); err != nil {
		return err
	}

	// Navigation inside note view
	if err := g.SetKeybinding("noteDetail", gocui.KeyCtrlF, gocui.ModNone, notePageDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("noteDetail", gocui.KeyCtrlB, gocui.ModNone, notePageUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("noteDetail", gocui.KeyCtrlC, gocui.ModNone, noteDisableEditable); err != nil {
		return err
	}
	if err := g.SetKeybinding("noteDetail", hKey, gocui.ModNone, noteCursorLeft); err != nil {
		return err
	}
	if err := g.SetKeybinding("noteDetail", iKey, gocui.ModNone, noteEnableEditable); err != nil {
		return err
	}
	if err := g.SetKeybinding("noteDetail", jKey, gocui.ModNone, noteCursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("noteDetail", kKey, gocui.ModNone, noteCursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("noteDetail", lKey, gocui.ModNone, noteCursorRight); err != nil {
		return err
	}
	if err := g.SetKeybinding("noteDetail", aKey, gocui.ModNone, noteEnableEditableNextChar); err != nil {
		return err
	}
	if err := g.SetKeybinding("noteDetail", OKey, gocui.ModNone, noteEnableEditableInsertAbove); err != nil {
		return err
	}
	if err := g.SetKeybinding("noteDetail", oKey, gocui.ModNone, noteEnableEditableInsertBelow); err != nil {
		return err
	}
	if err := g.SetKeybinding("noteDetail", xKey, gocui.ModNone, noteDeleteChar); err != nil {
		return err
	}

	// Navigation inside edit note title view
	if err := g.SetKeybinding("editNoteTitle", aKey, gocui.ModNone, noteTitleEnableEditableNextChar); err != nil {
		return err
	}
	if err := g.SetKeybinding("editNoteTitle", hKey, gocui.ModNone, noteTitleCursorLeft); err != nil {
		return err
	}
	if err := g.SetKeybinding("editNoteTitle", iKey, gocui.ModNone, noteTitleEnableEditable); err != nil {
		return err
	}
	if err := g.SetKeybinding("editNoteTitle", lKey, gocui.ModNone, noteTitleCursorRight); err != nil {
		return err
	}
	if err := g.SetKeybinding("editNoteTitle", xKey, gocui.ModNone, noteTitleDeleteChar); err != nil {
		return err
	}
	if err := g.SetKeybinding("editNoteTitle", gocui.KeyCtrlC, gocui.ModNone, noteTitleDisableEditable); err != nil {
		return err
	}

	return nil
}
