package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSuscriber(t *testing.T) {
	// Arrange

	// Act
	result := NewSusciber()
	// Assert
	assert.NotNil(t, result)
}

func TestAddStructMapWhenElementDoesNotExist(t *testing.T) {
	//Arrange
	type struct1 struct {
	}
	type struct2 struct {
	}

	suscriber := NewSusciber()

	//Act & Assert
	assert.NotPanics(t, func() {
		_ = suscriber.AddStructMapp(struct1{}, struct2{})
	})
}

func TestAddStructMapWhenElementAlreadyExist(t *testing.T) {
	//Arrange
	type struct1 struct {
	}
	type struct2 struct {
	}

	suscriber := NewSusciber()

	//Act & Assert
	assert.Panics(t, func() {
		_ = suscriber.AddStructMapp(struct1{}, struct2{})
		_ = suscriber.AddStructMapp(struct1{}, struct2{})
	})
}

func TestCheckIfKeyExist(t *testing.T) {
	//Arrange
	suscriber := NewSusciber()
	dict := suscriber.structsMaps
	dict["test1"] = dtoMaps{}
	dict["test2"] = dtoMaps{}
	dict["test4"] = dtoMaps{}
	dict["test5"] = dtoMaps{}

	//Act & Assert

	for key := range dict {
		result := suscriber.checkIfKeyExist(key)
		assert.True(t, result)
	}
}
