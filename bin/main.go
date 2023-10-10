package main

import (
	"fmt"
	"os"
)

func main() {
	envVariableName := "TEST_ENV_VARIABLE"

	// Read the environment variable
	envValue := os.Getenv(envVariableName)

	if envValue == "" {
		fmt.Printf("The environment variable '%s' is not set.\n", envVariableName)
	} else {
		fmt.Printf("The value of the environment variable '%s' is: %s\n", envVariableName, envValue)
	}
}
