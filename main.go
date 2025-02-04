package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

func main() {
	h := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(h)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		headers := []any{}
		for key, vals := range r.Header {
			fmt.Println("key:", key)
			headers = append(headers, slog.String(key, strings.Join(vals, ";")))
		}

		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Error("failed to read body", slog.Any("err", err))
		}

		logger.InfoContext(ctx,
			"got request",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.String("proto", r.Proto),
			slog.Int64("contentLength", r.ContentLength),
			slog.String("host", r.Host),
			slog.String("remoteAddr", r.RemoteAddr),
			slog.Group("headers", headers...),
			slog.String("body", string(bytes)),
		)
	})

	port := os.Getenv("PORT")
	if port == "" {
		logger.Error("PORT cannot be blank")
		os.Exit(1)
	}
	logger.Info("server listening", slog.String("port", port))
	if err := http.ListenAndServe(":"+port, nil); err != nil && err != http.ErrServerClosed {
		logger.Error("server quit unexpectedly", slog.Any("err", err))
	}
}
