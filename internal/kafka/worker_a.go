package kafka

import (
	"context"
	"crypto/tls"
	"fmt"
	"go-boilerplate/configs"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
)

type WorkerA interface {
	Start()
}

type workerA struct {
	cli       *kgo.Client
	ibItmxcfg configs.WorkerAConsumerCfg
	cliId     string
}

func (wa *workerA) Start() {
	go wa.consume()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	wa.cli.Close()
}

func (wa *workerA) consume() {
	for {
		fetches := wa.cli.PollFetches(context.Background())
		if fetches.IsClientClosed() {
			return
		}

		fetches.EachError(func(t string, p int32, err error) {
			log.Panicf("%s: fetch err on topic %s partition %d: %v", "Errors", t, p, err)
		})

		fetches.EachRecord(func(r *kgo.Record) {
			// Process the record
		})
	}
}

func NewWorkerA(ibItmxcfg configs.WorkerAConsumerCfg) WorkerA {
	cliId := fmt.Sprintf("worker-a-%s-%d", ibItmxcfg.WokerAGroupId, os.Getpid())
	broker := strings.Split(ibItmxcfg.Host, ",")

	timeOut, err := time.ParseDuration(ibItmxcfg.WorkerATimeOut)
	if err != nil {
		log.Fatalf("Error parsing timeout: %v", err)
	}

	opts := []kgo.Opt{
		kgo.SeedBrokers(broker...),
		kgo.ClientID(cliId),
		kgo.ConsumerGroup(ibItmxcfg.WokerAGroupId),
		kgo.ConsumeTopics(ibItmxcfg.WorkerATopic),
		kgo.Dialer((&tls.Dialer{NetDialer: &net.Dialer{Timeout: timeOut}}).DialContext),
		kgo.WithLogger(kgo.BasicLogger(os.Stderr, kgo.LogLevelInfo, nil)),
	}

	cli, err := kgo.NewClient(opts...)
	if err != nil {
		log.Panicf("Failed to create client: %v", err)
	}

	return &workerA{
		cli:       cli,
		ibItmxcfg: ibItmxcfg,
		cliId:     cliId,
	}
}
