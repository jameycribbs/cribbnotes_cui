package kbFunctions

import (
	"errors"

	"github.com/jroimartin/gocui"
)

// KeybindingsCommon contains all the commong bindings.
func KeybindingsCommon(g *gocui.Gui) error {
	var err error
	var nKey, yKey rune = 110, 121

	// Navigate between toc and note views.
	if err = g.SetKeybinding("toc", gocui.KeyEnter, gocui.ModNone, nextView); err != nil {
		return errors.New("(Keybindings) error setting keybinding for nextView: " + err.Error())
	}

	// Quit application.
	if err = g.SetKeybinding("", gocui.KeyCtrlQ, gocui.ModNone, quit); err != nil {
		return errors.New("(Keybindings) error setting keybinding for quit: " + err.Error())
	}

	// Create note.
	if err = g.SetKeybinding("newTitle", gocui.KeyEnter, gocui.ModNone, createNote); err != nil {
		return errors.New("(Keybindings) error setting keybinding for createNote: " + err.Error())
	}

	// Edit note title.
	if err = g.SetKeybinding("editNoteTitle", gocui.KeyEnter, gocui.ModNone, saveNoteTitle); err != nil {
		return errors.New("(Keybindings) error setting keybinding for saveNoteTitle: " + err.Error())
	}
	if err = g.SetKeybinding("toc", gocui.KeyCtrlT, gocui.ModNone, showEditNoteTitle); err != nil {
		return errors.New("(Keybindings) error setting keybinding for showEditNoteTitle: " + err.Error())
	}
	if err = g.SetKeybinding("noteDetail", gocui.KeyCtrlT, gocui.ModNone, showEditNoteTitle); err != nil {
		return errors.New("(Keybindings) error setting keybinding for showEditNoteTitle: " + err.Error())
	}
	if err = g.SetKeybinding("editNoteTitle", gocui.KeyEsc, gocui.ModNone, abortEditNoteTitle); err != nil {
		return errors.New("(Keybindings) error setting keybinding for abortEditNoteTitle: " + err.Error())
	}

	// Delete note.
	if err = g.SetKeybinding("toc", gocui.KeyCtrlD, gocui.ModNone, showDeleteNoteConfirm); err != nil {
		return errors.New("(Keybindings) error setting keybinding for showDeleteNoteConfirm: " + err.Error())
	}
	if err = g.SetKeybinding("noteDetail", gocui.KeyCtrlD, gocui.ModNone, showDeleteNoteConfirm); err != nil {
		return errors.New("(Keybindings) error setting keybinding for showDeleteNoteConfirm: " + err.Error())
	}
	// Confirm delete note.
	if err = g.SetKeybinding("deleteNoteConfirm", yKey, gocui.ModNone, deleteNote); err != nil {
		return errors.New("(Keybindings) error setting keybinding for deleteNote: " + err.Error())
	}
	if err = g.SetKeybinding("deleteNoteConfirm", nKey, gocui.ModNone, abortDeleteNote); err != nil {
		return errors.New("(Keybindings) error setting keybinding for abortDeleteNote: " + err.Error())
	}
	if err = g.SetKeybinding("deleteNoteConfirm", gocui.KeyEsc, gocui.ModNone, abortDeleteNote); err != nil {
		return errors.New("(Keybindings) error setting keybinding for abortDeleteNote: " + err.Error())
	}

	// Search string.
	if err = g.SetKeybinding("", gocui.KeyF3, gocui.ModNone, searchString); err != nil {
		return errors.New("(Keybindings) error setting keybinding for searchString: " + err.Error())
	}
	if err = g.SetKeybinding("search", gocui.KeyEsc, gocui.ModNone, abortSearch); err != nil {
		return errors.New("(Keybindings) error setting keybinding for AbortSearch: " + err.Error())
	}
	// Find notes.
	if err = g.SetKeybinding("search", gocui.KeyEnter, gocui.ModNone, findNotes); err != nil {
		return errors.New("(Keybindings) error setting keybinding for findNotes: " + err.Error())
	}
	// Clear current filter.
	if err = g.SetKeybinding("", gocui.KeyF4, gocui.ModNone, clearCurrentFilter); err != nil {
		return errors.New("(Keybindings) error setting keybinding for clearCurrentFilter: " + err.Error())
	}

	return nil
}
