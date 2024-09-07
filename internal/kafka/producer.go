package kafka

import (
	"context"
	"crypto/tls"
	"fmt"
	"go-boilerplate/configs"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Producer interface {
	SendMessage(
		ctx context.Context,
		topic string,
		key string,
		payload []byte,
	) (err error)
}

type producer struct {
	cli   *kgo.Client
	pdCfg configs.ProducerCfg
	cliId string
}

func (p *producer) SendMessage(
	ctx context.Context,
	topic string,
	key string,
	payload []byte,
) (err error) {
	record := &kgo.Record{
		Topic: topic,
		Key:   []byte(key),
		Value: payload,
	}

	if err := p.cli.ProduceSync(ctx, record).FirstErr(); err != nil {
		return err
	}
	return nil
}

func NewProducer(
	pdCfg configs.ProducerCfg,
) Producer {
	cliId := fmt.Sprintf("producer-%d", os.Getpid())
	brokers := strings.Split(pdCfg.Host, ",")

	// TODO add more options for production
	opts := []kgo.Opt{
		kgo.SeedBrokers(brokers...),
		kgo.ClientID(cliId),
		kgo.Dialer((&tls.Dialer{NetDialer: &net.Dialer{Timeout: 10 * time.Second}}).DialContext),
	}

	cl, err := kgo.NewClient(opts...)
	if err != nil {
		log.Panicf("Error to create kafka client: %v", err)
	}

	return &producer{
		cli:   cl,
		pdCfg: pdCfg,
		cliId: cliId,
	}
}
