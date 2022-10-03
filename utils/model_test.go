package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type baseModel struct {
	ID int
}

func TestCopy(t *testing.T) {
	type sourceType struct {
		Name    string
		Age     int
		IsValid bool
		Address []string
	}

	type destType struct {
		baseModel
		Name       string
		Age        int
		IsValid    bool
		Address    []string
		Additional string
	}

	source := sourceType{
		Name:    "name",
		Age:     18,
		IsValid: true,
		Address: []string{"address1", "address2"},
	}
	dest := destType{}
	err := Copy(&source, &dest)
	assert.Nilf(t, err, "copy failed: %v", err)
	assert.Equal(t, source.Name, dest.Name)
	assert.Equal(t, source.Age, dest.Age)
	assert.Equal(t, source.IsValid, dest.IsValid)
	assert.Equal(t, source.Address, dest.Address)
	assert.Equal(t, 0, dest.ID)
	assert.Equal(t, "", dest.Additional)

	// test exclude
	dest = destType{}
	err = Copy(&source, &dest, "Name")
	assert.Nilf(t, err, "copy failed: %v", err)
	assert.Equal(t, "", dest.Name)
	assert.Equal(t, source.Age, dest.Age)
}

// TODO: TestCopyEmbed not work
//func TestCopyEmbed(t *testing.T) {
//	type sourceType struct {
//		Name string
//		baseModel
//	}
//
//	type destType struct {
//		ID   int
//		Name string
//	}
//
//	source := sourceType{
//		Name: "name",
//		baseModel: baseModel{
//			ID: 1,
//		},
//	}
//	dest := destType{}
//	err := Copy(&source, &dest)
//	assert.Nilf(t, err, "copy failed: %v", err)
//	assert.Equal(t, source.Name, dest.Name)
//	assert.Equal(t, source.ID, dest.ID)
//}
