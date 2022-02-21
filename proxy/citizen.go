package main

type citizen struct {
	president President
}

type Citizen interface {
	TellTrouble(string)
	WantJob(string)
}

func NewCitizen(presidency Presidency) Citizen {
	president := presidency.Get()
	return &citizen{president: president}
}

func (c *citizen) TellTrouble(trouble string) {
	c.president.ListenTrouble(trouble)
}

func (c *citizen) WantJob(s string) {
	c.president.FindJob(s)
}
