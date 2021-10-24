package clickup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpace_GetID(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Space *Space
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Space is nil, GetID should return its zero value",
			fields: fields{Space: nil},
			want:   "",
		},
		{
			name:   "If ID is nil, GetID should return its zero value",
			fields: fields{Space: &Space{ID: nil}},
			want:   "",
		},
		{
			name:   "If ID has a valid value, GetID should return it",
			fields: fields{Space: &Space{ID: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Space.GetID())
		})
	}
}
