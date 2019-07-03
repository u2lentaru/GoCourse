package statistic

//Average calculate average value of xs
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

//Sumxs calculate sum of xs
func Sumxs(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total
}
