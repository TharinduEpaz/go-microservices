package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"tharinduEpaz/go-microservice/application"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel() // call at the end of the current function

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app : ", err)
	}

}
