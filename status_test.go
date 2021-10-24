package clickup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatus_GetStatus(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Status *Status
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Status is nil, GetStatus should return its zero value",
			fields: fields{Status: nil},
			want:   "",
		},
		{
			name:   "If Status is nil, GetStatus should return its zero value",
			fields: fields{Status: &Status{Status: nil}},
			want:   "",
		},
		{
			name:   "If Status has a valid value, GetStatus should return it",
			fields: fields{Status: &Status{Status: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Status.GetStatus())
		})
	}
}

func TestStatus_GetColor(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Status *Status
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Status is nil, GetColor should return its zero value",
			fields: fields{Status: nil},
			want:   "",
		},
		{
			name:   "If Color is nil, GetColor should return its zero value",
			fields: fields{Status: &Status{Color: nil}},
			want:   "",
		},
		{
			name:   "If Color has a valid value, GetColor should return it",
			fields: fields{Status: &Status{Color: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Status.GetColor())
		})
	}
}

func TestStatus_GetOrderIndex(t *testing.T) {
	intValue := 1

	type fields struct {
		Status *Status
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "If Status is nil, GetOrderIndex should return its zero value",
			fields: fields{Status: nil},
			want:   0,
		},
		{
			name:   "If OrderIndex is nil, GetOrderIndex should return its zero value",
			fields: fields{Status: &Status{OrderIndex: nil}},
			want:   0,
		},
		{
			name:   "If OrderIndex has a valid value, GetOrderIndex should return it",
			fields: fields{Status: &Status{OrderIndex: &intValue}},
			want:   intValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Status.GetOrderIndex())
		})
	}
}

func TestStatus_GetType(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Status *Status
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Status is nil, GetType should return its zero value",
			fields: fields{Status: nil},
			want:   "",
		},
		{
			name:   "If Type is nil, GetType should return its zero value",
			fields: fields{Status: &Status{Type: nil}},
			want:   "",
		},
		{
			name:   "If Type has a valid value, GetType should return it",
			fields: fields{Status: &Status{Type: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Status.GetType())
		})
	}
}
