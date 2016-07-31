package kb_functions

import "github.com/jroimartin/gocui"

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	var line string
	var err error

	if v != nil {
		cx, cy := v.Cursor()

		if line, err = v.Line(cy + 1); err != nil {
			// if error, means toc is empty so do not scroll.
			return nil
		}

		if line != "" {
			if err := v.SetCursor(cx, cy+1); err != nil {
				ox, oy := v.Origin()
				if err := v.SetOrigin(ox, oy+1); err != nil {
					return err
				}
			}
		}

		if err := ShowNote(g); err != nil {
			return err
		}

	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}

		if err := ShowNote(g); err != nil {
			return err
		}

	}
	return nil
}

func nextView(g *gocui.Gui, v *gocui.View) error {
	if v == nil || v.Name() == "toc" {
		updateStatus(g, "Ctrl+s to save note.")

		return g.SetCurrentView("note")
	}

	updateStatus(g, "")

	return g.SetCurrentView("toc")
}

func noteCursorDown(g *gocui.Gui, v *gocui.View) error {
	var line string
	var err error

	if v != nil {
		if v.Editable {
			return nil
		}

		cx, cy := v.Cursor()

		if line, err = v.Line(cy + 1); err != nil {
			return nil
		}

		if line != "" {
			if err := v.SetCursor(cx, cy+1); err != nil {
				ox, oy := v.Origin()
				if err := v.SetOrigin(ox, oy+1); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func noteCursorLeft(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
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
	}
	return nil
}

func noteCursorRight(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		if v.Editable {
			return nil
		}

		cx, cy := v.Cursor()

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
	if v != nil {
		if v.Editable {
			return nil
		}

		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

func noteEnableEditable(g *gocui.Gui, v *gocui.View) error {
	noteView, err := g.View("note")
	if err != nil {
		return err
	}

	noteView.Editable = true

	updateStatus(g, "-- INSERT --")

	return nil
}

func noteDisableEditable(g *gocui.Gui, v *gocui.View) error {
	noteView, err := g.View("note")
	if err != nil {
		return err
	}

	noteView.Editable = false

	updateStatus(g, "Ctrl+s to save note.")

	return nil
}

func tocPageDown(g *gocui.Gui, v *gocui.View) error {
	var line string
	var err error

	if v != nil {
		_, vy := v.Size()

		for i := 0; i < vy-1; i++ {
			cx, cy := v.Cursor()

			if line, err = v.Line(cy + 1); err != nil {
				break
			}

			if line != "" {
				if err := v.SetCursor(cx, cy+1); err != nil {
					ox, oy := v.Origin()
					if err := v.SetOrigin(ox, oy+1); err != nil {
						return err
					}
				}
			}
		}

		if err := ShowNote(g); err != nil {
			return err
		}

	}
	return nil
}

func tocPageUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		_, vy := v.Size()

		for i := 0; i < vy-1; i++ {
			ox, oy := v.Origin()
			cx, cy := v.Cursor()
			if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
				if err := v.SetOrigin(ox, oy-1); err != nil {
					return err
				}
			}
		}

		if err := ShowNote(g); err != nil {
			return err
		}

	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
