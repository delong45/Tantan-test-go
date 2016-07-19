package service

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/facebookgo/httpdown"
)

func Run(sev *Server) error {
	conf := sev.Conf
	addr := conf.HTTPAddr
	log.Printf("Starting HTTP Server: %s\n", addr)

	s := &http.Server{
		Addr:           addr,
		Handler:        sev.Handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	hd := &httpdown.HTTP{}
	hs, err := hd.ListenAndServe(s)
	if err != nil {
		return fmt.Errorf("failed to listen and serve: %v", err)
	}

	go func(hs httpdown.Server) {
		waiterr := make(chan error, 1)
		go func() {
			defer close(waiterr)
			waiterr <- hs.Wait()
		}()

		select {
		case err := <-waiterr:
			if err != nil {
				log.Printf("failed to wait HTTP Server: %v", err)
			}
		case <-sev.Closing:
			err := hs.Stop()
			if err != nil {
				log.Printf("failed to stop HTTP Server: %v", err)
			}
		}

		log.Printf("HTTP Server shutdown\n")
		sev.HTTPDone <- struct{}{}
	}(hs)

	return nil
}
