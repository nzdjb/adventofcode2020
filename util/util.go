package util

// Check Panic if errror.
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
