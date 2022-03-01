package main

import (
	"github.com/mytoko2796/sdk-go/stdlib/grace"
	"github.com/mytoko2796/sdk-go/stdlib/health"
	"github.com/mytoko2796/sdk-go/stdlib/httpmiddleware"
	"github.com/mytoko2796/sdk-go/stdlib/httpmux"
	"github.com/mytoko2796/sdk-go/stdlib/httpserver"
	log "github.com/mytoko2796/sdk-go/stdlib/logger"
	"github.com/mytoko2796/sdk-go/stdlib/parser"
	sqlx "github.com/mytoko2796/sdk-go/stdlib/sql"
	"github.com/mytoko2796/sdk-go/stdlib/telemetry"
)

type Options struct {
	Group       string
	Namespace   string
	Version     string
	Description string
	Host        string
	BasePath    string
	Go          string
	BuildTime   string
	CommitHash  string
}

type Conf struct {
	Meta Options
	Parser parser.Options
	Log            log.Options

	SQL     map[string]sqlx.Options

	HTTPMiddleware httpmiddleware.Options

	// Server Infrastructure Options
	Telemetry      telemetry.Options
	Health         health.Options
	HTTPMux        httpmux.Options
	GraceApp       grace.Options
	HTTPServer     httpserver.Options

}