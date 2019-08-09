package boolGenerator

import (
	"testing"
	"time"
)

func Test_boolGenerator_New_ShouldReturnTrue(t *testing.T){
	var time = time.Date(2000,10,10,10,10,10,10,time.Local)
	random := New(time)

	if !random {
		t.Fail()
	}
}

func Test_boolGenerator_New_ShouldReturnFalse(t *testing.T){
	var time = time.Date(2000,10,10,10,10,10,11,time.Local)
	random := New(time)

	if random {
		t.Fail()
	}
}