package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/emicklei/go-restful"
	"github.com/floreks/go-currency/service/converter"
	"github.com/spf13/pflag"
)

var (
	argPort = pflag.Int("port", 8080, "The port to listen on for incoming HTTP requests")
)

func main() {
	// Set logging out to standard console out
	log.SetOutput(os.Stdout)

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	// Register handler
	restful.Add(converter.NewConverterService().Handler())

	log.Printf("Listening on port: %d", *argPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *argPort), nil))
}
