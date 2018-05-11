package wserest

import (
	"github.com/openfresh/wse-rest-library-go/entity/application/helper"
	"github.com/openfresh/wse-rest-library-go/entity/base"
)

// Server is server utility
type Server struct {
	wowza
}

// NewServer creates Server object
func NewServer(settings *helper.Settings) *Server {
	s := new(Server)
	s.init(settings)
	s.baseURI = s.host() + "/servers/" + s.serverInstance()
	return s
}

// GetUsers retrieves the list of server Users
func (s *Server) GetUsers() (map[string]interface{}, error) {
	s.setRestURI(s.baseURI + "/users")

	return s.sendRequest(s.preparePropertiesForRequest(), []base.Entity{}, GET, "")
}

// CreateUser adds a new server User to the list
func (s *Server) CreateUser(name string, password string, groups []string) (map[string]interface{}, error) {
	s.setRestURI(s.baseURI + "/users/" + name)
	s.AddAdditionalParameter("name", name)
	s.AddAdditionalParameter("password", password)
	s.AddAdditionalParameter("groups", groups)

	return s.sendRequest(s.preparePropertiesForRequest(), []base.Entity{}, POST, "")
}

// RemoveUser deletes the specified User configuration
func (s *Server) RemoveUser(name string) (map[string]interface{}, error) {
	s.setRestURI(s.baseURI + "/users/" + name)

	return s.sendRequest(s.preparePropertiesForRequest(), []base.Entity{}, DELETE, "")
}
