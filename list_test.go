package clickup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_GetID(t *testing.T) {
	stringValue := "1"

	type fields struct {
		List *List
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If List is nil, GetID should return its zero value",
			fields: fields{List: nil},
			want:   "",
		},
		{
			name:   "If ID is nil, GetID should return its zero value",
			fields: fields{List: &List{ID: nil}},
			want:   "",
		},
		{
			name:   "If ID has a valid value, GetID should return it",
			fields: fields{List: &List{ID: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.List.GetID())
		})
	}
}
