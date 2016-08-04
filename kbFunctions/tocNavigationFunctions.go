package kbFunctions

import (
	"errors"

	"github.com/jroimartin/gocui"
)

func tocCursorDown(g *gocui.Gui, v *gocui.View) error {
	var line string
	var err error
	var cx, cy, ox, oy int

	if v != nil {
		cx, cy = v.Cursor()

		if line, err = v.Line(cy + 1); err != nil {
			// if error, means toc is empty so do not scroll.
			return nil
		}

		if line != "" {
			if err = v.SetCursor(cx, cy+1); err != nil {
				ox, oy = v.Origin()
				if err = v.SetOrigin(ox, oy+1); err != nil {
					return errors.New("(tocCursorDown) error setting origin: " + err.Error())
				}
			}
		}

		if err = ShowNote(g); err != nil {
			return err
		}

	}
	return nil
}

func tocCursorUp(g *gocui.Gui, v *gocui.View) error {
	var ox, oy, cx, cy int
	var err error

	if v != nil {
		ox, oy = v.Origin()
		cx, cy = v.Cursor()
		if err = v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err = v.SetOrigin(ox, oy-1); err != nil {
				return errors.New("(tocCursorUp) error setting origin: " + err.Error())
			}
		}

		if err = ShowNote(g); err != nil {
			return err
		}

	}
	return nil
}

func tocPageDown(g *gocui.Gui, v *gocui.View) error {
	var line string
	var err error
	var vy, cx, cy, ox, oy int

	_, vy = v.Size()

	for i := 0; i < vy-1; i++ {
		cx, cy = v.Cursor()

		if line, err = v.Line(cy + 1); err != nil {
			break
		}

		if line != "" {
			if err = v.SetCursor(cx, cy+1); err != nil {
				ox, oy = v.Origin()
				if err = v.SetOrigin(ox, oy+1); err != nil {
					return errors.New("(tocPageDown) error setting origin: " + err.Error())
				}
			}
		}
	}

	if err = ShowNote(g); err != nil {
		return err
	}
	return nil
}

func tocPageUp(g *gocui.Gui, v *gocui.View) error {
	var vy, ox, oy, cx, cy int
	var err error

	_, vy = v.Size()

	for i := 0; i < vy-1; i++ {
		ox, oy = v.Origin()
		cx, cy = v.Cursor()
		if err = v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err = v.SetOrigin(ox, oy-1); err != nil {
				return errors.New("(tocPageUp) error setting origin: " + err.Error())
			}
		}
	}

	if err = ShowNote(g); err != nil {
		return err
	}
	return nil
}
