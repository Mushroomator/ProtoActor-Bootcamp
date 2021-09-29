package main

import (
	"fmt"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	actors "github.com/Mushroomator/ProtoActorGo-Bootcamp/Module_03/Lesson_02/Actors"
	msgs "github.com/Mushroomator/ProtoActorGo-Bootcamp/Module_03/Lesson_02/Msgs"
)

// Main function
func main() {
	// Creating actor system
	system := actor.NewActorSystem()
	fmt.Println("Actor system created")

	// Creating props for the playbackActor
	props := actor.PropsFromProducer(func() actor.Actor {
		return &actors.PlaybackActor{}
	})

	// Creating the playbackActor and receiving its PID which we can use to reference the actor
	pid := system.Root.Spawn(props)

	// Send messages using PID
	system.Root.Send(pid, msgs.NewPlayMovieMessage("The Movie", 44))
	system.Root.Send(pid, msgs.NewPlayMovieMessage("The Movie 2", 54))
	system.Root.Send(pid, msgs.NewPlayMovieMessage("The Movie 3", 64))
	system.Root.Send(pid, msgs.NewPlayMovieMessage("The Movie 4", 74))

	// wait a little while to let the actor process the messages first (not required, but makes output appear in order of execution)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Press [enter] key to restart actor")
	fmt.Scanln()
	// Send Recoverable message to playbackActor which will be forwarded to the first child actor (= childActor) and there cause a panic
	// which in turn causes the parent actor (=playbackActor) to restart the child actor
	system.Root.Send(pid, *msgs.NewRecoverable())

	// wait a little while to let the actor process the message first (not required, but makes output appear in order of execution)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Press [enter] to stop actor")
	fmt.Scanln()
	// stop playbackActor
	system.Root.Stop(pid)

	// Program will terminate immediately after main is run.
	// As all operations are non-blocking this fmt.Scanln() is necessary to prevent the program
	// from terminating before any other actions can be performed
	fmt.Scanln()
}
