package model

import (
	"testing"
)

func TestCheckDay(t *testing.T) {
	var x BlackListedDays = NewBlackListedDays()
	x.AddDay(0)
	r := x.CheckDay(0)
	if !r {
		t.Errorf("x = xxxxxxx1; x.CheckDay(0) = %t; want true", r)
	}
	r = x.CheckDay(1)
	if r {
		t.Errorf("x = xxxxxxx1; x.CheckDay(1) = %t; want false", r)
	}
	x.AddDay(4)
	r = x.CheckDay(4)
	if !r {
		t.Errorf("x = xxx1xxxx; x.CheckDay(4) = %t; want true", r)
	}
	r = x.CheckDay(2)
	if r {
		t.Errorf("x = xxx1xxxx; x.CheckDay(2) = %t; want false", r)
	}
}
