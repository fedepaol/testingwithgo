package testingwithgo

import "testing"

func TestReverse(t *testing.T) {
	r := Reverse("hellogdg")
	if r != "gdgolleh" {
		t.Errorf(`Reverse("hellogdg") = %s`, r)
	}
}
