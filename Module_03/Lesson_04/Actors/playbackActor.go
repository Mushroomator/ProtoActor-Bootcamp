package actors04

import (
	"fmt"

	"github.com/AsynkronIT/protoactor-go/actor"
	msgs "github.com/Mushroomator/ProtoActorGo-Bootcamp/Module_03/Lesson_04/Msgs"
	consolepainter "github.com/Mushroomator/ProtoActorGo-Bootcamp/Module_03/consolePainter"
)

// Defining playback actor
type PlaybackActor struct{}

// Create new PlaybackActor
func NewPlaybackActor() *PlaybackActor {
	fmt.Println("Creating a PlaybackActor")
	return &PlaybackActor{}
}

// Process system message *actor.Started
func (state *PlaybackActor) ProcessStartedMessage(msg *actor.Started) {
	fmt.Printf("PlaybackActor started\n")
}

// Process PlayMovieMessage
func (state *PlaybackActor) ProcessPlayMovieMessage(msg *msgs.PlayMovieMessage) {
	cp := consolepainter.NewConsolePainter()
	cp.PrintfYellow("PlayMovieMessage %v for user %v\n", msg.MovieTitle, msg.UserId)
}

// Process StoppingMessage
func (state *PlaybackActor) ProcessStoppingMessage(msg *actor.Stopping) {
	fmt.Printf("PlaybackActor stopping\n")
}

// Process StoppedMessage
func (state *PlaybackActor) ProcessStoppedMessage(msg *actor.Stopped) {
	fmt.Printf("PlaybackActor stopped\n")
}

func (state *PlaybackActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Stopping:
		state.ProcessStoppingMessage(msg)
	case *actor.Stopped:
		state.ProcessStoppedMessage(msg)
	case *msgs.PlayMovieMessage:
		state.ProcessPlayMovieMessage(msg)
	}
}
