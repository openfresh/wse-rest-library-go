package wserest

import (
	"testing"

	"github.com/openfresh/wse-rest-library-go/entity/application/helper"
)

func TestPublishers(t *testing.T) {
	// Create settings
	settings := helper.NewDefaultSettings()
	settings.SetHost(WowzaHost)
	settings.SetUseDigest(true)
	settings.SetUsername(WowzaUsername)
	settings.SetPassword(WowzaPassword)

	sf := NewPublisher(settings, "myUser")

	if UseWowza == "" {
		return
	}

	response, err := sf.Create("myPass")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)

	response, err = sf.Remove()
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
