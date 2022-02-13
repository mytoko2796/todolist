package main

import "flag"

var (
	staticConfigPath string
)

func main()  {
	// get template config file path
	flag.StringVar(&staticConfigPath, "staticConfigPath", "./etc/")
}