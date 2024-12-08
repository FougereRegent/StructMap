package main

import ()

func hash(source, destination interface{}) string {
	sourceName := getTypeName(source)
	destinationName := getTypeName(destination)

	sourceDest := sourceName + destinationName
	return sourceDest
}
