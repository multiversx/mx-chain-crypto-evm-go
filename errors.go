package crypto

import (
	"errors"
)

// ErrImplementationNotDefinedForSelector is returned when an implementation was not defined for a selector
var ErrImplementationNotDefinedForSelector = errors.New("implementation not defined for selector")
