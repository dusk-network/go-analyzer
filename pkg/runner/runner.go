// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT License was not distributed with this
// file, you can obtain one at https://opensource.org/licenses/MIT.
//
// Copyright (c) DUSK NETWORK. All rights reserved.

// Package runner provides an easy interface to run any available lint, as well
// as running them all in sequence.
package runner

import (
	"fmt"

	"github.com/dusk-network/go-analyzer/pkg/lints"
)

// A list of all checks defined in the package.
var allChecks = map[string]lints.Fn{
	"license": lints.CheckLicenseHeaders,
}

// RunAll checks in the `lints` package.
func RunAll(dirPath string) (errs []error) {
	for _, f := range allChecks {
		errs = append(errs, f(dirPath)...)
	}

	return
}

// Run a check for a given name. Returns an error if the lint does not exist in
// `allChecks`.
func Run(name, dirPath string) []error {
	checkFunc, ok := allChecks[name]
	if !ok {
		return []error{fmt.Errorf("unknown lint - %v", name)}
	}

	return checkFunc(dirPath)
}
