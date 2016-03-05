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
type PhotoCriterias struct {
	Feature       Feature
	Only          Categories
	Exclude       Categories
	Sort          SortBy
	SortDirection SortOrder
	ImageSize     Size
}

// Valid returns true if PhotoCriterias a valid for futher usage
func (p *PhotoCriterias) Valid() bool {
	return p.Feature.Valid()
}

// Vals converts PhotoCriterias to url.Values
func (p *PhotoCriterias) Vals() url.Values {
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
		if int(p.ImageSize) > 0 {
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
