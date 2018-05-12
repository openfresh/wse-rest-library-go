package wserest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/openfresh/wse-rest-library-go/entity/application/helper"
	"github.com/openfresh/wse-rest-library-go/entity/base"
)

// VerbType is verb type
type VerbType int

const (
	// POST is post method
	POST VerbType = iota
	// GET is get method
	GET
	// DELETE is delete method
	DELETE
	// PUT is put method
	PUT
)

func (v VerbType) String() string {
	switch v {
	case POST:
		return "POST"
	case GET:
		return "GET"
	case DELETE:
		return "DELETE"
	case PUT:
		return "PUT"
	default:
		return "Unknown"
	}
}

type wowza struct {
	skip       map[string]interface{}
	additional map[string]interface{}
	settings   *helper.Settings
	props      map[string]interface{}
	baseURI    string
}

func (w *wowza) init(settings *helper.Settings) {
	w.settings = settings
	w.skip = make(map[string]interface{})
	w.additional = make(map[string]interface{})
	w.props = make(map[string]interface{})
}

/*func (w *wowza) restURI() string {
	return w.props["restURI"].(string)
}*/

func (w *wowza) setRestURI(uri string) {
	w.props["restURI"] = uri
}

func (w *wowza) host() string {
	return w.settings.Host()
}

func (w *wowza) serverInstance() string {
	return w.settings.ServerInstance()
}

func (w *wowza) vHostInstance() string {
	return w.settings.VhostInstance()
}

func (w *wowza) getEntities(args []base.Entity, baseURI string) []base.Entity {
	var entities []base.Entity
	for _, arg := range args {
		if arg.RestURI() == "" {
			if baseURI == "" {
				arg.ResetURI()
			} else {
				arg.SetURI(baseURI)
			}
		}
		entities = append(entities, arg)
	}
	return entities
}

func (w *wowza) debug(v ...interface{}) {
	if w.settings.IsDebug() {
		println(fmt.Sprint(v...))
	}
}

func (w *wowza) debugf(format string, v ...interface{}) {
	if w.settings.IsDebug() {
		println(fmt.Sprintf(format, v...))
	}
}

func (w *wowza) sendRequest(props map[string]interface{}, entities []base.Entity, verbType VerbType, queryParams string) (map[string]interface{}, error) {
	if restURI, ok := props["restURI"].(string); ok {
		for _, entity := range entities {
			name := entity.EntityName()
			props[name] = entity
		}
		jsonb, err := json.Marshal(props)
		if err != nil {
			return nil, err
		}

		if queryParams != "" {
			restURI += "?" + queryParams
		}
		w.debugf("JSON REQUEST to %s with verb %s: %+v", restURI, verbType, props)

		client := &http.Client{Timeout: 10 * time.Second}

		req, err := http.NewRequest(verbType.String(), restURI, bytes.NewReader(jsonb))
		req.Header.Add("Accept", "application/json; charset=utf-8")
		req.Header.Add("Content-type", "application/json; charset=utf-8")
		req.Header.Add("Content-Length", strconv.Itoa(len(jsonb)))
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		if w.settings.IsUseDigest() && resp.StatusCode == http.StatusUnauthorized {
			resp.Body.Close()
			var (
				auth     *authorization
				wa       *wwwAuthenticate
				waString string
			)
			if waString = resp.Header.Get("WWW-Authenticate"); waString == "" {
				return nil, fmt.Errorf("failed to get WWW-Authenticate header, please check your server configuration")
			}
			wa = newWwwAuthenticate(waString)
			dr := new(digestRequest)
			dr.UpdateRequest(w.settings.Username(), w.settings.Password(), verbType.String(), restURI, string(jsonb))
			dr.CertVal = true
			dr.Wa = wa
			if auth, err = newAuthorization(dr); err != nil {
				return nil, err
			}
			req, err = http.NewRequest(verbType.String(), restURI, bytes.NewReader(jsonb))
			req.Header.Add("Accept", "application/json; charset=utf-8")
			req.Header.Add("Content-type", "application/json; charset=utf-8")
			req.Header.Add("Content-Length", strconv.Itoa(len(jsonb)))
			req.Header.Add("Authorization", auth.toString())
			resp, err = client.Do(req)
			if err != nil {
				return nil, err
			}
		}
		defer resp.Body.Close()
		contents := make(map[string]interface{})
		json.NewDecoder(resp.Body).Decode(&contents)

		w.debugf("RETURN: %+v", contents)

		w.skip = make(map[string]interface{})
		w.additional = make(map[string]interface{})

		return contents, nil
	}
	return nil, errors.New("no restURI")
}

func (w *wowza) AddAdditionalParameter(key string, value interface{}) {
	w.additional[key] = value
}

func (w *wowza) AddSkipParameter(key string) {
	w.skip[key] = true
}

func (w *wowza) preparePropertiesForRequest() map[string]interface{} {
	newProps := make(map[string]interface{})
	for k, v := range w.props {
		if _, ok := w.skip[k]; ok {
			continue
		}
		newProps[k] = v
	}

	for k, v := range w.additional {
		newProps[k] = v
	}

	return newProps
}
