package wserest

import (
	"net/http"
)

type digestRequest struct {
	Body     string
	Method   string
	Password string
	URI      string
	Username string
	Header   http.Header
	Auth     *authorization
	Wa       *wwwAuthenticate
	CertVal  bool
}

func (dr *digestRequest) UpdateRequest(username, password, method, uri, body string) *digestRequest {
	dr.Body = body
	dr.Method = method
	dr.Password = password
	dr.URI = uri
	dr.Username = username
	dr.Header = make(map[string][]string)
	return dr
}
