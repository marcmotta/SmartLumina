// cmd/smartlumina/main.go
package main

import (
	"flag"
	"log"
	"os"

	"smartlumina/internal/smartlumina"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := smartlumina.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
