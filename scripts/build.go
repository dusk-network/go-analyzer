// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT License was not distributed with this
// file, you can obtain one at https://opensource.org/licenses/MIT.
//
// Copyright (c) DUSK NETWORK. All rights reserved.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	// GOLANGCI_VERSION to be used for linting.
	GOLANGCI_VERSION = "github.com/golangci/golangci-lint/cmd/golangci-lint@v1.31.0"
)

// GOBIN environment variable.
func GOBIN() string {
	if os.Getenv("GOBIN") == "" {
		log.Fatal("GOBIN not set")
	}

	return os.Getenv("GOBIN")
}

func main() {
	if _, err := os.Stat(filepath.Join("scripts", "build.go")); os.IsNotExist(err) {
		log.Fatal("should run build from root dir")
	}

	if len(os.Args) < 2 {
		log.Fatal("cmd required, eg: install")
	}

	switch os.Args[1] {
	case "lint":
		lint()
	default:
		log.Fatal("cmd not found: ", os.Args[1])
	}
}

//nolint:gosec
func lint() {
	v := flag.Bool("v", false, "log verbosely")

	// Make sure GOLANGCI is downloaded and available.
	argsGet := []string{"get", GOLANGCI_VERSION}
	cmd := exec.Command(filepath.Join(runtime.GOROOT(), "bin", "go"), argsGet...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("could not list pkgs: %v\n%s", err, string(out))
	}

	cmd = exec.Command(filepath.Join(GOBIN(), "golangci-lint"))
	cmd.Args = append(cmd.Args, "run", "--config", ".golangci.yml")

	if *v {
		cmd.Args = append(cmd.Args, "-v")
	}

	fmt.Println("Linting...", strings.Join(cmd.Args, " \\\n"))
	cmd.Stderr, cmd.Stdout = os.Stderr, os.Stdout

	if err := cmd.Run(); err != nil {
		log.Fatal("Error: Could not Lint... ", "error: ", err, ", cmd: ", cmd)
	}
}
