// Lesson 7: PlaybackActor with custom message
package main

import (
	"fmt"

	"github.com/AsynkronIT/protoactor-go/actor"
)

// Defining playback actor
type PlaybackActor struct{}

// Defining custom message PlayMovieMessage
type PlayMovieMessage struct {
	MovieTitle string
	UserId     int
}

func (state *PlaybackActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *PlayMovieMessage:
		// Received message is of type *PlayMovieMessage
		fmt.Printf("Reiceived movie title %v\n", msg.MovieTitle)
		fmt.Printf("Reiceived user ID %v\n", msg.UserId)
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
	pid := system.Root.Spawn(props)

	// Send messages using PID
	// Message of type PlayMovieMessage
	system.Root.Send(pid, &PlayMovieMessage{
		MovieTitle: "The Movie",
		UserId:     44,
	})

	// Program will terminate immediately after main is run.
	// As all operations are non-blocking this fmt.Scanln() is necessary to prevent the program
	// from terminating before any other actions can be performed
	fmt.Scanln()
}
