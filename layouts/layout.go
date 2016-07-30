package layouts

import "github.com/jroimartin/gocui"

func Layout(g *gocui.Gui) error {
	if err := headlineLayout(g); err != nil {
		return err
	}

	if err := tocLayout(g); err != nil {
		return err
	}

	if err := keybindingsLayout(g); err != nil {
		return err
	}

	if err := noteTitleLayout(g); err != nil {
		return err
	}

	if err := noteNumberLayout(g); err != nil {
		return err
	}

	if err := noteLayout(g); err != nil {
		return err
	}

	if err := statusLayout(g); err != nil {
		return err
	}

	return nil
}
