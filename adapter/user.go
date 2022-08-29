package main

type User struct {
	player Player
}

func (u *User) setPlayer(player Player) {
	u.player = player
}

func (u *User) watch(work string) {
	u.player.Play(work)
}
