package reverse

//START OMIT
import "testing"

func TestReverse(t *testing.T) {
	r := Reverse("hellogdg")
	if r != "gdgolleh" {
		t.Errorf(`Reverse("hellogdg") = %s`, r)
	}
}

// END OMIT
