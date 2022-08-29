package main

func main() {
	var videoPlayer *VideoPlayer
	var display *Display
	var user User

	var displayAdapter = &DisplayAdapter{Display: display}

	user.setPlayer(videoPlayer)
	user.watch("Game of Thrones")

	user.setPlayer(displayAdapter)
	user.watch("Breaking Bad")
}
