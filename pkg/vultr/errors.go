package vultr

import (
	"errors"
)

// General errrors
var (
	ErrNoAPIKey = errors.New("No API Key")
	ErrNoDCID   = errors.New("No DCID")
)
