package wserest

import (
	"strconv"

	"github.com/openfresh/wse-rest-library-go/entity/application/helper"
	"github.com/openfresh/wse-rest-library-go/entity/base"
)

// Publisher is publisher utility
type Publisher struct {
	wowza
}

// NewPublisher creates Publisher object
func NewPublisher(settings *helper.Settings, publisherName string) *Publisher {
	p := new(Publisher)
	p.init(settings)
	p.props["name"] = publisherName
	p.props["password"] = ""
	p.baseURI = p.host() + "/servers/" + p.serverInstance() + "/publishers"
	return p
}

// Create adds a new Publisher to the list
func (p *Publisher) Create(password string) (map[string]interface{}, error) {
	p.props["password"] = password
	p.setRestURI(p.baseURI)
	response, err := p.sendRequest(p.preparePropertiesForRequest(), []base.Entity{}, POST, "")

	return response, err
}

// GetAll retrieves the list of server Publishers
func (p *Publisher) GetAll() (map[string]interface{}, error) {
	p.AddSkipParameter("name")
	p.AddSkipParameter("password")

	p.setRestURI(p.baseURI)
	return p.sendRequest(p.preparePropertiesForRequest(), []base.Entity{}, GET, "")
}

// Remove deletes the specified Publisher configuration
func (p *Publisher) Remove() (map[string]interface{}, error) {
	p.setRestURI(p.baseURI + "/" + p.props["name"].(string))

	return p.sendRequest(p.preparePropertiesForRequest(), []base.Entity{}, DELETE, "")
}

func (p *Publisher) getAdvancedSettings(urlProps map[string]interface{}) []*helper.AdvancedSettingItem {
	items := make([]*helper.AdvancedSettingItem, 0)
	for k, v := range urlProps {
		item := helper.NewAdvancedSettingItem()
		item.Name = k
		switch t := v.(type) {
		default:
		case bool:
			item.Value = strconv.FormatBool(t)
			item.Type = "Boolean"
		case int:
			item.Value = strconv.Itoa(t)
			item.Type = "Integer"
		case string:
			item.Value = t
			item.Type = "String"
		}
		items = append(items, item)
	}

	return items
}
