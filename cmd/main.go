package main

import (
	"context"
	broker "github.com/xnok/tasq-renovate/broker"
	"github.com/xnok/tasq-renovate/tasks"
	"log"
	"log/slog"
)

func main() {
	ctx := context.Background()

	broker.MainQueue.Consumer().Stop()

	if err := broker.MainQueue.Consumer().Start(ctx); err != nil {
		log.Fatal(err)
	}

	err := broker.MainQueue.Add(tasks.DiscoverTask.WithArgs(ctx))
	if err != nil {
		slog.Error("failed to add initial discover task to queue")
	}

	sig := broker.WaitSignal()
	log.Println(sig.String())

	err = broker.QueueFactory.Close()
	if err != nil {
		log.Fatal(err)
	}
}
