// Server
package main

import (
	"log"

	"github.com/srg-bnd/keeper/internal/server/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Panicln(err)
	}
}
