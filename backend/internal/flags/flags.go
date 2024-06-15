package flags

import (
	"flag"
	"fmt"
	"os"
)

type Options struct {
	Config string
}

func usage() {
	fmt.Print("--config, to print config")
}

var help = flag.Bool("help", false, "show help")
var config = flag.String("config", "", "config file path")

func init() {
	flag.StringVar(config, "c", "", "config file path")
}

func Parse() Options {
	flag.Usage = usage
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	return Options{
		Config: *config,
	}
}
