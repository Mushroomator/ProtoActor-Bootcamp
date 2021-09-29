package actors

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	msgs "github.com/Mushroomator/ProtoActorGo-Bootcamp/Module_03/Lesson_02/Msgs"
	consolepainter "github.com/Mushroomator/ProtoActorGo-Bootcamp/Module_03/consolePainter"
)

type ChildActor struct{}

// Process system message *actor.Started
func (state *ChildActor) ProcessRecoverableMessage(msg msgs.RecoverableMessage) {
	panic("Oops! An exception/ panic occured.")
}

// Process system message *actor.Restarting
func (state *ChildActor) ProcessRestartingMessage(msg *actor.Restarting) {
	cp := consolepainter.NewConsolePainter()
	cp.PrintfGreen("ChildActor restarting\n")
}

func (state *ChildActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Restarting:
		state.ProcessRestartingMessage(msg)
	case msgs.RecoverableMessage:
		state.ProcessRecoverableMessage(msg)
	}

}

// Create new ChildActor
func NewChildActor() *ChildActor {
	return &ChildActor{}
}
