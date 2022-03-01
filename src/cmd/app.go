package main

import (
	"context"
	"crypto/tls"
	"flag"
	sqlx "github.com/mytoko2796/sdk-go/stdlib/sql"
	"runtime"

	errors "github.com/mytoko2796/sdk-go/stdlib/error"
	"github.com/mytoko2796/sdk-go/stdlib/httpmux"
	"github.com/mytoko2796/sdk-go/stdlib/httpserver"
	//"flag"
	"github.com/mytoko2796/sdk-go/stdlib/parser"
	"github.com/mytoko2796/todolist/src/business/domain"
	"github.com/mytoko2796/todolist/src/business/usecase"
	resthandler "github.com/mytoko2796/todolist/src/handler/rest"

	cfg "github.com/mytoko2796/sdk-go/stdlib/config"
	"github.com/mytoko2796/sdk-go/stdlib/grace"
	"github.com/mytoko2796/sdk-go/stdlib/health"
	"github.com/mytoko2796/sdk-go/stdlib/httpmiddleware"
	log "github.com/mytoko2796/sdk-go/stdlib/logger"
	"github.com/mytoko2796/sdk-go/stdlib/telemetry"
)

var (
	staticConfPath              string
	conf Conf

	// Resource - Storage
	sqlClient0   sqlx.SQL

	// Server Infrastructure
	logger     log.Logger
	staticConf cfg.Conf
	remoteConf cfg.Conf
	secretConf cfg.Conf
	tele       telemetry.Telemetry
	healt      health.Health
	parse      parser.Parser
	httpMware  httpmiddleware.HttpMiddleware
	httpMux    httpmux.HttpMux
	httpServer httpserver.HTTPServer
	app        grace.App

	//business
	uc *usecase.Usecase
	dom *domain.Domain
)

func init()  {
	//Init flag setting
	//Init static configuration path file
	flag.StringVar(&staticConfPath, "staticConfPath", "./etc/conf/development.yaml", "config path")

	//Parse all flag
	flag.Parse()

	logger := log.Init(log.Options{})

	//read config
	staticConf = cfg.Init(logger, cfg.AppStaticConfig, cfg.Options{
		Enabled:         true,
		Type:            "yaml",
		Path:            staticConfPath,
		RestartOnChange: true,
	})
	staticConf.ReadAndWatch(&conf)

	// Override All Config : Static
	// Swagger config is initialized
	// Injected build variable is initialized
	OverrideStaticConfig(staticConf, &conf)

	// Override Logger Settings with new Logger Settings
	// Logger Default Fields is initialized
	logger.SetOptions(conf.Log)


	//init infra dependencies
	parse = parser.Init(conf.Parser)

	sqlClient0 = sqlx.Init(logger, conf.SQL["sql-0"])

	// Register Liveness HealthCheck Function
	// You have to adjust your liveness check function based on your application behaviour
	// Add LivenessCheck to check the liveness as internal process monitor
	// number of goroutines, etc
	// Context will be initialized with timeout in the static config - health.liveness.CheckTimeout
	// Cancel Func must be called as deferred function to prevent context leak
	conf.Health.Liveness.CheckF = func(ctx context.Context, cancel context.CancelFunc) error {
		defer cancel()
		numgoroutine := runtime.NumGoroutine()
		if numgoroutine > 3000 {
			err := errors.New(`too many process`)
			logger.Error(err)
			return err
		}
		return nil
	}

	// Register Readiness HealthCheck Function
	// Add ReadinessCheck to SQL Database - Ping to db leader Or PingF to db followers
	// Add ReadinessCheck to Redis - Ping to redis
	// Add ReadinessCheck to check the liveness to another dependencies/ services if applicable
	// Context will be initialized with timeout in the static config - health.readiness.CheckTimeout
	// Cancel Func must be called as deferred function to prevent context leak
	conf.Health.Readiness.CheckF = func(ctx context.Context, cancel context.CancelFunc) error {
		defer cancel()
		//if sqlClient0 != nil {
		//	if err := sqlClient0.Leader().Ping(ctx); err != nil {
		//		logger.Error(err)
		//		return err
		//	}
		//}

		return nil
	}

	// Healthcheck Init
	// Healthcheck will be blocking the app until the first readiness and liveness status is set
	// This behavior is intended to prevent this app to serve any kinds of TCP services so
	// The application grace upgrader won't switch before apps is ready
	// This behavior is expected to handle TCP traffic hiccup during application upgrade
	// See options : Health.waitBeforeContinue to disable this behavior
	healt = health.Init(logger, conf.Health)

	// Telemetry with opencensus - spawning telemetry server based on config
	// Telemetry spawns 2 server if enabled :
	// 	Stats Server ( only if you enabled Zpage Or Prometheus in the options)
	//		see HTTP Address& Port Configuration at telemetry.Exporters.Stats
	// 	Pprof Server
	//		see HTTP Address& Port Configuration at telemetry.Exporters.Profiler
	tele = telemetry.Init(logger, conf.Telemetry)

	// Middleware Init
	// Default Middleware :
	//	CatchPanicAndReport - Catching Panic from any HTTP Handler and Report this to Prometheus
	// 	HealthCheck - Block any incoming traffic if the readiness and liveness check fails above threshold
	// 	RequestDump - Dump and log all incoming HTTP Requests
	// 	Secure : Block any incoming HTTP Requests that does not satisfy the Security Rules
	//		Security settings can be found in Static Options : security
	httpMware = httpmiddleware.Init(logger, healt, conf.HTTPMiddleware)

	httpMux = httpmux.Init(httpmux.HTTPROUTER, logger, staticConf, remoteConf, httpMware, tele, healt, conf.HTTPMux)

	//init business dependencies
	dom = domain.Init(logger, sqlClient0)
	uc = usecase.Init(dom)

	// REST Handler Initialization
	_ = resthandler.Init(logger,parse, uc, httpMux)

	// HTTPServer TLS Config
	// Only Will be initialized if HTTPS Server is enabled
	conf.HTTPServer.TLSConfig = &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	httpServer = httpserver.Init(logger, tele, httpMux, conf.HTTPServer)

	// Grace Upgrader Init
	// Graceapp listens to SIGHUP Signal
	// It waits the new App to be ready before switch the traffic
	app = grace.Init(logger, tele, httpServer, conf.GraceApp)
}

func main()  {
	defer func() {
		if healt != nil {
			healt.Stop()
		}
		if sqlClient0 != nil {
			sqlClient0.Stop()
		}
		if staticConf != nil {
			staticConf.Stop()
		}
		if remoteConf != nil {
			remoteConf.Stop()
		}
		if secretConf != nil {
			secretConf.Stop()
		}
		if logger != nil {
			logger.Stop()
		}
		app.Stop()
	}()

	app.Serve()
}