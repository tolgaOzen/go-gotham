package flags

import (
	"flag"
)

var (
	Production *bool
	Migrate    *bool
	Seed       *bool
)

func init() {
	Production = flag.Bool("production", false, "a bool")
	Migrate = flag.Bool("migrate", false, "a bool")
	Seed = flag.Bool("seed", false, "a bool")
	flag.Parse()
}
