package wserest

import (
	"testing"

	"github.com/openfresh/wse-rest-library-go/entity/application/helper"
)

func TestUser(t *testing.T) {
	// Create settings
	settings := helper.NewDefaultSettings()
	settings.SetHost(WowzaHost)
	settings.SetUseDigest(true)
	settings.SetUsername(WowzaUsername)
	settings.SetPassword(WowzaPassword)

	sf := NewUser(settings, "newuser3")

	if UseWowza == "" {
		return
	}

	response, err := sf.Create("newpass4", []string{"admin"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)

	response, err = sf.GetAll()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)

	response, err = sf.Remove()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}
