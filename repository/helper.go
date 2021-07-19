package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// avoidPanic() catches an error and terminates the program.
func avoidPanic(err error) {
	if err != nil {
		panic(err)
	}
}

// convertToInt() converts string into ObjectID
func convertToObjectId(datum string) primitive.ObjectID {
	convertedDatum, err := primitive.ObjectIDFromHex(datum)
	avoidPanic(err)

	return convertedDatum
}
