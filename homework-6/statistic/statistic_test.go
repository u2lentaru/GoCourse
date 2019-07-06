package statistic

import "testing"

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5 got ", v)
	}
}

func TestSumxs(t *testing.T) {
	var v float64
	v = Sumxs([]float64{1, 2})
	if v != 3 {
		t.Error("Expected 3 got ", v)
	}
}

//Running tool: C:\Go\bin\go.exe test -timeout 30s GoCourse\homework-6\statistic -run ^(TestSumxs)$
//
//ok  	GoCourse/homework-6/statistic	(cached)
//Success: Tests passed.
