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
