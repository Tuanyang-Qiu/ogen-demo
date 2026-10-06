package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	componentGroup, err := AssemblyAppComponents(logger)
	if err != nil {
		log.Fatalf("failed to assembly components due to %v", err)
	}

	logger.Info("listening", slog.String("addr", "http://localhost:8080"))
	if err = http.ListenAndServe(":8080", componentGroup.handler); err != nil {
		log.Fatalf("failed to run application due to %v", err)
	}
}
