package wserest

import (
	"testing"

	"github.com/openfresh/wse-rest-library-go/entity/application/helper"
)

func TestStatistics(t *testing.T) {
	// Create settings
	settings := helper.NewDefaultSettings()
	settings.SetHost(WowzaHost)
	settings.SetUsername(WowzaUsername)
	settings.SetPassword(WowzaPassword)

	sf := NewStatistics(settings)

	if UseWowza == "" {
		return
	}

	// get stats per application
	wowzaApplication := NewApplication(settings, "vod", "", "", "", "")

	// get total server stats
	server := NewServer(settings)

	response, err := sf.GetServerStatistics(server)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)

	// get stats historical for given application
	response, err = sf.GetApplicationStatisticsHistory(wowzaApplication)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)

	response, err = sf.GetApplicationStatistics(wowzaApplication)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)

	// get incoming stream stats for given application
	response, err = sf.GetIncomingApplicationStatistics(wowzaApplication, "sample.mp4", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}
