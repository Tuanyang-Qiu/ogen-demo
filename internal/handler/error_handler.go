package handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/tuanyang-qiu/ogen-demo/internal/handler/ogen"
)

// NewErrorHandler returns the handler ogen calls for every error that does not
// come back as a typed response: request decode/validation failures and any
// error a handler returns. The body uses the spec's Error schema.
func NewErrorHandler(logger *slog.Logger) ogen.ErrorHandler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
		code := ogenerrors.ErrorCode(err)
		message := err.Error()
		// A 4xx is the caller's mistake: tell them what to fix.
		// A 5xx is ours: log the real error, return only the status text,
		// because err.Error() can carry internal hostnames.
		if code >= http.StatusInternalServerError {
			logger.ErrorContext(ctx, "request failed",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Int("status", code),
				slog.Any("error", err),
			)
			message = http.StatusText(code)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		_ = json.NewEncoder(w).Encode(&ogen.Error{Message: message})
	}
}
