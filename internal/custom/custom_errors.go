package custom

import (
	"flag"
	"fmt"
	"os"
)

func NoProjectArgs() {
	fmt.Fprintf(flag.CommandLine.Output(),
		"Project name need to have at least 2 chars, passing: %s\n", os.Args[1])
}
