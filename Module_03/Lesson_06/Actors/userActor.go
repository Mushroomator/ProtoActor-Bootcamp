package actors

import (
	"fmt"

	"github.com/AsynkronIT/protoactor-go/actor"
	msgs "github.com/Mushroomator/ProtoActorGo-Bootcamp/Module_03/Lesson_06/Msgs"
	consolepainter "github.com/Mushroomator/ProtoActorGo-Bootcamp/Module_03/consolePainter"
)

type UserActor struct {
	behavior          actor.Behavior
	currentlyWatching *string
}

// Creates new UserActor
func NewUserActor() *UserActor {
	fmt.Printf("Creating a UserActor\n")
	cp := consolepainter.NewConsolePainter()
	cp.PrintfCyan("Setting initial behavior to stopped\n")
	// Create actor with behavior
	act := &UserActor{
		behavior: actor.NewBehavior(),
	}
	// Initialize actor with default behavior
	act.behavior.Become(act.Stopped)
	return act
}

// Processes all messages from mailbox and delegates them to
// the current behavior.
func (state *UserActor) Receive(ctx actor.Context) {
	// let the current behavior process ALL the messages
	state.behavior.Receive(ctx)
}

// Behavior for when actor is in "Playing" state.
func (state *UserActor) Playing(ctx actor.Context) {
	cp := consolepainter.NewConsolePainter()

	switch ctx.Message().(type) {
	case *msgs.PlayMovieMessage:
		cp.PrintfRed("Error: cannot start playing another movie before stopping existing one.\n")
	case *msgs.StopMovieMessage:
		cp.PrintfYellow("User has stopped watching %v\n", *state.currentlyWatching)
		state.currentlyWatching = nil
		state.behavior.Become(state.Stopped)
	}
	cp.PrintfCyan("UserActor has now become Playing\n")
}

// Behavior for when actor is in "Stopped" state.
func (state *UserActor) Stopped(ctx actor.Context) {
	cp := consolepainter.NewConsolePainter()

	switch msg := ctx.Message().(type) {
	case *msgs.PlayMovieMessage:
		state.currentlyWatching = &msg.MovieTitle
		cp.PrintfYellow("User is currently watching \"%v\"\n", *state.currentlyWatching)
		state.behavior.Become(state.Playing)
	case *msgs.StopMovieMessage:
		cp.PrintfRed("Error: cannot stop if nothing is playing\n")
	}
	cp.PrintfCyan("UserActor has now become Stopped\n")
}
