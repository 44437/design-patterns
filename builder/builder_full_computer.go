package main

func newFullComputerBuilder() IBuilder {
	return &FullComputer{}
}

func (f *FullComputer) buildRam() {
	f.Ram = "3GB"
}

func (f *FullComputer) buildCPU() {
	f.CPU = "3.6GHz"
}

func (f *FullComputer) buildHardDrive() {
	f.HardDrive = "1TB"
}

func (f *FullComputer) buildDisplay() {
	f.Display = `32 inch 144hz 1ms`
}

func (f *FullComputer) buildGraphicCard() {
	f.GraphicCard = "6GB"
}

func (f *FullComputer) buildKeyboard() {
	f.Keyboard = "RGB Q Keyboard"
}

func (f *FullComputer) buildMouse() {
	f.Mouse = "RGB 8000 DPI"
}

func (f *FullComputer) getComputer() Computer {
	return f
}
