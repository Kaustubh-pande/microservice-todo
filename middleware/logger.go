package middleware

import (
	"context"
	"os"

	"github.com/micro/go-micro/server"
	log "github.com/sirupsen/logrus"
)

// Setup initilize logrus logging
// You can also use text formatter instead of json
// log.SetFormatter(&log.TextFormatter{
// 	DisableColors: true,
// 	FullTimestamp: true,
// })
func Setup() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	// See different Log level https://github.com/sirupsen/logrus#level-logging
	log.SetLevel(log.WarnLevel)

	// Set this to true If you wish to add the calling method as a field
	log.SetReportCaller(false)
}

// Logger implements the server.HandlerWrapper
func Logger(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		err := fn(ctx, req, resp)
		if err != nil {
			log.Error(err)
		}
		return err
	}
}
