package wserest

import (
	"testing"

	"github.com/openfresh/wse-rest-library-go/entity/application/helper"
)

func TestPushPublish(t *testing.T) {
	// Create settings
	settings := helper.NewDefaultSettings()
	settings.SetHost(WowzaHost)
	settings.SetUsername(WowzaUsername)
	settings.SetPassword(WowzaPassword)

	sf := NewStreamTarget(settings, "live")

	if UseWowza == "" {
		return
	}

	response, err := sf.Create("myStream", "ppsource", "rtmp", "locahost", "", "", "myStream", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)

	response, err = sf.GetAll()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)

	response, err = sf.Remove("ppsource")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}
