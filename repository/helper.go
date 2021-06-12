package repository

import "strconv"

// avoidPanic() catches an error and terminates the program.
func avoidPanic(err error) {
	if err != nil {
		panic(err)
	}
}

// convertToInt() converts string datum into int datum
func convertToInt(datum string) int {
	convertedDatum, err := strconv.Atoi(datum)
	avoidPanic(err)

	return convertedDatum
}
