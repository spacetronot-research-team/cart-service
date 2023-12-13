package eraerr

import "fmt"

// Wrap return wrapped inner error with outer error, the format is
// outer:: inner.
func Wrap(inner error, outer error) error {
	return fmt.Errorf("%w:: %w", outer, inner)
}
