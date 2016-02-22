// Package five00px provides main 500px API implementation
package five00px

// 500px client
type Five00px struct {
	oa oAuthInfo
}

// The call creates new Five00px structure with provided
// consumer key and secret
func New(consumerKey, consumerSecret string) *Five00px {
	return &Five00px{
		oa: newOAuth(consumerKey, consumerSecret),
	}
}

// OAuth authentication call. Default Web broser will be popped up
// during authentication.
func (f00 *Five00px) Auth() {
	f00.oa.Auth()
}
