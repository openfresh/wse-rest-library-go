package wserest

import (
	"testing"
	"time"

	"github.com/openfresh/wse-rest-library-go/entity/application/helper"
)

func TestDvrClip(t *testing.T) {
	// Create settings
	settings := helper.NewDefaultSettings()
	settings.SetHost(WowzaHost)
	settings.SetUsername(WowzaUsername)
	settings.SetPassword(WowzaPassword)

	sf := NewDvrClipExtraction(settings, "ndvr", "")

	if UseWowza == "" {
		return
	}

	response, err := sf.ClearCache()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)

	response, err = sf.DebugConversions("tmp127")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)

	now := time.Now()
	duration := time.Duration(5) * time.Second
	response, err = sf.ConvertByDurationWithEndTime("tmp123", &now, &duration, "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)

	hourago := time.Now().Add(-1 * time.Hour)
	response, err = sf.ConvertByDurationWithStartTime("tmp123", &hourago, &duration, "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)

	now = time.Now()
	response, err = sf.Convert("tmp127", &now, nil, "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)

	response, err = sf.ConvertGroup([]string{"tmp123", "tmp124"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)

	response, err = sf.Convert("tmp123", nil, nil, "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)

	response, err = sf.GetItem("tmp123")
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
