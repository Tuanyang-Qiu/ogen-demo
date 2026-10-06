package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/tuanyang-qiu/ogen-demo/internal/handler"
	"github.com/tuanyang-qiu/ogen-demo/internal/handler/ogen"
)

type ComponentGroup struct {
	handler http.Handler
}

func AssemblyAppComponents(logger *slog.Logger) (*ComponentGroup, error) {
	combinedHandler := handler.NewCombinedHandler(logger)
	ogenServer, err := ogen.NewServer(combinedHandler)
	if err != nil {
		return nil, fmt.Errorf("failed to create ogen server: %w", err)
	}
	return &ComponentGroup{handler: ogenServer}, nil
}
