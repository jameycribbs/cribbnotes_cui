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
		updateStatus(g, "Ctrl+s - save note")

		return g.SetCurrentView("note")
	}

	updateStatus(g, "j/k - scroll | Enter - view note")

	return g.SetCurrentView("toc")
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
