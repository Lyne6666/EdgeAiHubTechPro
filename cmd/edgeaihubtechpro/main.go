// cmd/edgeaihubtechpro/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"edgeaihubtechpro/internal/edgeaihubtechpro"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := edgeaihubtechpro.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
