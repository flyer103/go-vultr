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
	APIServerReinstall = fmt.Sprintf("%s/v1/server/reinstall", APIPrefix)
	APIServerHalt      = fmt.Sprintf("%s/v1/server/halt", APIPrefix)
	APIServerStart     = fmt.Sprintf("%s/v1/server/start", APIPrefix)
	APIServerReboot    = fmt.Sprintf("%s/v1/server/reboot", APIPrefix)
)
