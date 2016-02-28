package five00px

import "testing"

func TestTransportError(t *testing.T) {
	err := httpError{42, "msg"}

	if err.Error() != "HTTP error 42. Message: msg" {
		t.Fatal("Invalid error message")
	}
}
