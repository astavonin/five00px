// Package five00px provides ...
package five00px

import "errors"

// User 500px structure. See https://github.com/500px/api-documentation/blob/master/basics/formats_and_terms.md#full-format
// for additional information.
type User struct {
	About                 string        `json:"about"`
	Affection             int           `json:"affection"`
	AllowSaleRequests     int           `json:"allow_sale_requests"`
	AnalyticsCode         interface{}   `json:"analytics_code"`
	Auth                  Auth          `json:"auth"`
	Avatars               Avatars       `json:"avatars"`
	Birthday              string        `json:"birthday"`
	City                  string        `json:"city"`
	Contacts              Contacts      `json:"contacts"`
	Country               string        `json:"country"`
	CoverURL              interface{}   `json:"cover_url"`
	CustomLicensePrices   []interface{} `json:"custom_license_prices"`
	Domain                string        `json:"domain"`
	Email                 string        `json:"email"`
	Equipment             Equipment     `json:"equipment"`
	Firstname             string        `json:"firstname"`
	FollowersCount        int           `json:"followers_count"`
	Following             bool          `json:"following"`
	FotomotoOn            bool          `json:"fotomoto_on"`
	FriendsCount          int           `json:"friends_count"`
	Fullname              string        `json:"fullname"`
	GalleriesCount        int           `json:"galleries_count"`
	ID                    int           `json:"id"`
	InFavoritesCount      int           `json:"in_favorites_count"`
	InviteAccepted        bool          `json:"invite_accepted"`
	InvitePending         bool          `json:"invite_pending"`
	Lastname              string        `json:"lastname"`
	Locale                string        `json:"locale"`
	PhotosCount           int           `json:"photos_count"`
	PresubmitForLicensing interface{}   `json:"presubmit_for_licensing"`
	RegistrationDate      string        `json:"registration_date"`
	Sex                   int           `json:"sex"`
	ShowNude              bool          `json:"show_nude"`
	State                 string        `json:"state"`
	StoreOn               bool          `json:"store_on"`
	UpgradeStatus         int           `json:"upgrade_status"`
	UpgradeStatusExpiry   string        `json:"upgrade_status_expiry"`
	UpgradeType           int           `json:"upgrade_type"`
	UploadLimit           int           `json:"upload_limit"`
	UploadLimitExpiry     string        `json:"upload_limit_expiry"`
	Username              string        `json:"username"`
	UserpicHTTPSURL       string        `json:"userpic_https_url"`
	UserpicURL            string        `json:"userpic_url"`
	Usertype              int           `json:"usertype"`
}

// Equipment structure contains dictionary of a user's equipment
type Equipment struct {
	Camera []string `json:"camera"`
	Lens   []string `json:"lens"`
	Misc   []string `json:"misc"`
}

// Auth structure contains dictionary of a user's social network authentications.
type Auth struct {
	Facebook     int `json:"facebook"`
	GoogleOauth2 int `json:"google_oauth2"`
	Twitter      int `json:"twitter"`
}

// Contacts structure contacts dictionary of user’s contacts, object
type Contacts struct {
	Facebook     string `json:"facebook"`
	Facebookpage string `json:"facebookpage"`
	Googleplus   string `json:"googleplus"`
	Skype        string `json:"skype"`
	Twitter      string `json:"twitter"`
	Website      string `json:"website"`
}

// URLInfo is used in Avatars structure
type URLInfo struct {
	HTTP  string `json:"http"`
	HTTPS string `json:"https"`
}

// Avatars is a dictionary of different avatar sizes
type Avatars struct {
	Default URLInfo `json:"default"`
	Large   URLInfo `json:"large"`
	Small   URLInfo `json:"small"`
	Tiny    URLInfo `json:"tiny"`
}

// Friends structure is a list of friends for the specified user
type Friends struct {
	Users        []User `json:"friends"`
	FriendsCount int    `json:"friends_count"`
	FriendsPages int    `json:"friends_pages"`
	Page         int    `json:"page"`
}

// ErrUserNotFound error.
var ErrUserNotFound = errors.New("User not found")
