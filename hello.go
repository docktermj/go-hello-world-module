package hello

import (
	"fmt"
)

// Values updated via "go install -ldflags" parameters.

var programName string = "go-hello-world-module"
var buildVersion string = "1.0.0"
var buildIteration string = "0"
var helloName string = "world"

func Hello() string {
	return fmt.Sprintf("Hello, %s! from %s version %s-%s", helloName, programName, buildVersion, buildIteration)
}
