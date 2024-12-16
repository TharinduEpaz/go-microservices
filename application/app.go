package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

// by defining the router inside the App struct make this decoupled from chi

type App struct {
	router http.Handler
	rdb    *redis.Client
}

func New() *App {
	app := &App{
		router: loadRoutes(),
		rdb:    redis.NewClient(&redis.Options{}),
	}
	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	// connet to the database
	rdb := redis.NewClient(&redis.Options{
        Addr: "54.86.44.104:6379", // Replace with your EC2 public IP address
        Password: "", // No password set
        DB: 0,        // use default DB
    })

    pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println(pong)

	// close the database at the end of the start function
	defer func() {
		if err := a.rdb.Close(); err != nil {
			fmt.Println("failed to close redis", err)
		}
	}()

	fmt.Println("Starting server")

	ch := make(chan error, 1) // second arguement is the buffer size

	// run the server in a separate go routine
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server : %w", err)
		}
		close(ch)
	}()

	// This waits for either:
	// An error from the server (through the channel)
	// A signal to shut down (through the context)
	// If shutdown is requested,
	// it gives the server 10 seconds to finish any ongoing requests before forcing a shutdown.
	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		fmt.Println("Shutting down the server")
		return server.Shutdown(timeout)
	}
}
