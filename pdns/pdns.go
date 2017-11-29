// Package pdns adapts the lego powerdns DNS
// provider for Caddy. Importing this package plugs it in.
package pdns

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/xenolf/lego/providers/dns/pdns"
)

func init() {
	caddytls.RegisterDNSProvider("pdns", NewDNSProvider)
}

// NewDNSProvider returns a new pdns DNS challenge provider.
// The credentials are detected automatically; see underlying
// package docs for details:
// https://godoc.org/github.com/xenolf/lego/providers/dns/pdns
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return pdns.NewDNSProvider()
	default:
		return nil, errors.New("invalid credentials length")
	}
}
