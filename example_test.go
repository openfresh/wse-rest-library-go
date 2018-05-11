package wserest_test

import (
	"fmt"
	"log"
	"os"

	wserest "github.com/openfresh/wse-rest-library-go"
	"github.com/openfresh/wse-rest-library-go/entity/application/helper"
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

func Example() {
	// It is simple to create a setup object for transporting our settings
	setup := helper.NewDefaultSettings()
	setup.SetHost(WowzaHost)
	setup.SetUsername(WowzaUsername)
	setup.SetPassword(WowzaPassword)

	// Connect to the server or deal with statistics NOTICE THE CAPS IN COM AND WOWZA
	server := wserest.NewServer(setup)
	sf := wserest.NewStatistics(setup)

	if UseWowza == "" {
		fmt.Println(2)
		return
	}

	response, err := sf.GetServerStatistics(server)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(response))
	// Output: 2
}
