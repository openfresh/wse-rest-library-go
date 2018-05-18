# Go REST Library for Wowza Streaming Engine
Wowza Streaming Engine [media server software](https://www.wowza.com/products/streaming-engine) includes a REST API that you can wrap with a Go library to configure, manage, and monitor your streaming media server through Go requests.

## Prerequisites
Wowza Streaming Engine™ 4.0.0 or later is required.

### Install

Go get it:

	$ go get github.com/openfresh/wse-rest-library-go
	
Then import it:

	import wserest "github.com/openfresh/wse-rest-library-go"

## Example

	// It is simple to create a setup object for transporting our settings
	setup := helper.NewDefaultSettings()
	setup.SetHost("http://111.111.123.123:8087/v2")
	setup.SetUseDigest(true)
	setup.SetUsername("my_secret_username")
	setup.SetPassword("my_super_cool_password")

	// Connect to the server or deal with statistics
	server := wserest.NewServer(setup)
	sf := wserest.NewStatistics(setup)

	response, err := sf.GetServerStatistics(server)
	if err != nil {
		log.Fatal(err)
	}

## License
See [LICENSE](LICENSE).

Copyright © CyberAgent, Inc. All Rights Reserved.
