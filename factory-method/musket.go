package main

type musket struct {
	gun
}

func NewMusket() Gun {
	return &musket{
		gun: gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
