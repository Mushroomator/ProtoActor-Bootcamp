package actors

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	msgs "github.com/Mushroomator/ProtoActorGo-Bootcamp/Module_03/Lesson_02/Msgs"
	consolepainter "github.com/Mushroomator/ProtoActorGo-Bootcamp/Module_03/consolePainter"
)

// Defining playback actor
type PlaybackActor struct{}

// Process system message *actor.Started
func (state *PlaybackActor) ProcessStartedMessage(msg *actor.Started) {
	cp := consolepainter.NewConsolePainter()
	cp.PrintfGreen("PlaybackActor started\n")
}

// Process system message *actor.Restarting
func (state *PlaybackActor) ProcessRecoverableMessage(ctx actor.Context, msg msgs.RecoverableMessage) {
	var child *actor.PID
	if children := ctx.Children(); len(children) == 0 {
		// this actor does not have any children yet
		props := actor.PropsFromProducer(func() actor.Actor {
			return NewChildActor()
		})
		// spawn a child actor
		child = ctx.Spawn(props)
	} else {
		// Get first child
		child = children[0]
	}
	// forward the current message of type Recoverable to the child actor
	ctx.Forward(child)
}

// Process PlayMovieMessage
func (state *PlaybackActor) ProcessPlayMovieMessage(msg *msgs.PlayMovieMessage) {
	cp := consolepainter.NewConsolePainter()
	cp.PrintfYellow("PlayMovieMessage %v for user %v\n", msg.MovieTitle, msg.UserId)
}

// Process StoppingMessage
func (state *PlaybackActor) ProcessStoppingMessage(msg *actor.Stopping) {
	cp := consolepainter.NewConsolePainter()
	cp.PrintfGreen("PlaybackActor stopping\n")
}

func (state *PlaybackActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		state.ProcessStartedMessage(msg)
	case msgs.RecoverableMessage:
		state.ProcessRecoverableMessage(ctx, msg)
	case *actor.Stopping:
		state.ProcessStoppingMessage(msg)
	case *msgs.PlayMovieMessage:
		state.ProcessPlayMovieMessage(msg)
	}

}
