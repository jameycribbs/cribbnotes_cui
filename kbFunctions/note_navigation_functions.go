package kbFunctions

import "github.com/jroimartin/gocui"

func notePageDown(g *gocui.Gui, v *gocui.View) error {
	var err error

	_, vy := v.Size()

	for i := 0; i < vy-1; i++ {
		cx, cy := v.Cursor()

		if err = v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err = v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func notePageUp(g *gocui.Gui, v *gocui.View) error {
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
	return nil
}
