package helper

import ()

func Hash(source, destination interface{}) string {
	sourceName := GetTypeName(source)
	destinationName := GetTypeName(destination)

	sourceDest := sourceName + destinationName
	return sourceDest
}
