package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	kitlog "github.com/go-kit/kit/log"
	"github.com/spiffcp/borkbot"
)

func main() {
	var (
		httpAddr          = flag.String("listen", ":9000", "HTTP listen and serve address for service")
		verificationToken = flag.String("verification_token", "", "Slack token used to verify requests come from slack")
	)
	flag.Parse()
	errChan := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	service := borkbot.NewService(*verificationToken)
	w := kitlog.NewSyncWriter(os.Stderr)
	logger := kitlog.NewLogfmtLogger(w)
	logService := borkbot.NewLoggingService(logger, service)
	endpoints := borkbot.MakeEndpoints(logService)
	go func() {
		log.Println("http:", *httpAddr)
		handler := borkbot.NewHTTPServer(endpoints, logger)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()
	// Block from exiting unless error is recieved
	log.Fatalln(<-errChan)
}
