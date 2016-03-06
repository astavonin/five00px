// Package five00px provides ...
package five00px

import (
	"net/url"
	"strconv"
	"strings"
)

// Helper interfaces section

// Validator interface provides ability to test argument validity
type Validator interface {
	Valid() bool
}

// PhotoCriterias section

// PhotoCriterias ...
type StreamCriterias struct {
	Feature       Feature
	Only          Categories
	Exclude       Categories
	Sort          SortBy
	SortDirection SortOrder
	ImageSize     Size
}

// Valid returns true if PhotoCriterias a valid for futher usage
func (p *StreamCriterias) Valid() bool {
	return p.Feature.Valid()
}

func NewPhotoCriterias() *StreamCriterias {
	return &StreamCriterias{
		Feature: FeaturePopular,
	}
}

// Vals converts PhotoCriterias to url.Values
func (p *StreamCriterias) Vals() url.Values {
	vals := url.Values{}
	if p != nil {
		vals.Add("feature", p.Feature.String())
		if p.Only.Valid() {
			vals.Add("only", p.Only.String())
		}
		if p.Exclude.Valid() {
			vals.Add("exclude", p.Exclude.String())
		}
		if p.Sort.Valid() {
			vals.Add("sort", p.Sort.String())
		}
		if p.SortDirection.Valid() {
			vals.Add("sort_direction", p.SortDirection.String())
		}
		if p.ImageSize.Valid() {
			vals.Add("image_size", p.ImageSize.String())
		}
		vals.Add("include_store", "1")
		vals.Add("include_states", "1")
		vals.Add("tags", "1")
	}
	return vals
}

// Feature section

// Feature to string
func (f Feature) String() string {
	return string(f)
}

// Valid if Feature contains valid value
func (f Feature) Valid() bool {
	switch f {
	case
		FeaturePopular,
		FeatureHRated,
		FeatureUpcoming,
		FeatureEditors,
		FeatureFreshToday,
		FeatureFreshYesterday,
		FeatureFreshWeek:
		return true
	}
	return false
}

// Size section

// Size to string
func (s Size) String() string {
	return strconv.Itoa(int(s))
}

// Valid if Size contains valid size
func (s Size) Valid() bool {
	switch s {
	case
		Size70x70,
		Size140x140,
		Size280x280,
		Size100x100,
		Size200x200,
		Size440x440,
		Size600x600,
		Size900l,
		Size1170l,
		Size1080h,
		Size300h,
		Size600h,
		Size256l,
		Size450h,
		Size1080l,
		Size1600l,
		Size2048l:
		return true
	}
	return false
}

// Category section

// String representation of Categories
func (c Categories) String() string {
	tmp := []string{}
	for _, v := range c {
		tmp = append(tmp, v.String())
	}
	return strings.Join(tmp, ",")
}

// Valid if Categories contains valid values
func (c Categories) Valid() bool {
	if len(c) <= 0 {
		return false
	}
	for _, v := range c {
		if !v.Valid() {
			return false
		}
	}
	return true
}

// String representation for Category
func (c Category) String() string {
	return categoriesMap[c]
}

// Valid if Category contains valid values
func (c Category) Valid() bool {
	switch c {
	case
		CategoryUncategorized,
		CategoryAbstract,
		CategoryAnimals,
		CategoryBW,
		CategoryCelebrities,
		CategoryArch,
		CategoryCommercial,
		CategoryConcert,
		CategoryFamily,
		CategoryFashion,
		CategoryFilm,
		CategoryFineArt,
		CategoryFood,
		CategoryJournalism,
		CategoryLandscapes,
		CategoryMacro,
		CategoryNature,
		CategoryNude,
		CategoryPeople,
		CategoryPerformingArts,
		CategorySport,
		CategoryStillLife,
		CategoryStreet,
		CategoryTransportation,
		CategoryTravel,
		CategoryUnderwater,
		CategoryUrbanExploration,
		CategoryWedding:
		return true
	}
	return false
}

// SortBy section

// String representation for SortBy
func (s SortBy) String() string {
	return string(s)
}

//Valid if SortBy contains acceptable value
func (s SortBy) Valid() bool {
	switch s {
	case SortByRating,
		SortByTakenAt,
		SortByCreatedAt,
		SortByTimesViewed,
		SortByCommentsCount,
		SortByVotesCount,
		SortByHighestRating:
		return true
	}
	return false
}

// SortBy section

// String representation for SortOrder
func (s SortOrder) String() string {
	return string(s)
}

// Valid if SortOrder contains acceptable value
func (s SortOrder) Valid() bool {
	switch s {
	case SortOrderAsk,
		SortOrderDesk:
		return true

	}
	return false
}

// Search section

// SearchCriterias intro information for photo search request
type SearchCriterias struct {
	Term        string
	Tag         string
	Geo         Geo
	Only        Categories
	Exclude     Categories
	ImageSize   Size
	LicenseType License
	Sort        SortBy
}

func NewSearchCriterias() *SearchCriterias {
	return &SearchCriterias{
		LicenseType: LicAll,
	}
}

func (s *SearchCriterias) Valid() bool {
	return s.Term != "" || s.Tag != ""
}

func (s *SearchCriterias) Vals() url.Values {
	vals := url.Values{}
	if s != nil {
		if s.Term != "" {
			vals.Add("term", s.Term)
		}
		if s.Tag != "" {
			vals.Add("tag", s.Tag)
		}
		if s.Geo.Valid() {
			vals.Add("geo", s.Geo.String())
		}
		if s.Exclude.Valid() {
			vals.Add("only", s.Exclude.String())
		}
		if s.Only.Valid() {
			vals.Add("only", s.Only.String())
		}
		if s.ImageSize.Valid() {
			vals.Add("image_size", s.ImageSize.String())
		}
		if s.LicenseType.Valid() {
			vals.Add("license_type", s.LicenseType.String())
		}
		vals.Add("tags", "1")
	}
	return vals
}

// Geo geo-location point
type Geo struct {
	Latitude  string
	Longitude string
	Radius    string
	Units     Units
}

func (u Units) String() string {
	return string(u)
}

// Valid units are km and mi
func (u Units) Valid() bool {
	switch u {
	case UnitsMi,
		UnitsKm:
		return true
	}
	return false
}

// Valid units are km and mi
func (g Geo) Valid() bool {
	return g.Units.Valid()
}

func (g Geo) String() string {
	return g.Latitude + "," + g.Longitude + "," + g.Radius + "<" + g.Units.String() + ">"
}

// Valid license type
func (l License) Valid() bool {
	switch l {
	case Lic500px,
		LicCrCommonNonComAttr,
		LicCrCommonNonComNoDeriv,
		LicCrCommonNonComShare,
		LicCrCommonAttr,
		LicCrCommonNoDeriv,
		LicCrCommonShare:
		return true
	}
	return false
}

func (l License) String() string {
	return strconv.Itoa(int(l))
}

type PhotoInfo struct {
	ImageSize    Size
	Comments     bool
	CommentsPage int
}

func (p *PhotoInfo) Vals() url.Values {
	vals := url.Values{}
	if p != nil {
		if p.ImageSize.Valid() {
			vals.Add("image_size", p.ImageSize.String())
		}
		if p.Comments != false {
			vals.Add("comments", "1")
		}
		if p.CommentsPage != 0 {
			vals.Add("comments_page", strconv.Itoa(p.CommentsPage))
		}
	}
	vals.Add("tags", "1")
	return vals
}
