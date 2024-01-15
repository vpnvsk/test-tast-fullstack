package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserUpdateValidate(t *testing.T) {
	// Test case: Valid input
	validInput := UserUpdate{Name: StringPointer("John"), Age: IntPointer(25)}
	err := validInput.Validate()

	assert.NoError(t, err, "Unexpected error for valid input")

	// Test case: All fields are nil (invalid)
	invalidInput := UserUpdate{}
	err = invalidInput.Validate()

	assert.EqualError(t, err, InvalidRequest.Error(), "Unexpected error for invalid input")

	// Test case: Age is less than 1 (invalid)
	invalidAgeInput := UserUpdate{Age: IntPointer(0)}
	err = invalidAgeInput.Validate()

	// Check if the error matches the expected error
	assert.EqualError(t, err, InvalidRequest.Error(), "Unexpected error for invalid age")

}

func StringPointer(s string) *string {
	return &s
}

func IntPointer(i int) *int {
	return &i
}
