package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateHash(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedErr bool
	}{
		{name: "Hash Success", input: "secret", expectedErr: false},
		{name: "Hash Failed", input: "", expectedErr: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := GenerateHash(test.input)

			if test.expectedErr {
				assert.NotNil(t, err)
				assert.Equal(t, "", result)
			} else {
				assert.NotNil(t, result)
				assert.Nil(t, err)
			}
		})
	}
}

func TestValidateHash(t *testing.T) {
	text := "secret"
	hashedText, _ := GenerateHash(text)
	tests := []struct {
		name           string
		text           string
		hashedText     string
		expectedResult bool
	}{
		{name: "Valid", text: text, hashedText: hashedText, expectedResult: true},
		{name: "Invalid Hash", text: text, hashedText: "", expectedResult: false},
		{name: "Invalid Text", text: "", hashedText: hashedText, expectedResult: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			isValid := ValidateHash(test.text, test.hashedText)
			assert.Equal(t, test.expectedResult, isValid)
		})
	}
}
