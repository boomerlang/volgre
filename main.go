//
// main package
// Author Bogdan Peta
//
package main

import (
	// "os"
	"fmt"
	"log"
	"time"
	"net/http"
	"errors"
	"os"
	"context"
	"os/signal"
	"syscall"
	"flag"
	"github.com/gorilla/mux"

	"github.com/boomerlang/volgre/controllers"
)

var (
    host string
    port string
)

func main () {
	
	fmt.Println()

	flag.StringVar(&host, "host", "10.10.11.66", "host.domain.tld")
	flag.StringVar(&port, "port", "8082", "8080")
	flag.Parse()
	
	controllers.PreloadRuleEngines()

	log.Println("Server is starting...")

	r := mux.NewRouter()
	
	r.HandleFunc("/run/engine/{rule_engine}", controllers.RunRuleEngineHandler).Methods("POST")
	r.HandleFunc("/refresh/engine/{rule_engine}", controllers.RefreshRuleEngineHandler).Methods("GET")
	r.HandleFunc("/version/engine/{rule_engine}", controllers.VersionRuleEngineHandler).Methods("GET")
	
	srv := &http.Server{
		Handler: r, 
		Addr: host + ":" + port,
		WriteTimeout: 15 * time.Second, 
		ReadTimeout: 15 * time.Second, 
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("Listen: %s\n", err)
		}
	}()
	
	shutdown_channel := make(chan os.Signal)
	signal.Notify(shutdown_channel, syscall.SIGINT, syscall.SIGTERM)
	s := <-shutdown_channel
	
	log.Println("Server received signal:", s)
	log.Println("Shutting down server...")
	
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	
	log.Println("Server exiting")
}
