package kb_functions

import "github.com/jroimartin/gocui"

func Keybindings(g *gocui.Gui) error {
	var jKey, kKey, slashKey rune = 106, 107, 47

	// Navigate between toc and main views.
	if err := g.SetKeybinding("toc", gocui.KeyCtrlSpace, gocui.ModNone, nextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("main", gocui.KeyCtrlSpace, gocui.ModNone, nextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("toc", gocui.KeyCtrlJ, gocui.ModNone, nextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("main", gocui.KeyCtrlK, gocui.ModNone, nextView); err != nil {
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

	// Quit application.
	if err := g.SetKeybinding("", gocui.KeyCtrlQ, gocui.ModNone, quit); err != nil {
		return err
	}

	// Show a record.
	if err := g.SetKeybinding("toc", gocui.KeyEnter, gocui.ModNone, showRec); err != nil {
		return err
	}

	// Save a record.
	if err := g.SetKeybinding("main", gocui.KeyCtrlS, gocui.ModNone, saveRec); err != nil {
		return err
	}

	// New record.
	if err := g.SetKeybinding("toc", gocui.KeyCtrlN, gocui.ModNone, newRec); err != nil {
		return err
	}
	if err := g.SetKeybinding("main", gocui.KeyCtrlN, gocui.ModNone, newRec); err != nil {
		return err
	}

	// Abort new title.
	if err := g.SetKeybinding("newTitle", gocui.KeyCtrlN, gocui.ModNone, AbortNewTitle); err != nil {
		return err
	}

	// Create record.
	if err := g.SetKeybinding("newTitle", gocui.KeyEnter, gocui.ModNone, createRec); err != nil {
		return err
	}

	// Get file id to delete.
	if err := g.SetKeybinding("toc", gocui.KeyCtrlD, gocui.ModNone, getFileIdToDelete); err != nil {
		return err
	}
	if err := g.SetKeybinding("main", gocui.KeyCtrlD, gocui.ModNone, getFileIdToDelete); err != nil {
		return err
	}

	// Abort get file id to delete.
	if err := g.SetKeybinding("fileIdToDelete", gocui.KeyCtrlD, gocui.ModNone, AbortGetFileIdToDelete); err != nil {
		return err
	}

	// Delete record.
	if err := g.SetKeybinding("fileIdToDelete", gocui.KeyEnter, gocui.ModNone, deleteRec); err != nil {
		return err
	}

	// Search string.
	if err := g.SetKeybinding("toc", gocui.KeyCtrlF, gocui.ModNone, searchString); err != nil {
		return err
	}
	if err := g.SetKeybinding("toc", slashKey, gocui.ModNone, searchString); err != nil {
		return err
	}
	if err := g.SetKeybinding("main", gocui.KeyCtrlF, gocui.ModNone, searchString); err != nil {
		return err
	}
	if err := g.SetKeybinding("main", slashKey, gocui.ModNone, searchString); err != nil {
		return err
	}

	// Abort Search.
	if err := g.SetKeybinding("search", gocui.KeyCtrlF, gocui.ModNone, AbortSearch); err != nil {
		return err
	}
	if err := g.SetKeybinding("search", slashKey, gocui.ModNone, AbortSearch); err != nil {
		return err
	}

	// Find notes.
	if err := g.SetKeybinding("search", gocui.KeyEnter, gocui.ModNone, findNotes); err != nil {
		return err
	}

	return nil
}
