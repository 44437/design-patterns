package main

func newBaseComputerBuilder() IBuilder {
	return &BaseComputer{}
}

func (b *BaseComputer) buildRam() {
	b.Ram = "3GB"
}

func (b *BaseComputer) buildCPU() {
	b.CPU = "3.6GHz"
}

func (b *BaseComputer) buildHardDrive() {
	b.HardDrive = "1TB"
}

func (b *BaseComputer) buildDisplay() {}

func (b *BaseComputer) buildGraphicCard() {}

func (b *BaseComputer) buildKeyboard() {}

func (b *BaseComputer) buildMouse() {}

func (b *BaseComputer) getComputer() Computer {
	return b
}
