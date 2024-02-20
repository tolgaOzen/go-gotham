package flags

import (
	"flag"
)

var (
	Production *bool
	Dingo      *bool
	Migrate    *bool
	Seed       *bool
	Server     *bool
)

func init() {
	Production = flag.Bool("production", false, "a bool")
	Dingo = flag.Bool("dingo", false, "a bool")
	Migrate = flag.Bool("migrate", false, "a bool")
	Seed = flag.Bool("seed", false, "a bool")
	Server = flag.Bool("server", false, "a bool")
	flag.Parse()
}
