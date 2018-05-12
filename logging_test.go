package wserest

import (
	"testing"

	"github.com/openfresh/wse-rest-library-go/entity/application/helper"
)

func TestLogging(t *testing.T) {
	// Create settings
	settings := helper.NewDefaultSettings()
	settings.SetHost(WowzaHost)
	settings.SetUseDigest(true)
	settings.SetUsername(WowzaUsername)
	settings.SetPassword(WowzaPassword)

	sf := NewLogging(settings)

	if UseWowza == "" {
		return
	}

	response, err := sf.Search("MediaCasterStreamValidator.init")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)

	response, err = sf.GetLineCount(10)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)

	response, err = sf.GetNewestFirst()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}
