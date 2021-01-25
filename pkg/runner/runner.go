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
func RunAll() (errs []error) {
	for _, f := range allChecks {
		errs = append(errs, f()...)
	}

	return
}

// Run a check for a given name. Returns an error if the lint does not exist in
// `allChecks`.
func Run(name string) []error {
	checkFunc, ok := allChecks[name]
	if !ok {
		return []error{fmt.Errorf("unknown lint - %v", name)}
	}

	return checkFunc()
}
