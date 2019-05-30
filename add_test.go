package add

import "testing"

func TestAdd(t *testing.T) {
	actualResult := Add(1, 1)
	expected := 2

	if actualResult != 2 {
		t.Errorf("expected '%d' but got '%d'", expected, actualResult)
	}
}
