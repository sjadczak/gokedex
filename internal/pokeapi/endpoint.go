package pokeapi

import (
	"net/url"
	"strings"
)

func preprocessEndpoint(endpoint string) (*url.URL, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	if u.IsAbs() {
		u.Scheme = ""
		u.Host = ""
	}

	np := strings.Replace(u.Path, "/api/v2", "", 1)
	u.Path = np

	return u, nil
}
