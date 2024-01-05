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
