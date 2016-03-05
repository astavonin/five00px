// Package five00px provides ...
package five00px

// Photo is an array of prhotos for selected criterias
type Photo struct {
	Aperture               string   `json:"aperture"`
	Camera                 string   `json:"camera"`
	Category               Category `json:"category"`
	CollectionsCount       int      `json:"collections_count"`
	CommentsCount          int      `json:"comments_count"`
	Converted              int      `json:"converted"`
	ConvertedBits          int      `json:"converted_bits"`
	CreatedAt              string   `json:"created_at"`
	CropVersion            int      `json:"crop_version"`
	Description            string   `json:"description"`
	Disliked               bool     `json:"disliked"`
	Favorited              bool     `json:"favorited"`
	FavoritesCount         int      `json:"favorites_count"`
	FocalLength            string   `json:"focal_length"`
	ForSale                bool     `json:"for_sale"`
	ForSaleDate            string   `json:"for_sale_date"`
	Height                 int      `json:"height"`
	HiResUploaded          int      `json:"hi_res_uploaded"`
	HighestRating          float32  `json:"highest_rating"`
	HighestRatingDate      string   `json:"highest_rating_date"`
	ID                     int      `json:"id"`
	ImageFormat            string   `json:"image_format"`
	ImageURL               string   `json:"image_url"`
	Iso                    string   `json:"iso"`
	Latitude               float32  `json:"latitude"`
	Lens                   string   `json:"lens"`
	LicenseRequestsEnabled bool     `json:"license_requests_enabled"`
	LicenseType            int      `json:"license_type"`
	LicensingRequested     bool     `json:"licensing_requested"`
	Liked                  bool     `json:"liked"`
	Location               string   `json:"location"`
	Longitude              float32  `json:"longitude"`
	Name                   string   `json:"name"`
	Nsfw                   bool     `json:"nsfw"`
	PositiveVotesCount     int      `json:"positive_votes_count"`
	Privacy                bool     `json:"privacy"`
	Profile                bool     `json:"profile"`
	Purchased              bool     `json:"purchased"`
	Rating                 float32  `json:"rating"`
	RequestToBuyEnabled    bool     `json:"request_to_buy_enabled"`
	SalesCount             int      `json:"sales_count"`
	ShutterSpeed           string   `json:"shutter_speed"`
	Status                 int      `json:"status"`
	StoreDownload          bool     `json:"store_download"`
	StoreLicense           bool     `json:"store_license"`
	StorePrint             bool     `json:"store_print"`
	Tags                   []string `json:"tags"`
	TakenAt                string   `json:"taken_at"`
	TimesViewed            int      `json:"times_viewed"`
	URL                    string   `json:"url"`
	User                   User     `json:"user"`
	UserID                 int      `json:"user_id"`
	Voted                  bool     `json:"voted"`
	VotesCount             int      `json:"votes_count"`
	Watermark              bool     `json:"watermark"`
	Width                  int      `json:"width"`
}

// Filters NOTE: Category and Exclude are interfaces as they are false or int values
type Filters struct {
	Category interface{} `json:"category"`
	Exclude  interface{} `json:"exclude"`
}

// Photos structure contains a listing of (up to one hundred) photos for a specified photo stream.
type Photos struct {
	CurrentPage int     `json:"current_page"`
	Feature     Feature `json:"feature"`
	Filters     Filters `json:"filters"`
	Photos      []Photo `json:"photos"`
	TotalItems  int     `json:"total_items"`
	TotalPages  int     `json:"total_pages"`
}

// Category helper type
type Category int

// Categories helper type
type Categories []Category

// Size helper type
type Size int

// Feature helper type
type Feature string

// SortBy helper type
type SortBy string

// SortOrder helper type
type SortOrder string

// Units helper type
type Units string

// License type
type License int

// Constants
const (
	FeaturePopular        = Feature("popular")
	FeatureHRated         = Feature("highest_rated")
	FeatureUpcoming       = Feature("upcoming")
	FeatureEditors        = Feature("editors")
	FeatureFreshToday     = Feature("fresh_today")
	FeatureFreshYesterday = Feature("fresh_yesterday")
	FeatureFreshWeek      = Feature("fresh_week")

	Size70x70   = Size(1)
	Size140x140 = Size(2)
	Size280x280 = Size(3)
	Size100x100 = Size(100)
	Size200x200 = Size(200)
	Size440x440 = Size(440)
	Size600x600 = Size(600)
	Size900l    = Size(4)
	Size1170l   = Size(5)
	Size1080h   = Size(6)
	Size300h    = Size(20)
	Size600h    = Size(21)
	Size256l    = Size(30)
	Size450h    = Size(31)
	Size1080l   = Size(1080)
	Size1600l   = Size(1600)
	Size2048l   = Size(2048)

	CategoryAll              = Category(-1)
	CategoryUncategorized    = Category(0)
	CategoryAbstract         = Category(10)
	CategoryAnimals          = Category(11)
	CategoryBW               = Category(5)
	CategoryCelebrities      = Category(1)
	CategoryArch             = Category(9)
	CategoryCommercial       = Category(15)
	CategoryConcert          = Category(16)
	CategoryFamily           = Category(20)
	CategoryFashion          = Category(14)
	CategoryFilm             = Category(2)
	CategoryFineArt          = Category(24)
	CategoryFood             = Category(23)
	CategoryJournalism       = Category(3)
	CategoryLandscapes       = Category(8)
	CategoryMacro            = Category(12)
	CategoryNature           = Category(18)
	CategoryNude             = Category(4)
	CategoryPeople           = Category(7)
	CategoryPerformingArts   = Category(19)
	CategorySport            = Category(17)
	CategoryStillLife        = Category(6)
	CategoryStreet           = Category(21)
	CategoryTransportation   = Category(26)
	CategoryTravel           = Category(13)
	CategoryUnderwater       = Category(22)
	CategoryUrbanExploration = Category(27)
	CategoryWedding          = Category(25)

	SortByCreatedAt     = SortBy("created_at")
	SortByRating        = SortBy("rating")
	SortByHighestRating = SortBy("highest_rating")
	SortByTimesViewed   = SortBy("times_viewed")
	SortByVotesCount    = SortBy("votes_count")
	SortByCommentsCount = SortBy("comments_count")
	SortByTakenAt       = SortBy("taken_at")

	SortOrderAsk  = SortOrder("ask")
	SortOrderDesk = SortOrder("desk")

	UnitsKm = Units("km")
	UnitsMi = Units("mi")

	LicAll                   = License(-1)
	Lic500px                 = License(0)
	LicCrCommonNonComAttr    = License(1)
	LicCrCommonNonComNoDeriv = License(2)
	LicCrCommonNonComShare   = License(3)
	LicCrCommonAttr          = License(4)
	LicCrCommonNoDeriv       = License(5)
	LicCrCommonShare         = License(6)
)

var categoriesMap = map[Category]string{
	CategoryUncategorized:    "Uncategorized",
	CategoryAbstract:         "Abstract",
	CategoryAnimals:          "Animals",
	CategoryBW:               "Black and White",
	CategoryCelebrities:      "Celebrities",
	CategoryArch:             "City and Architecture",
	CategoryCommercial:       "Commercial",
	CategoryConcert:          "Concert",
	CategoryFamily:           "Family",
	CategoryFashion:          "Fashion",
	CategoryFilm:             "Film",
	CategoryFineArt:          "Fine Art",
	CategoryFood:             "Food",
	CategoryJournalism:       "Journalism",
	CategoryLandscapes:       "Landscapes",
	CategoryMacro:            "Macro",
	CategoryNature:           "Nature",
	CategoryNude:             "Nude",
	CategoryPeople:           "People",
	CategoryPerformingArts:   "Performing Arts",
	CategorySport:            "Sport",
	CategoryStillLife:        "Still Life",
	CategoryStreet:           "Street",
	CategoryTransportation:   "Transportation",
	CategoryTravel:           "Travel",
	CategoryUnderwater:       "Underwater",
	CategoryUrbanExploration: "Urban Exploration",
	CategoryWedding:          "Wedding",
}
