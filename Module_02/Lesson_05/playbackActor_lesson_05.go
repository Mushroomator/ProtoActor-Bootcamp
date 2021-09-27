// Lesson 5: PlaybackActor
package main

import (
	"fmt"

	"github.com/AsynkronIT/protoactor-go/actor"
)

// Defining playback actor
type PlaybackActor struct{}

// Constructor for PlaybackActor
func NewPlaybackActor() *PlaybackActor {
	// Alternative: print creation message here
	// fmt.Println("Creating a PlaybackActor")
	return &PlaybackActor{}
}

func (state *PlaybackActor) Receive(ctx actor.Context) {
	switch ctx.Message().(type) {
	case *actor.Started:
		// In C# we can send this within the constructor but as there is no such thing as an constructor
		// for structs in Go we must log this message when we receive the *actor.Started message which is
		// a system message sent whenever the actor is created

		// There is an alternative: one could create a func NewPlaybackActor() *PlaybackActor message which creates the playback actor
		// struct and therefore acts as an constructor. But there is nothing that prevents the user from creating the struct himself/herself.
		fmt.Println("Creating a PlaybackActor")
	}

}

func main() {
	// Creating actor system
	system := actor.NewActorSystem()
	fmt.Println("Actor system created")

	// Creating props for the playbackActor
	props := actor.PropsFromProducer(func() actor.Actor {
		return &PlaybackActor{}
	})

	// Creating the playbackActor and receiving its PID which we can use to reference the actor
	system.Root.Spawn(props)

	// Program will terminate immediately after main is run.
	// As all operations are non-blocking this fmt.Scanln() is necessary to prevent the program
	// from terminating before any other actions can be performed
	fmt.Scanln()
}
