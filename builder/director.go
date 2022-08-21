package main

type Director struct {
	Builder IBuilder
}

func newDirector(builder IBuilder) *Director {
	return &Director{Builder: builder}
}

func (d *Director) setBuilder(builder IBuilder) {
	d.Builder = builder
}

func (d Director) buildComputer() Computer {
	d.Builder.buildRam()
	d.Builder.buildCPU()
	d.Builder.buildHardDrive()
	d.Builder.buildDisplay()
	d.Builder.buildGraphicCard()
	d.Builder.buildKeyboard()
	d.Builder.buildMouse()
	return d.Builder.getComputer()
}
