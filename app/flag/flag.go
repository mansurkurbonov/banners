package flag

import (
	"flag"
	"log"
	"os"
)

// configDefault default config file to parse.
const configDefault = "dev.json"

// f package instance of _Flags.
var f flags

// Peek provides secure access to flag options.
func Peek() flags {
	return f
}

// init parses program run flags.
func init() {
	flag.BoolVar(&f.Help, "help", false, "show this help")
	flag.StringVar(&f.ConfigFile, "config", configDefault, "name of config json file located in dir ./configs")

	flag.Usage = func() {
		flag.PrintDefaults()
	}

	flag.Parse()
	if !flag.Parsed() {
		log.Panicln("flag: failed to parse command flags")
		return
	}

	if f.Help {
		flag.Usage()
		os.Exit(0)
	}
	return
}

// flags is a structure which holds application command line options.
type flags struct {
	Help       bool
	ConfigFile string
}
