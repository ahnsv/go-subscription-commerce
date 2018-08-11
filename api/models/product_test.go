package models

import (
	"testing"
)

func TestAllProducts(t *testing.T) {
	plist := AllProducts

	if len(plist) != AllProducts() {
		t.Fail()
	}

	for i, v := range plist {
		if v.name != plist[i].name {
			t.Fail()
			break
		}
	}
}
