package listeners

import (
	"github.com/nats-io/nats.go"
	"github.com/saiprasadkrishnamurthy/web-api/config"
)

// BaseListener base listener.
type BaseListener struct {
	NatsConn *nats.Conn
}

// OnMessage on message function of all the listeners.
func (l *BaseListener) OnMessage(subject string, queue string, msgHandler nats.MsgHandler) {
	l.NatsConn.QueueSubscribe(subject, queue, msgHandler)
}

// InitializeAllListeners initializes all listeners.
func InitializeAllListeners(nats *nats.Conn) {
	baseListener := &BaseListener{NatsConn: config.GetNats()}
	initializeFooListener(baseListener)

	// List all your listeners here.

}

func initializeFooListener(baseListener *BaseListener) {
	foo := &FooListener{BaseListener: baseListener}
	foo.OnMessage("foo", "queue", foo.Handle)
}
