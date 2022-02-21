package main

type presidency struct {
	president President
}

type Presidency interface {
	Get() President
}

func NewPresidency(realPresident President) Presidency {
	return &presidency{president: NewProxyPresident(realPresident)}
}

func (p *presidency) Get() President {
	return p.president
}
