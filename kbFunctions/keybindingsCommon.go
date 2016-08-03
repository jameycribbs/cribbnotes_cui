package kbFunctions

import "github.com/jroimartin/gocui"

// KeybindingsCommon contains all the commong bindings.
func KeybindingsCommon(g *gocui.Gui) error {
	var nKey, yKey rune = 110, 121

	// Navigate between toc and note views.
	if err := g.SetKeybinding("toc", gocui.KeyEnter, gocui.ModNone, nextView); err != nil {
		return err
	}

	// Quit application.
	if err := g.SetKeybinding("", gocui.KeyCtrlQ, gocui.ModNone, quit); err != nil {
		return err
	}

	// Create note.
	if err := g.SetKeybinding("newTitle", gocui.KeyEnter, gocui.ModNone, createNote); err != nil {
		return err
	}

	// Edit note title.
	if err := g.SetKeybinding("editNoteTitle", gocui.KeyEnter, gocui.ModNone, saveNoteTitle); err != nil {
		return err
	}
	if err := g.SetKeybinding("toc", gocui.KeyCtrlT, gocui.ModNone, showEditNoteTitle); err != nil {
		return err
	}
	if err := g.SetKeybinding("note", gocui.KeyCtrlT, gocui.ModNone, showEditNoteTitle); err != nil {
		return err
	}
	if err := g.SetKeybinding("editNoteTitle", gocui.KeyCtrlC, gocui.ModNone, AbortEditNoteTitle); err != nil {
		return err
	}

	// Delete note.
	if err := g.SetKeybinding("toc", gocui.KeyCtrlD, gocui.ModNone, showDeleteNoteConfirm); err != nil {
		return err
	}
	if err := g.SetKeybinding("note", gocui.KeyCtrlD, gocui.ModNone, showDeleteNoteConfirm); err != nil {
		return err
	}
	// Confirm delete note.
	if err := g.SetKeybinding("deleteNoteConfirm", yKey, gocui.ModNone, deleteNote); err != nil {
		return err
	}
	if err := g.SetKeybinding("deleteNoteConfirm", nKey, gocui.ModNone, abortDeleteNote); err != nil {
		return err
	}

	// Find notes.
	if err := g.SetKeybinding("search", gocui.KeyEnter, gocui.ModNone, findNotes); err != nil {
		return err
	}
	return nil
}
