package wserest

import (
	"os"
)

var UseWowza = getenv("USE_WOWZA", "")
var WowzaHost = getenv("WOWZA_HOST", "http://111.111.123.123:8087/v2")
var WowzaServerInstance = getenv("WOWZA_SERVER_INSTANCE", "_defaultServer_")
var WowzaVhostInstance = getenv("WOWZA_VHOST_INSTANCE", "_defaultVHost_")

var WowzaUsername = getenv("WOWZA_USERNAME", "my_secret_username")
var WowzaPassword = getenv("WOWZA_PASSWORD", "my_super_cool_password")

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
