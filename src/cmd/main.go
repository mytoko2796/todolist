package main

import (
	cfg "github.com/mytoko2796/sdk-go/stdlib/config"
	"runtime"
)

var (
	Namespace    string
	GoVersion    string
	BuildVersion string
	BuildTime    string
	CommitHash   string
)


// OverrideConfig will override the static Config as it merges the setting from static config
// and injected build variable. This will be used to update the value in the main config to be displayed
// in the platform HTTP Handler, see (/platform/)
// Swagger Config is updated here as well
func OverrideStaticConfig(staticConf cfg.Conf, conf *Conf) {

	GoVersion = runtime.Version()
	conf.Meta.Go = GoVersion
	staticConf.Set(`meta.Go`, GoVersion)
	staticConf.Set(`log.defaultfields.go`, GoVersion)

	if Namespace != "" {
		conf.Meta.Namespace = Namespace
		staticConf.Set(`meta.Namespace`, Namespace)
		staticConf.Set(`log.defaultfields.namespace`, Namespace)
	}

	if BuildVersion != "" {
		conf.Meta.Version = BuildVersion
		staticConf.Set(`meta.version`, BuildVersion)
		staticConf.Set(`log.defaultfields.version`, BuildVersion)
	}

	if BuildTime != "" {
		conf.Meta.BuildTime = BuildTime
		staticConf.Set(`meta.build`, BuildTime)
		staticConf.Set(`log.defaultfields.build`, BuildTime)
	}

	if CommitHash != "" {
		conf.Meta.CommitHash = CommitHash
		staticConf.Set(`meta.commit`, CommitHash)
		staticConf.Set(`log.defaultfields.commit`, CommitHash)
	}

	conf.Log.DefaultFields = staticConf.GetStringMapString(`log.defaultfields`)
}

