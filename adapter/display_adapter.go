package main

type DisplayAdapter struct {
	Display *Display
}

func (d *DisplayAdapter) Play(work string) {
	d.Display.Show(work)
}
