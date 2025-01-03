package config

import (
	"flag"
	"github.com/sirupsen/logrus"
)

var (
	Port   = flag.String("port", "8080", "The port to listen on")
	Logger = lorgus.New()
)
