package application

import (
	"github.com/openfresh/wse-rest-library-go/entity/base"
)

type SecurityConfig struct {
	base.EntityBase
	SecureTokenVersion               int    `json:"secureTokenVersion"`
	ClientStreamWriteAccess          string `json:"clientStreamWriteAccess"`
	PublishRequirePassword           bool   `json:"publishRequirePassword"`
	PublishPasswordFile              string `json:"publishPasswordFile"`
	PublishRTMPSecureURL             string `json:"publishRTMPSecureURL"`
	PublishIPBlackList               string `json:"publishIPBlackList"`
	PublishIPWhiteList               string `json:"publishIPWhiteList"`
	PublishBlockDuplicateStreamNames bool   `json:"publishBlockDuplicateStreamNames"`
	PublishValidEncoders             string `json:"publishValidEncoders"`
	PublishAuthenticationMethod      string `json:"publishAuthenticationMethod"`
	PlayMaximumConnections           int    `json:"playMaximumConnections"`
	PlayRequireSecureConnection      bool   `json:"playRequireSecureConnection"`
	SecureTokenSharedSecret          string `json:"secureTokenSharedSecret"`
	SecureTokenUseTEAForRTMP         bool   `json:"secureTokenUseTEAForRTMP"`
	SecureTokenIncludeClientIPInHash bool   `json:"secureTokenIncludeClientIPInHash"`
	SecureTokenHashAlgorithm         string `json:"secureTokenHashAlgorithm"`
	SecureTokenQueryParametersPrefix string `json:"secureTokenQueryParametersPrefix"`
	SecureTokenOriginSharedSecret    string `json:"secureTokenOriginSharedSecret"`
	PlayIPBlackList                  string `json:"playIPBlackList"`
	PlayIPWhiteList                  string `json:"playIPWhiteList"`
	PlayAuthenticationMethod         string `json:"playAuthenticationMethod"`
}

func NewSecurityConfig() *SecurityConfig {
	s := new(SecurityConfig)
	s.SecureTokenVersion = 0
	s.ClientStreamWriteAccess = "*"
	s.PublishRequirePassword = true
	s.PublishPasswordFile = ""
	s.PublishRTMPSecureURL = ""
	s.PublishIPBlackList = ""
	s.PublishIPWhiteList = ""
	s.PublishBlockDuplicateStreamNames = false
	s.PublishValidEncoders = ""
	s.PublishAuthenticationMethod = "digest"
	s.PlayMaximumConnections = 0
	s.PlayRequireSecureConnection = false
	s.SecureTokenSharedSecret = ""
	s.SecureTokenUseTEAForRTMP = false
	s.SecureTokenIncludeClientIPInHash = false
	s.SecureTokenHashAlgorithm = ""
	s.SecureTokenQueryParametersPrefix = ""
	s.SecureTokenOriginSharedSecret = ""
	s.PlayIPBlackList = ""
	s.PlayIPWhiteList = ""
	s.PlayAuthenticationMethod = "none"
	return s
}

func (s *SecurityConfig) EntityName() string {
	return "securityConfig"
}

func (s *SecurityConfig) SetURI(baseURI string) {
	s.SetRestURI(baseURI + "/security")
}
