// Lesson 6: PlaybackActor with string and int messages
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
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		// In C# we can send this within the constructor but as there is no such thing as an constructor
		// for structs in Go we must log this message when we receive the *actor.Started message which is
		// a system message sent whenever the actor is created

		// There is an alternative: one could create a func NewPlaybackActor() *PlaybackActor message which creates the playback actor
		// struct and therefore acts as an constructor. But there is nothing that prevents the user from creating the struct himself/herself.
		fmt.Println("Creating a PlaybackActor")
	case string:
		// Received message is of type string
		fmt.Printf("Reiceived movie title %v\n", msg)
	case int:
		// Received message is of type int
		fmt.Printf("Reiceived user ID %v\n", msg)
	}

}

func main() {
	// Creating actor system
	system := actor.NewActorSystem()
	fmt.Println("Actor system created")

	// Creating props for the playbackActor
	props := actor.PropsFromProducer(func() actor.Actor {
		return &PlaybackActor{}
		// Alternative:
		// return NewPlaybackActor()
	})

	// Creating the playbackActor and receiving its PID which we can use to reference the actor
	pid := system.Root.Spawn(props)

	// Send messages using PID
	// Message of type string
	system.Root.Send(pid, "The Movie")
	// Message of type int
	system.Root.Send(pid, 44)

	// Program will terminate immediately after main is run.
	// As all operations are non-blocking this fmt.Scanln() is necessary to prevent the program
	// from terminating before any other actions can be performed
	fmt.Scanln()
}
