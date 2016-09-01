package kbFunctions

import (
	"errors"

	"github.com/jroimartin/gocui"
)

// KeybindingsVim contains all the vim bindings.
func KeybindingsVim(g *gocui.Gui) error {
	var dollarSignKey, zeroKey, colonKey, OKey, aKey, hKey, iKey, jKey, kKey, lKey, oKey, xKey rune = 36, 48, 58, 79, 97, 104, 105, 106, 107, 108, 111, 120
	var err error

	// Ex Commands
	if err = g.SetKeybinding("", colonKey, gocui.ModNone, exMode); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for exMode: " + err.Error())
	}
	if err = g.SetKeybinding("status", gocui.KeyEnter, gocui.ModNone, exExecuteCommand); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for exExecuteCommand: " + err.Error())
	}
	if err = g.SetKeybinding("status", gocui.KeyEsc, gocui.ModNone, abortExMode); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for abortExMode: " + err.Error())
	}

	// Navigate between toc and note views.
	if err = g.SetKeybinding("toc", gocui.KeyCtrlJ, gocui.ModNone, nextView); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for nextView: " + err.Error())
	}
	if err = g.SetKeybinding("noteDetail", gocui.KeyCtrlK, gocui.ModNone, nextView); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for nextView: " + err.Error())
	}

	// Navigate inside toc view.
	if err = g.SetKeybinding("toc", jKey, gocui.ModNone, tocCursorDown); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for tocCursorDown: " + err.Error())
	}
	if err = g.SetKeybinding("toc", kKey, gocui.ModNone, tocCursorUp); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for tocCursorUp: " + err.Error())
	}
	if err = g.SetKeybinding("toc", gocui.KeyCtrlF, gocui.ModNone, tocPageDown); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for tocPageDown: " + err.Error())
	}
	if err = g.SetKeybinding("toc", gocui.KeyCtrlB, gocui.ModNone, tocPageUp); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for tocPageUp: " + err.Error())
	}

	// Save a record.
	if err = g.SetKeybinding("noteDetail", gocui.KeyCtrlW, gocui.ModNone, saveNote); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for saveNote: " + err.Error())
	}

	// New record.
	if err = g.SetKeybinding("toc", gocui.KeyCtrlI, gocui.ModNone, newRec); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for newRec: " + err.Error())
	}
	if err = g.SetKeybinding("noteDetail", gocui.KeyCtrlI, gocui.ModNone, newRec); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for newRec: " + err.Error())
	}
	if err = g.SetKeybinding("newTitle", gocui.KeyEsc, gocui.ModNone, abortNewTitle); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for abortNewTitle: " + err.Error())
	}

	// Navigation inside note view
	if err = g.SetKeybinding("noteDetail", zeroKey, gocui.ModNone, noteBeginningOfLine); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for noteBeginningOfLine: " + err.Error())
	}
	if err = g.SetKeybinding("noteDetail", dollarSignKey, gocui.ModNone, noteEndOfLine); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for noteEndOfLine: " + err.Error())
	}
	if err = g.SetKeybinding("noteDetail", gocui.KeyCtrlF, gocui.ModNone, notePageDown); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for notePageDown: " + err.Error())
	}
	if err = g.SetKeybinding("noteDetail", gocui.KeyCtrlB, gocui.ModNone, notePageUp); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for notePageUp: " + err.Error())
	}
	if err = g.SetKeybinding("noteDetail", gocui.KeyEsc, gocui.ModNone, noteDisableEditable); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for noteDisableEditable: " + err.Error())
	}
	if err = g.SetKeybinding("noteDetail", hKey, gocui.ModNone, noteCursorLeft); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for noteCursorLeft: " + err.Error())
	}
	if err = g.SetKeybinding("noteDetail", iKey, gocui.ModNone, noteEnableEditable); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for noteEnableEditable: " + err.Error())
	}
	if err = g.SetKeybinding("noteDetail", jKey, gocui.ModNone, noteCursorDown); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for noteCursorDown: " + err.Error())
	}
	if err = g.SetKeybinding("noteDetail", kKey, gocui.ModNone, noteCursorUp); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for noteCursorUp: " + err.Error())
	}
	if err = g.SetKeybinding("noteDetail", lKey, gocui.ModNone, noteCursorRight); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for noteCursorRight: " + err.Error())
	}
	if err = g.SetKeybinding("noteDetail", aKey, gocui.ModNone, noteEnableEditableNextChar); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for noteEnableEditableNextChar: " + err.Error())
	}
	if err = g.SetKeybinding("noteDetail", OKey, gocui.ModNone, noteEnableEditableInsertAbove); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for noteEnableEditableInsertAbove: " + err.Error())
	}
	if err = g.SetKeybinding("noteDetail", oKey, gocui.ModNone, noteEnableEditableInsertBelow); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for noteEnableEditableInsertBelow: " + err.Error())
	}
	if err = g.SetKeybinding("noteDetail", xKey, gocui.ModNone, noteDeleteChar); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for noteDeleteChar: " + err.Error())
	}

	// Navigation inside edit note title view
	if err = g.SetKeybinding("editNoteTitle", aKey, gocui.ModNone, noteTitleEnableEditableNextChar); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for noteTitleEnableEditableNextChar: " + err.Error())
	}
	if err = g.SetKeybinding("editNoteTitle", hKey, gocui.ModNone, noteTitleCursorLeft); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for noteTitleCursorLeft: " + err.Error())
	}
	if err = g.SetKeybinding("editNoteTitle", iKey, gocui.ModNone, noteTitleEnableEditable); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for noteTitleEnableEditable: " + err.Error())
	}
	if err = g.SetKeybinding("editNoteTitle", lKey, gocui.ModNone, noteTitleCursorRight); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for noteTitleCursorRight: " + err.Error())
	}
	if err = g.SetKeybinding("editNoteTitle", xKey, gocui.ModNone, noteTitleDeleteChar); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for noteTitleDeleteChar: " + err.Error())
	}
	if err = g.SetKeybinding("editNoteTitle", gocui.KeyEsc, gocui.ModNone, noteTitleDisableEditable); err != nil {
		return errors.New("(KeybindingsVim) error setting keybinding for noteTitleDisableEditable: " + err.Error())
	}

	return nil
}
