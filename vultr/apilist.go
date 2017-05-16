package vultr

import (
	"fmt"
)

const (
	APIPrefix = "https://api.vultr.com"
)

// Vultr API List
var (
	// Server
	APIServerList      = fmt.Sprintf("%s/v1/server/list", APIPrefix)
	APIServerReInstall = fmt.Sprintf("%s/v1/server/reinstall", APIPrefix)
)
