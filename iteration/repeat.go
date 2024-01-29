package iteration

func Repeate(a string, repetedCount int) string {
	var repeated string

	for i := 0; i < repetedCount; i++ {
		repeated += a
	}
	return repeated
}
