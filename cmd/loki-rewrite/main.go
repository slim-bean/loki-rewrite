package main

import (
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/slim-bean/loki-rewrite/pkg/loki"
	"os"
	"time"
)

func main() {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowInfo())
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.Caller(3))

	l := loki.New(logger)

	query := `{job="screencap"} | logfmt | label_format thumb="" | line_format "ts={{.ts}} type={{.type}} loc={{.loc}}"`

	from, err := time.Parse(time.RFC3339, "2023-01-01T00:00:00Z")
	if err != nil {
		panic(err)
	}

	to, err := time.Parse(time.RFC3339, "2023-02-05T00:00:00Z")
	if err != nil {
		panic(err)
	}

	l.Process(query, from, to, 1000)
}
