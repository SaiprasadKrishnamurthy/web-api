package listeners

import (
	"log"

	"github.com/nats-io/nats.go"
)

// FooListener foo listener.
type FooListener struct {
	*BaseListener
}

// Handle on message function of all the listeners.
func (l *FooListener) Handle(msg *nats.Msg) {
	log.Println(" Handle Message Received in Foo Listener...", string(msg.Data))
}
