package wserest

import (
	"testing"

	"github.com/openfresh/wse-rest-library-go/entity/application/helper"
)

func TestStreamfile(t *testing.T) {
	// Create settings
	settings := helper.NewDefaultSettings()
	settings.SetHost(WowzaHost)
	settings.SetUsername(WowzaUsername)
	settings.SetPassword(WowzaPassword)

	sf := NewStreamFile(settings, "live", "myStream")

	if UseWowza == "" {
		return
	}

	urlProps := map[string]string{
		"uri":                "rtsp://localhost/vod/mp4:BigBuckBunny_115k.mov",
		"streamTimeout":      "1200",
		"rtspSessionTimeout": "800",
	}

	response, err := sf.Create(urlProps, "rtp", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)

	// Complete an update on the stream file
	urlProps = map[string]string{
		"uri":                "rtsp://184.72.239.149/vod/mp4:BigBuckBunny_115k.mov",
		"streamTimeout":      "1100",
		"rtspSessionTimeout": "600",
	}
	response, err = sf.Update(urlProps, "", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)

	// Connect to the stream identified in the file
	response, err = sf.Connect("")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)

	response, err = sf.Disconnect()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)

	// Remove the stream file
	response, err = sf.Remove()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}
