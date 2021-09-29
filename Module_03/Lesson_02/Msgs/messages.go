package msgs

// PLAYMOVIE
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

// RECOVERABLE
type RecoverableMessage struct{}

// Creates new Recoverable message
func NewRecoverable() *RecoverableMessage {
	return &RecoverableMessage{}
}
