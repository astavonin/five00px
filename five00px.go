// Package five00px provides main 500px API implementation
package five00px

// Five00px client
type Five00px struct {
	oa oAuth
}

// New call creates new Five00px structure with provided
// consumer key and secret
func New(authParams OAuthData) *Five00px {
	return &Five00px{
		oa: newOAuth(authParams),
	}
}

// Auth initiate OAuth authentication call. Default Web broser will be
// popped up during authentication. Returns error on authorization failure
func (f00 *Five00px) Auth() error {
	return f00.oa.Auth()
}

// AuthParams returns pointer to internal OAuthData structure. The call
// might be usefull for reusing tocken and verifier between sessions
func (f00 *Five00px) AuthParams() *OAuthData {
	return &f00.oa.params
}
