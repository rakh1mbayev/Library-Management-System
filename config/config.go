package config

import "flag"

var (
	Port = flag.String("port", "8080", "The port to listen on")
)
