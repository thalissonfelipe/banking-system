package vos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCPF(t *testing.T) {
	testCases := []struct {
		name        string
		cpf         string
		expectedCPF string
		expectedErr error
	}{
		{"cpf must be valid #1", "648.446.967-93", "64844696793", nil},
		{"cpf must be valid #2", "626.413.228-46", "62641322846", nil},
		{"cpf must be valid #3", "871.957.260-37", "87195726037", nil},
		{"cpf must be valid #4", "64844696793", "64844696793", nil},
		{"cpf must be valid #5", "62641322846", "62641322846", nil},
		{"cpf must be invalid #2", "000.000.000-00", "", ErrInvalidCPF},
		{"cpf must be invalid #1", "111.111.111-11", "", ErrInvalidCPF},
		{"cpf must be invalid #3", "222.222.222-33", "", ErrInvalidCPF},
		{"cpf must be invalid #4", "123.456.789-01", "", ErrInvalidCPF},
		{"cpf must be invalid #4", "00000000000", "", ErrInvalidCPF},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			cpf, err := NewCPF(tt.cpf)

			assert.Equal(t, tt.expectedCPF, cpf.String())
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
