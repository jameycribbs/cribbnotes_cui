package kbFunctions

import "github.com/jroimartin/gocui"

func noteDeleteChar(g *gocui.Gui, v *gocui.View) error {
	if v.Editable {
		return nil
	}
	v.EditDelete(false)

	return nil
}

func noteCursorDown(g *gocui.Gui, v *gocui.View) error {
	if v.Editable {
		return nil
	}

	cx, cy := v.Cursor()

	line, err := v.Line(cy + 1)
	if err != nil {
		return nil
	}

	if len(line) == 0 {
		cx = 0
	} else if len(line)-1 < cx {
		cx = len(line) - 1
	}

	if err := v.SetCursor(cx, cy+1); err != nil {
		ox, oy := v.Origin()
		if err := v.SetOrigin(ox, oy+1); err != nil {
			return err
		}
	}
	return nil
}

func noteCursorLeft(g *gocui.Gui, v *gocui.View) error {
	if v.Editable {
		return nil
	}

	cx, cy := v.Cursor()

	if cx == 0 {
		return nil
	}

	if err := v.SetCursor(cx-1, cy); err != nil {
		ox, oy := v.Origin()
		if err := v.SetOrigin(ox-1, oy); err != nil {
			return err
		}
	}
	return nil
}

func noteCursorRight(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		if v.Editable {
			return nil
		}

		cx, cy := v.Cursor()

		line, err := v.Line(cy)
		if err != nil {
			return err
		}

		if len(line) < cx+2 {
			return nil
		}

		if err := v.SetCursor(cx+1, cy); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox+1, oy); err != nil {
				return err
			}
		}
	}
	return nil
}

func noteCursorUp(g *gocui.Gui, v *gocui.View) error {
	if v.Editable {
		return nil
	}

	cx, cy := v.Cursor()

	if cy == 0 {
		return nil
	}

	ox, oy := v.Origin()

	line, err := v.Line(cy - 1)
	if err != nil {
		return err
	}

	if len(line) == 0 {
		cx = 0
	} else if len(line)-1 < cx {
		cx = len(line) - 1
	}

	if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
		if err := v.SetOrigin(ox, oy-1); err != nil {
			return err
		}
	}

	return nil
}

func noteEnableEditable(g *gocui.Gui, v *gocui.View) error {
	if v.Editable {
		return nil
	}

	v.Editable = true

	updateStatus(g, "-- INSERT --")

	return nil
}

func noteEnableEditableNextChar(g *gocui.Gui, v *gocui.View) error {
	if v.Editable {
		return nil
	}

	v.Editable = true

	updateStatus(g, "-- INSERT --")

	cx, cy := v.Cursor()

	if err := v.SetCursor(cx+1, cy); err != nil {
		ox, oy := v.Origin()
		if err := v.SetOrigin(ox, oy); err != nil {
			return err
		}
	}

	return nil
}

func noteEnableEditableInsertAbove(g *gocui.Gui, v *gocui.View) error {
	if v.Editable {
		return nil
	}

	v.Editable = true

	updateStatus(g, "-- INSERT --")

	v.EditNewLine()

	_, cy := v.Cursor()

	if err := v.SetCursor(0, cy-1); err != nil {
		_, oy := v.Origin()
		if err := v.SetOrigin(0, oy-1); err != nil {
			return err
		}
	}

	return nil
}

func noteEnableEditableInsertBelow(g *gocui.Gui, v *gocui.View) error {
	if v.Editable {
		return nil
	}

	v.Editable = true

	updateStatus(g, "-- INSERT --")

	_, cy := v.Cursor()

	if err := v.SetCursor(0, cy+1); err != nil {
		_, oy := v.Origin()
		if err := v.SetOrigin(0, oy+1); err != nil {
			return err
		}
	}

	v.EditNewLine()

	_, cy = v.Cursor()

	if err := v.SetCursor(0, cy-1); err != nil {
		_, oy := v.Origin()
		if err := v.SetOrigin(0, oy-1); err != nil {
			return err
		}
	}

	return nil
}

func noteDisableEditable(g *gocui.Gui, v *gocui.View) error {
	noteView, err := g.View("noteDetail")
	if err != nil {
		return err
	}

	noteView.Editable = false

	updateStatus(g, "")

	return nil
}
