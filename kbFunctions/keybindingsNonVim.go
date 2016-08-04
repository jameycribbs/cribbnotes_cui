package kbFunctions

import (
	"errors"

	"github.com/jroimartin/gocui"
)

// KeybindingsNonVim contains all the non-vim bindings.
func KeybindingsNonVim(g *gocui.Gui) error {
	var err error

	// Navigate between toc and note views.
	if err = g.SetKeybinding("toc", gocui.KeyCtrlSpace, gocui.ModNone, nextView); err != nil {
		return errors.New("(KeybindingsNonVim) error setting keybinding for nextView: " + err.Error())
	}
	if err = g.SetKeybinding("noteDetail", gocui.KeyCtrlSpace, gocui.ModNone, nextView); err != nil {
		return errors.New("(KeybindingsNonVim) error setting keybinding for nextView: " + err.Error())
	}

	// Navigate inside toc view.
	if err = g.SetKeybinding("toc", gocui.KeyArrowDown, gocui.ModNone, tocCursorDown); err != nil {
		return errors.New("(KeybindingsNonVim) error setting keybinding for tocCursorDown: " + err.Error())
	}
	if err = g.SetKeybinding("toc", gocui.KeyArrowUp, gocui.ModNone, tocCursorUp); err != nil {
		return errors.New("(KeybindingsNonVim) error setting keybinding for tocCursorUp: " + err.Error())
	}
	if err = g.SetKeybinding("toc", gocui.KeyPgdn, gocui.ModNone, tocPageDown); err != nil {
		return errors.New("(KeybindingsNonVim) error setting keybinding for tocPageDown: " + err.Error())
	}
	if err = g.SetKeybinding("toc", gocui.KeyPgup, gocui.ModNone, tocPageUp); err != nil {
		return errors.New("(KeybindingsNonVim) error setting keybinding for tocPageUp: " + err.Error())
	}

	// Navigate inside note view.
	if err = g.SetKeybinding("noteDetail", gocui.KeyPgdn, gocui.ModNone, notePageDown); err != nil {
		return errors.New("(KeybindingsNonVim) error setting keybinding for notePageDown: " + err.Error())
	}
	if err = g.SetKeybinding("noteDetail", gocui.KeyPgup, gocui.ModNone, notePageUp); err != nil {
		return errors.New("(KeybindingsNonVim) error setting keybinding for notePageUp: " + err.Error())
	}

	// Save a record.
	if err = g.SetKeybinding("noteDetail", gocui.KeyCtrlS, gocui.ModNone, saveNote); err != nil {
		return errors.New("(KeybindingsNonVim) error setting keybinding for saveNote: " + err.Error())
	}

	// Search string.
	if err = g.SetKeybinding("toc", gocui.KeyCtrlF, gocui.ModNone, searchString); err != nil {
		return errors.New("(KeybindingsNonVim) error setting keybinding for searchString: " + err.Error())
	}
	if err = g.SetKeybinding("noteDetail", gocui.KeyCtrlF, gocui.ModNone, searchString); err != nil {
		return errors.New("(KeybindingsNonVim) error setting keybinding for searchString: " + err.Error())
	}
	if err = g.SetKeybinding("search", gocui.KeyCtrlX, gocui.ModNone, abortSearch); err != nil {
		return errors.New("(KeybindingsNonVim) error setting keybinding for AbortSearch: " + err.Error())
	}

	// Clear current filter.
	if err = g.SetKeybinding("toc", gocui.KeyCtrlC, gocui.ModNone, clearCurrentFilter); err != nil {
		return errors.New("(KeybindingsNonVim) error setting keybinding for clearCurrentFilter: " + err.Error())
	}
	if err = g.SetKeybinding("noteDetail", gocui.KeyCtrlC, gocui.ModNone, clearCurrentFilter); err != nil {
		return errors.New("(KeybindingsNonVim) error setting keybinding for clearCurrentFilter: " + err.Error())
	}
	// New record.
	if err = g.SetKeybinding("toc", gocui.KeyCtrlN, gocui.ModNone, newRec); err != nil {
		return errors.New("(KeybindingsNonVim) error setting keybinding for newRec: " + err.Error())
	}
	if err = g.SetKeybinding("noteDetail", gocui.KeyCtrlN, gocui.ModNone, newRec); err != nil {
		return errors.New("(KeybindingsNonVim) error setting keybinding for newRec: " + err.Error())
	}
	if err = g.SetKeybinding("newTitle", gocui.KeyCtrlX, gocui.ModNone, abortNewTitle); err != nil {
		return errors.New("(KeybindingsNonVim) error setting keybinding for abortNewTitle: " + err.Error())
	}

	return nil
}
