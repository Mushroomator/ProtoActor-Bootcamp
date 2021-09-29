package main

import (
	"fmt"

	"github.com/AsynkronIT/protoactor-go/actor"
	actors "github.com/Mushroomator/ProtoActorGo-Bootcamp/Module_03/Lesson_04/Actors"
	msgs "github.com/Mushroomator/ProtoActorGo-Bootcamp/Module_03/Lesson_04/Msgs"
)

// Main function
func main() {
	// Creating actor system
	system := actor.NewActorSystem()
	fmt.Println("Actor system created")

	// Creating props for the playbackActor
	props := actor.PropsFromProducer(func() actor.Actor {
		return actors.NewPlaybackActor()
	})

	// Creating the playbackActor and receiving its PID which we can use to reference the actor
	pid := system.Root.Spawn(props)

	// Send messages using PID
	system.Root.Send(pid, msgs.NewPlayMovieMessage("The Movie", 44))

	// Send poison pill message
	system.Root.Poison(pid)

	fmt.Scanln()
}
