// Package lints defines all available lints for go-analyzer.

package lints

// Fn defines the standard function signature for a lint.
type Fn func() []error
