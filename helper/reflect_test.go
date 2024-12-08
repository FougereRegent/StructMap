package helper

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

type gender int

const (
	Male gender = iota
	Female
)

func TestGetPropertiesInt(t *testing.T) {
	type typeStructIntTest struct {
		Field1    int
		Field2    int8
		Field3    int32
		Field4    int64
		PtrField1 *int
		PtrField2 *int8
		PtrField3 *int32
		PtrField4 *int64
	}

	result := GetProperties(typeStructIntTest{})

	snaps.MatchJSON(t, result)
}

func TestGetPropertiesFloat(t *testing.T) {
	type typeStructFloatTest struct {
		Field1    float32
		Field2    float64
		PtrField1 *float32
		PtrField2 *float64
	}

	result := GetProperties(typeStructFloatTest{})

	snaps.MatchJSON(t, result)
}

func TestGetPropertiesString(t *testing.T) {
	type typeStructStringTest struct {
		Field1    string
		PtrField1 *string
	}

	result := GetProperties(typeStructStringTest{})

	snaps.MatchJSON(t, result)
}

func TestGetPropertiesBoolean(t *testing.T) {
	type typeStructBooleanTest struct {
		Field1    bool
		PtrField1 *bool
	}

	result := GetProperties(typeStructBooleanTest{})

	snaps.MatchJSON(t, result)
}

func TestGetPropertiesTags(t *testing.T) {
	type userStructTest struct {
		FirstName string `structmap:"FirstName"`
		LastName  string `structmap:"LastName" json:"Test"`
		Age       int
		Address   struct {
			ZipCode int16
			Country string
			Ciry    string
		} `structmap:"Address"`
		Gender gender
	}

	result := GetProperties(userStructTest{})

	snaps.MatchJSON(t, result)
}
