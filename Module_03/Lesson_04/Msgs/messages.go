package msgs

// PLAY MOVIE
// Defining custom message PlayMovieMessage
type PlayMovieMessage struct {
	MovieTitle string
	UserId     int
}

// Creates new PlayMovieMessage
func NewPlayMovieMessage(movieTitle string, userId int) *PlayMovieMessage {
	return &PlayMovieMessage{
		MovieTitle: movieTitle,
		UserId:     userId,
	}
}
