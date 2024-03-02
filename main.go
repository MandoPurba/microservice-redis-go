package main

import (
	"context"
	"fmt"
	"microservice-go/apps"
)

func main() {
	app := apps.New()
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app: ", err)
	}
}
