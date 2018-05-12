package wserest

import (
	"testing"

	"github.com/openfresh/wse-rest-library-go/entity/application/helper"
)

func TestSmilFiles(t *testing.T) {
	// Create settings
	settings := helper.NewDefaultSettings()
	settings.SetHost(WowzaHost)
	settings.SetUseDigest(true)
	settings.SetUsername(WowzaUsername)
	settings.SetPassword(WowzaPassword)

	sf := NewSmilFile(settings, "live")

	if UseWowza == "" {
		return
	}

	/*
	   * [
	          {
	             "systemLanguage": "en",
	             "src": "myfile_750.mp4",
	             "systemBitrate": "50000",
	             "type": "video",
	             "audioBitrate": "44100",
	             "videoBitrate": "750000",
	             "restURI": "http://localhost:8087/v2/servers/_defaultServer_/vhosts/_defaultVHost_/applications/live/smilfiles/mytestsmil",
	             "width": "640",
	             "height": "360"
	         },
	         {
	             "systemLanguage": "en",
	             "src": "myfile_1100.mp4",
	             "systemBitrate": "50000",
	             "type": "video",
	             "audioBitrate": "44100",
	             "videoBitrate": "1100000",
	             "restURI": "http://localhost:8087/v2/servers/_defaultServer_/vhosts/_defaultVHost_/applications/live/smilfiles/mytestsmil",
	             "width": "640",
	             "height": "360"
	         }
	*/
	stream0 := make(map[string]interface{})
	stream0["systemLanguage"] = "en"
	stream0["src"] = "myfile_750.mp4"
	stream0["systemBitrate"] = "50000"
	stream0["type"] = "video"
	stream0["audioBitrate"] = "44100"
	stream0["videoBitrate"] = "750000"
	stream0["restURI"] = "http://localhost:8087/v2/servers/_defaultServer_/vhosts/_defaultVHost_/applications/live/smilfiles/mytestsmil"
	stream0["width"] = "640"
	stream0["height"] = "360"

	stream1 := make(map[string]interface{})
	stream1["systemLanguage"] = "en"
	stream1["src"] = "myfile_1100.mp4"
	stream1["systemBitrate"] = "50000"
	stream1["type"] = "video"
	stream1["audioBitrate"] = "44100"
	stream1["videoBitrate"] = "1100000"
	stream1["restURI"] = "http://localhost:8087/v2/servers/_defaultServer_/vhosts/_defaultVHost_/applications/live/smilfiles/mytestsmil"
	stream1["width"] = "640"
	stream1["height"] = "360"

	streams := make([]map[string]interface{}, 2)
	streams[0] = stream0
	streams[1] = stream1

	response, err := sf.Create("newsmil", streams)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)

	response, err = sf.Remove("newsmil")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)

	response, err = sf.GetAll()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}
