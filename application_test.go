package wserest

import (
	"testing"

	"github.com/openfresh/wse-rest-library-go/entity/application"
	"github.com/openfresh/wse-rest-library-go/entity/application/helper"
)

func TestApplication(t *testing.T) {
	// example setting up a stream configuration element
	streamConfig := application.NewStreamConfig()
	streamConfig.StreamType = "live"
	streamConfig.LiveStreamPacketizer = []string{"sanjosestreamingpacketizer", "cupertinostreamingpacketizer"}

	// example setting up a security configuration element
	securityConfig := application.NewSecurityConfig()
	securityConfig.SecureTokenVersion = 0
	securityConfig.ClientStreamWriteAccess = "*"
	securityConfig.PublishRequirePassword = true
	securityConfig.PublishPasswordFile = ""
	securityConfig.PublishRTMPSecureURL = ""
	securityConfig.PublishIPBlackList = ""
	securityConfig.PublishIPWhiteList = ""
	securityConfig.PublishBlockDuplicateStreamNames = false
	securityConfig.PublishValidEncoders = ""
	securityConfig.PublishAuthenticationMethod = "digest"
	securityConfig.PlayMaximumConnections = 0
	securityConfig.PlayRequireSecureConnection = false
	securityConfig.SecureTokenSharedSecret = ""
	securityConfig.SecureTokenUseTEAForRTMP = false
	securityConfig.SecureTokenIncludeClientIPInHash = false
	securityConfig.SecureTokenHashAlgorithm = ""
	securityConfig.SecureTokenQueryParametersPrefix = ""
	securityConfig.SecureTokenOriginSharedSecret = ""
	securityConfig.PlayIPBlackList = ""
	securityConfig.PlayIPWhiteList = ""
	securityConfig.PlayAuthenticationMethod = "none"

	// example setting up module(s) configuration element
	modules := application.NewModules()
	modules.ModuleList = append(modules.ModuleList, modules.GetModuleItem("ModuleCoreSecurity", "ModuleCoreSecurity", "com.wowza.wms.security.ModuleCoreSecurity", -1))

	// Create settings
	settings := helper.NewDefaultSettings()
	settings.SetHost(WowzaHost)
	settings.SetUsername(WowzaUsername)
	settings.SetPassword(WowzaPassword)

	// Create this application
	wowzaApplication := NewApplication(settings, "live", "", "", "", "")

	if UseWowza == "" {
		return
	}

	response, err := wowzaApplication.GetAll()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)

	response, err = wowzaApplication.Get()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)

	wowzaApplication2 := NewApplication(settings, "live2", "", "", "", "")
	response, err = wowzaApplication2.Create(streamConfig, nil, nil, nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer wowzaApplication2.Remove()
	t.Log(response)

	wowzaApplication3 := NewApplication(settings, "live3", "", "", "", "")
	response, err = wowzaApplication3.Create(streamConfig, securityConfig, modules, nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer wowzaApplication3.Remove()
	t.Log(response)

	// Update the application
	response, err = wowzaApplication3.Update(streamConfig, nil, nil, nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}
