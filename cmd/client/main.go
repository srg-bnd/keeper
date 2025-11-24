// Client for keeper-server
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/srg-bnd/keeper/internal/client/bootstrap"
)

// handleSignals sets up OS signal handling and cancels the context on interrupt/termination.
func handleSignals(cancel context.CancelFunc) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-signalChan
		log.Println("Hey!")

		log.Printf("Received signal: %v. Shutting down...", sig)
		cancel()
	}()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	handleSignals(cancel)

	if err := bootstrap.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
