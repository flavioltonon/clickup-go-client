package clickup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFolder_GetID(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Folder *Folder
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Folder is nil, GetID should return its zero value",
			fields: fields{Folder: nil},
			want:   "",
		},
		{
			name:   "If ID is nil, GetID should return its zero value",
			fields: fields{Folder: &Folder{ID: nil}},
			want:   "",
		},
		{
			name:   "If ID has a valid value, GetID should return it",
			fields: fields{Folder: &Folder{ID: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Folder.GetID())
		})
	}
}
