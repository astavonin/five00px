package five00px

import "testing"

func TestTransportError(t *testing.T) {
	err := httpError{42, "msg"}

	if err.Error() != "HTTP error 42. Message: msg" {
		t.Fatal("Invalid error message")
	}
}

func TestBuildQuery(t *testing.T) {
	qExp := "?feature=fresh_today&image_size=4&include_states=1&include_store=1&only=Black%20and%20White,Abstract&sort=created_at&tags=1"
	c := PhotoCriterias{
		Feature:   FeatureFreshToday,
		Only:      Categories{CategoryBW, CategoryAbstract},
		Sort:      SortByCreatedAt,
		ImageSize: Size900l,
	}
	qGen := buildQuery(c.Vals())
	if qExp != qGen {
		t.Fatalf("Expecting:\t%s\nFound:\t%s", qExp, qGen)
	}
}
