package brokercore

import (
	"github.com/aliworkshop/errorslib"
)

type ActorFunc func(message Message) errorslib.ErrorModel

type Consumer interface {
	Subscribe(topics []string) errorslib.ErrorModel
	GetChan() chan Message
	Commit(message Message) errorslib.ErrorModel
	StartListening()
}

type Message interface {
	GetData() []byte
	SetMetadata(metadata Metadata)
	GetMetadata() Metadata
	GetTopic() string
	GetPartition() int32
	GetOffset() int64
}

type Producer interface {
	GetChan() chan Message
	StartProducing()
	SetFailActor(actorFunc ActorFunc)
	Stop()
}
