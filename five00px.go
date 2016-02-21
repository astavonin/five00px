// Package five00px provides main 500px API implementation
package five00px

import ()

type Five00px struct {
	oa OAuthInfo
}

func New(consumerKey, consumerSecret string) *Five00px {
	return &Five00px{
		oa: newOAuth(consumerKey, consumerSecret),
	}
}

func (f00 *Five00px) Auth() {
	f00.oa.Auth()
}
