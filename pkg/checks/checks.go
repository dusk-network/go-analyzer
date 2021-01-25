package checks

import "fmt"

// The signature of a top-level linter function.
type checkFunc func() []error

// A list of all checks defined in the package.
var allChecks = map[string]checkFunc{}

// RunAll checks in the package.
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
