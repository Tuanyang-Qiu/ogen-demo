package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/tuanyang-qiu/ogen-demo/internal/handler"
	"github.com/tuanyang-qiu/ogen-demo/internal/handler/ogen"
	productRepo "github.com/tuanyang-qiu/ogen-demo/internal/repo/product"
	productSvc "github.com/tuanyang-qiu/ogen-demo/internal/service/product"
)

type ComponentGroup struct {
	handler http.Handler
}

func AssemblyAppComponents(logger *slog.Logger) (*ComponentGroup, error) {
	prodRepo := productRepo.NewMemoryRepository()
	prodSvc := productSvc.NewService(prodRepo)

	combinedHandler := handler.NewCombinedHandler(logger, prodSvc)
	ogenServer, err := ogen.NewServer(combinedHandler,
		ogen.WithErrorHandler(handler.NewErrorHandler(logger)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create ogen server: %w", err)
	}
	return &ComponentGroup{handler: ogenServer}, nil
}
