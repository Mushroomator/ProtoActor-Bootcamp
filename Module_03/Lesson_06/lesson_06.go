package main

import (
	"fmt"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	actors "github.com/Mushroomator/ProtoActorGo-Bootcamp/Module_03/Lesson_06/Actors"
	msgs "github.com/Mushroomator/ProtoActorGo-Bootcamp/Module_03/Lesson_06/Msgs"
)

// Main function
func main() {
	// Creating actor system
	system := actor.NewActorSystem()
	fmt.Println("Actor system created")

	// Creating props for the playbackActor
	props := actor.PropsFromProducer(func() actor.Actor {
		return actors.NewUserActor()
	})

	// Creating the playbackActor and receiving its PID which we can use to reference the actor
	pid := system.Root.Spawn(props)

	// Send messages using PID
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Press [enter] to send PlayMovieMessage (The Movie)")
	fmt.Scanln()
	fmt.Println("Sending PlayMovieMessage (The Movie)")
	system.Root.Send(pid, msgs.NewPlayMovieMessage("The Movie", 44))
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Press [enter] to send PlayMovieMessage (The Movie 2)")
	fmt.Scanln()
	fmt.Println("Sending PlayMovieMessage (The Movie 2)")
	system.Root.Send(pid, msgs.NewPlayMovieMessage("The Movie 2", 54))

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Press [enter] to send StopMovieMessage")
	fmt.Scanln()
	fmt.Println("Sending a StopMovieMessage")
	system.Root.Send(pid, msgs.NewStopMovieMessage())

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Press [enter] to send another StopMovieMessage")
	fmt.Scanln()
	fmt.Println("Sending another StopMovieMessage")
	system.Root.Send(pid, msgs.NewStopMovieMessage())

	// Program will terminate immediately after main is run.
	// As all operations are non-blocking this fmt.Scanln() is necessary to prevent the program
	// from terminating before any other actions can be performed
	fmt.Scanln()
}
