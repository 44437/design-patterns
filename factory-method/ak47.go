package main

type ak47 struct {
	gun
}

func NewAk47() Gun {
	return &ak47{
		gun: gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
