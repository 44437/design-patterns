package main

func newBaseComputerWithDisplayBuilder() IBuilder {
	return &BaseComputerWithDisplay{}
}

func (b *BaseComputerWithDisplay) buildRam() {
	b.Ram = "3GB"
}

func (b *BaseComputerWithDisplay) buildCPU() {
	b.CPU = "3.6GHz"
}

func (b *BaseComputerWithDisplay) buildHardDrive() {
	b.HardDrive = "1TB"
}

func (b *BaseComputerWithDisplay) buildDisplay() {
	b.Display = "32 inch 144hz 1ms"
}

func (b *BaseComputerWithDisplay) buildGraphicCard() {
	b.GraphicCard = "6GB"
}

func (b *BaseComputerWithDisplay) buildKeyboard() {}

func (b *BaseComputerWithDisplay) buildMouse() {}

func (b *BaseComputerWithDisplay) getComputer() Computer {
	return b
}
