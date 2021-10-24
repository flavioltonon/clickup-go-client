package clickup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreator_GetID(t *testing.T) {
	intValue := 1

	type fields struct {
		Creator *Creator
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "If Creator is nil, GetID should return its zero value",
			fields: fields{Creator: nil},
			want:   0,
		},
		{
			name:   "If ID is nil, GetID should return its zero value",
			fields: fields{Creator: &Creator{ID: nil}},
			want:   0,
		},
		{
			name:   "If ID has a valid value, GetID should return it",
			fields: fields{Creator: &Creator{ID: &intValue}},
			want:   intValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Creator.GetID())
		})
	}
}

func TestCreator_GetUsername(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Creator *Creator
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Creator is nil, GetUsername should return its zero value",
			fields: fields{Creator: nil},
			want:   "",
		},
		{
			name:   "If Username is nil, GetUsername should return its zero value",
			fields: fields{Creator: &Creator{Username: nil}},
			want:   "",
		},
		{
			name:   "If Username has a valid value, GetUsername should return it",
			fields: fields{Creator: &Creator{Username: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Creator.GetUsername())
		})
	}
}

func TestCreator_GetColor(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Creator *Creator
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Creator is nil, GetColor should return its zero value",
			fields: fields{Creator: nil},
			want:   "",
		},
		{
			name:   "If Color is nil, GetColor should return its zero value",
			fields: fields{Creator: &Creator{Color: nil}},
			want:   "",
		},
		{
			name:   "If Color has a valid value, GetColor should return it",
			fields: fields{Creator: &Creator{Color: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Creator.GetColor())
		})
	}
}

func TestCreator_GetProfilePicture(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Creator *Creator
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Creator is nil, GetProfilePicture should return its zero value",
			fields: fields{Creator: nil},
			want:   "",
		},
		{
			name:   "If ProfilePicture is nil, GetProfilePicture should return its zero value",
			fields: fields{Creator: &Creator{ProfilePicture: nil}},
			want:   "",
		},
		{
			name:   "If ProfilePicture has a valid value, GetProfilePicture should return it",
			fields: fields{Creator: &Creator{ProfilePicture: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Creator.GetProfilePicture())
		})
	}
}
