package main

import (
	"flag"
	"fmt"
	"github.com/transcom/mymove/pkg/edi/gex"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Call this from command line with go run cmd/send_to_gex/main.go -edi <filepath>
func main() {
	ediFile := flag.String("edi", "", "The filepath to an edi file to send to GEX")
	transactionName := flag.String("transactionName", "test", "The required name sent in the url of the gex api request")
	flag.Parse()
	if *ediFile == "" {
		log.Fatal("Usage: go run cmd/send_to_gex/main.go  --edi <edi filepath>")
	}

	file, err := os.Open(*ediFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	edi, err := ioutil.ReadAll(file)

	ediString := string(edi[:])
	// make sure edi ends in new line
	ediString = strings.TrimSpace(ediString) + "\n"

	var logger = zap.NewNop()

	fmt.Println(ediString)
	statusCode, err := gex.SendInvoiceToGex(logger, ediString, *transactionName)

	fmt.Println("Sending to GEX. . .")
	fmt.Printf("status code: %v, error: %v", statusCode, err)
}
