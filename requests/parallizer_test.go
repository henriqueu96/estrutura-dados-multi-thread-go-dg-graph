package requests

import "testing"

func newTestParallizer() Parallelizer {
	return NewParallelizer(10)
}

func Test_Parallizer_Count_ShouldReturn1(t *testing.T){
	parallizer := newTestParallizer()

	parallizer.Add("teste")

	if(parallizer.requestCount() != 1){
		t.Fail()
	}
}

func Test_Parallizer_Count_ShouldReturn0(t *testing.T){
	parallizer := newTestParallizer()

	if(parallizer.requestCount() != 0){
		t.Fail()
	}
}