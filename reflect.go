package main

import (
	"reflect"
)

type Property struct {
	Type reflect.Kind
	Name string
	Tag  string
}

const tagName string = "structmap"

func getProperties(element any) []Property {
	t := reflect.TypeOf(element)
	nbProperties := t.NumField()
	result := make([]Property, nbProperties)

	for index := 0; index < nbProperties; index++ {
		field := t.Field(index)
		prop := &result[index]

		prop.Name = getName(&field)
		prop.Type = getType(&field)
		prop.Tag = getTag(&field)
	}

	return result
}

func getTypeName(element any) string {
	t := reflect.TypeOf(element)
	return t.Name()
}

func getType(field *reflect.StructField) reflect.Kind {
	t := field.Type.Kind()
	return t
}

func getName(field *reflect.StructField) string {
	name := field.Name
	return name
}

func getTag(field *reflect.StructField) string {
	tag := field.Tag.Get(tagName)
	return tag
}
