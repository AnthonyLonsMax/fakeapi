package server

import (
	"io"
	"log/slog"
)

func AddLogger(w io.Writer) {
	handler := slog.NewJSONHandler(w, &slog.HandlerOptions{})
	slog.SetDefault(slog.New(handler))
}
