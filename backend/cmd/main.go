package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sbsysdev/go-svelte-template/internal/infrastructure"
)

func main() {
	ctx := context.Background()
	// Load environment variables
	env := infrastructure.NewEnvironment()
	// Load db connection pool
	storage := infrastructure.NewStorage(ctx, env)
	// Load api server
	api := infrastructure.NewApiServer(env, storage)
	// Start the server
	if err := api.StartApiServer(); err != nil {
		fmt.Printf("Error starting the server: %v\n", err)
		os.Exit(1)
	}
}
