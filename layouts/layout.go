package layouts

import "github.com/jroimartin/gocui"

// Layout calls all the individual layout functions to create all the views.
func Layout(g *gocui.Gui) error {
	var err error

	if err = headlineLayout(g); err != nil {
		return err
	}

	if err = tocLayout(g); err != nil {
		return err
	}

	if err = keybindingsLayout(g); err != nil {
		return err
	}

	if err = noteTitleLayout(g); err != nil {
		return err
	}

	if err = noteNumberLayout(g); err != nil {
		return err
	}

	if err = noteDetailLayout(g); err != nil {
		return err
	}

	if err = statusLayout(g); err != nil {
		return err
	}

	return nil
}
