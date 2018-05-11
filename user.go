package wserest

import (
	"github.com/openfresh/wse-rest-library-go/entity/application/helper"
	"github.com/openfresh/wse-rest-library-go/entity/base"
)

// User is Server Users utiliyt
type User struct {
	wowza
}

// NewUser create User object
func NewUser(settings *helper.Settings, userName string) *User {
	u := new(User)
	u.init(settings)
	u.props["userName"] = userName
	u.props["password"] = ""
	u.props["groups"] = []string{}
	u.baseURI = u.host() + "/servers/" + u.serverInstance() + "/users"
	return u
}

// Create adds a new server User to the list
func (u *User) Create(password string, group []string) (map[string]interface{}, error) {
	u.props["password"] = password
	u.props["group"] = group
	u.setRestURI(u.baseURI)
	response, err := u.sendRequest(u.preparePropertiesForRequest(), []base.Entity{}, POST, "")

	return response, err
}

// GetAll retrieves the list of server Users
func (u *User) GetAll() (map[string]interface{}, error) {
	u.AddSkipParameter("userName")
	u.AddSkipParameter("password")
	u.AddSkipParameter("group")
	u.setRestURI(u.baseURI)

	return u.sendRequest(u.preparePropertiesForRequest(), []base.Entity{}, GET, "")
}

// Remove deletes the specified User configuration
func (u *User) Remove() (map[string]interface{}, error) {
	u.setRestURI(u.baseURI + "/" + u.props["userName"].(string))

	return u.sendRequest(u.preparePropertiesForRequest(), []base.Entity{}, DELETE, "")
}
