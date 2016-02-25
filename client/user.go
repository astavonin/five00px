// Package five00px provides ...
package five00px

type User struct {
	About             string      `json:"about"`
	Affection         int         `json:"affection"`
	AllowSaleRequests int         `json:"allow_sale_requests"`
	AnalyticsCode     interface{} `json:"analytics_code"`
	Auth              struct {
		Facebook     int `json:"facebook"`
		GoogleOauth2 int `json:"google_oauth2"`
		Twitter      int `json:"twitter"`
	} `json:"auth"`
	Avatars struct {
		Default struct {
			HTTP  string `json:"http"`
			HTTPS string `json:"https"`
		} `json:"default"`
		Large struct {
			HTTP  string `json:"http"`
			HTTPS string `json:"https"`
		} `json:"large"`
		Small struct {
			HTTP  string `json:"http"`
			HTTPS string `json:"https"`
		} `json:"small"`
		Tiny struct {
			HTTP  string `json:"http"`
			HTTPS string `json:"https"`
		} `json:"tiny"`
	} `json:"avatars"`
	Birthday string `json:"birthday"`
	City     string `json:"city"`
	Contacts struct {
		Facebook     string `json:"facebook"`
		Facebookpage string `json:"facebookpage"`
		Googleplus   string `json:"googleplus"`
		Skype        string `json:"skype"`
		Twitter      string `json:"twitter"`
		Website      string `json:"website"`
	} `json:"contacts"`
	Country             string        `json:"country"`
	CoverURL            interface{}   `json:"cover_url"`
	CustomLicensePrices []interface{} `json:"custom_license_prices"`
	Domain              string        `json:"domain"`
	Email               string        `json:"email"`
	Equipment           struct {
		Camera []string `json:"camera"`
		Lens   []string `json:"lens"`
	} `json:"equipment"`
	Firstname             string      `json:"firstname"`
	FollowersCount        int         `json:"followers_count"`
	FotomotoOn            bool        `json:"fotomoto_on"`
	FriendsCount          int         `json:"friends_count"`
	Fullname              string      `json:"fullname"`
	GalleriesCount        int         `json:"galleries_count"`
	ID                    int         `json:"id"`
	InFavoritesCount      int         `json:"in_favorites_count"`
	InviteAccepted        bool        `json:"invite_accepted"`
	InvitePending         bool        `json:"invite_pending"`
	Lastname              string      `json:"lastname"`
	Locale                string      `json:"locale"`
	PhotosCount           int         `json:"photos_count"`
	PresubmitForLicensing interface{} `json:"presubmit_for_licensing"`
	RegistrationDate      string      `json:"registration_date"`
	Sex                   int         `json:"sex"`
	ShowNude              bool        `json:"show_nude"`
	State                 string      `json:"state"`
	StoreOn               bool        `json:"store_on"`
	UpgradeStatus         int         `json:"upgrade_status"`
	UpgradeStatusExpiry   string      `json:"upgrade_status_expiry"`
	UpgradeType           int         `json:"upgrade_type"`
	UploadLimit           int         `json:"upload_limit"`
	UploadLimitExpiry     string      `json:"upload_limit_expiry"`
	Username              string      `json:"username"`
	UserpicHTTPSURL       string      `json:"userpic_https_url"`
	UserpicURL            string      `json:"userpic_url"`
	Usertype              int         `json:"usertype"`
}

type Friends struct {
	User []struct {
		Affection int `json:"affection"`
		Avatars   struct {
			Default struct {
				HTTPS string `json:"https"`
			} `json:"default"`
			Large struct {
				HTTPS string `json:"https"`
			} `json:"large"`
			Small struct {
				HTTPS string `json:"https"`
			} `json:"small"`
			Tiny struct {
				HTTPS string `json:"https"`
			} `json:"tiny"`
		} `json:"avatars"`
		City            string `json:"city"`
		Country         string `json:"country"`
		CoverURL        string `json:"cover_url"`
		Firstname       string `json:"firstname"`
		FollowersCount  int    `json:"followers_count"`
		Fullname        string `json:"fullname"`
		ID              int    `json:"id"`
		Lastname        string `json:"lastname"`
		StoreOn         bool   `json:"store_on"`
		UpgradeStatus   int    `json:"upgrade_status"`
		Username        string `json:"username"`
		UserpicHTTPSURL string `json:"userpic_https_url"`
		UserpicURL      string `json:"userpic_url"`
		Usertype        int    `json:"usertype"`
	} `json:"friends"`
	FriendsCount int `json:"friends_count"`
	FriendsPages int `json:"friends_pages"`
	Page         int `json:"page"`
}
