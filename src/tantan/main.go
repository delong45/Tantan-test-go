package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"tantan/config"
	"tantan/g"
	"tantan/handler"
	"tantan/service"
	"tantan/user"
)

const (
	ErrnoOK int = iota
	ErrnoConfigParseFailed
	ErrnoUserInitFailed
	ErrnoHTTPStartFailed
)

func main() {
	cfg := flag.String("c", "config.json", "configuration file")
	version := flag.Bool("v", false, "show version")

	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(ErrnoOK)
	}

	conf, err := config.Parse(*cfg)
	if err != nil {
		log.Printf("configuration file parse error: %v\n", err)
		os.Exit(ErrnoConfigParseFailed)
	}

	err = user.Init()
	if err != nil {
		log.Printf("user init error: %v\n", err)
		os.Exit(ErrnoUserInitFailed)
	}

	s := &service.Server{
		Conf:     conf,
		Closing:  make(chan struct{}),
		HTTPDone: make(chan struct{}),
		Handler:  handler.Init(),
	}

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
		<-signals
		log.Println("Initializing shutdown...")
		close(s.Closing)
	}()

	err = service.Run(s)
	if err != nil {
		log.Printf("failed to start HTTP Server: %v\n", err)
		os.Exit(ErrnoHTTPStartFailed)
	}

	<-s.HTTPDone
	log.Println("Tantan service Exit")
}
