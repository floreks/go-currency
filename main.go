// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
