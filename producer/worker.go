package producer

import "github.com/twmb/franz-go/pkg/kgo"

type WorkerProducer interface {
	ProduceMesage(key string, payload []byte) (err error)
}

type workerProducer struct {
	cl   *kgo.Client
	clId string
}
