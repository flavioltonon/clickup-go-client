package clickup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomField_GetID(t *testing.T) {
	stringValue := "1"

	type fields struct {
		CustomField *CustomField
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If CustomField is nil, GetID should return its zero value",
			fields: fields{CustomField: nil},
			want:   "",
		},
		{
			name:   "If ID is nil, GetID should return its zero value",
			fields: fields{CustomField: &CustomField{ID: nil}},
			want:   "",
		},
		{
			name:   "If ID has a valid value, GetID should return it",
			fields: fields{CustomField: &CustomField{ID: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.CustomField.GetID())
		})
	}
}

func TestCustomField_GetName(t *testing.T) {
	stringValue := "1"

	type fields struct {
		CustomField *CustomField
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If CustomField is nil, GetName should return its zero value",
			fields: fields{CustomField: nil},
			want:   "",
		},
		{
			name:   "If Name is nil, GetName should return its zero value",
			fields: fields{CustomField: &CustomField{Name: nil}},
			want:   "",
		},
		{
			name:   "If Name has a valid value, GetName should return it",
			fields: fields{CustomField: &CustomField{Name: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.CustomField.GetName())
		})
	}
}

func TestCustomField_GetType(t *testing.T) {
	stringValue := "1"

	type fields struct {
		CustomField *CustomField
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If CustomField is nil, GetType should return its zero value",
			fields: fields{CustomField: nil},
			want:   "",
		},
		{
			name:   "If Type is nil, GetType should return its zero value",
			fields: fields{CustomField: &CustomField{Type: nil}},
			want:   "",
		},
		{
			name:   "If Type has a valid value, GetType should return it",
			fields: fields{CustomField: &CustomField{Type: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.CustomField.GetType())
		})
	}
}

func TestCustomField_GetTypeConfig(t *testing.T) {
	boolValue := true

	customFieldTypeConfigValue := CustomFieldTypeConfig{
		IncludeGuests: &boolValue,
	}

	type fields struct {
		CustomField *CustomField
	}

	tests := []struct {
		name   string
		fields fields
		want   *CustomFieldTypeConfig
	}{
		{
			name:   "If CustomField is nil, GetTypeConfig should return its zero value",
			fields: fields{CustomField: nil},
			want:   nil,
		},
		{
			name:   "If TypeConfig is nil, GetTypeConfig should return its zero value",
			fields: fields{CustomField: &CustomField{TypeConfig: nil}},
			want:   nil,
		},
		{
			name:   "If TypeConfig has a valid value, GetTypeConfig should return it",
			fields: fields{CustomField: &CustomField{TypeConfig: &customFieldTypeConfigValue}},
			want:   &customFieldTypeConfigValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.CustomField.GetTypeConfig())
		})
	}
}

func TestCustomField_GetValue(t *testing.T) {
	stringValue := "1"

	type fields struct {
		CustomField *CustomField
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If CustomField is nil, GetValue should return its zero value",
			fields: fields{CustomField: nil},
			want:   "",
		},
		{
			name:   "If Value is nil, GetValue should return its zero value",
			fields: fields{CustomField: &CustomField{Value: nil}},
			want:   "",
		},
		{
			name:   "If Value has a valid value, GetValue should return it",
			fields: fields{CustomField: &CustomField{Value: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.CustomField.GetValue())
		})
	}
}

func TestCustomField_GetDateCreated(t *testing.T) {
	stringValue := "1"

	type fields struct {
		CustomField *CustomField
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If CustomField is nil, GetDateCreated should return its zero value",
			fields: fields{CustomField: nil},
			want:   "",
		},
		{
			name:   "If DateCreated is nil, GetDateCreated should return its zero value",
			fields: fields{CustomField: &CustomField{DateCreated: nil}},
			want:   "",
		},
		{
			name:   "If DateCreated has a valid value, GetDateCreated should return it",
			fields: fields{CustomField: &CustomField{DateCreated: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.CustomField.GetDateCreated())
		})
	}
}

func TestCustomField_GetHideFromGuests(t *testing.T) {
	boolValue := true

	type fields struct {
		CustomField *CustomField
	}

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "If CustomField is nil, GetHideFromGuests should return its zero value",
			fields: fields{CustomField: nil},
			want:   false,
		},
		{
			name:   "If HideFromGuests is nil, GetHideFromGuests should return its zero value",
			fields: fields{CustomField: &CustomField{HideFromGuests: nil}},
			want:   false,
		},
		{
			name:   "If HideFromGuests has a valid value, GetHideFromGuests should return it",
			fields: fields{CustomField: &CustomField{HideFromGuests: &boolValue}},
			want:   boolValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.CustomField.GetHideFromGuests())
		})
	}
}

func TestCustomField_GetRequired(t *testing.T) {
	boolValue := true

	type fields struct {
		CustomField *CustomField
	}

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "If CustomField is nil, GetRequired should return its zero value",
			fields: fields{CustomField: nil},
			want:   false,
		},
		{
			name:   "If Required is nil, GetRequired should return its zero value",
			fields: fields{CustomField: &CustomField{Required: nil}},
			want:   false,
		},
		{
			name:   "If Required has a valid value, GetRequired should return it",
			fields: fields{CustomField: &CustomField{Required: &boolValue}},
			want:   boolValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.CustomField.GetRequired())
		})
	}
}

func TestCustomFieldTypeConfig_GetIncludeGuests(t *testing.T) {
	boolValue := true

	type fields struct {
		CustomFieldTypeConfig *CustomFieldTypeConfig
	}

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "If CustomFieldTypeConfig is nil, GetIncludeGuests should return its zero value",
			fields: fields{CustomFieldTypeConfig: nil},
			want:   false,
		},
		{
			name:   "If IncludeGuests is nil, GetIncludeGuests should return its zero value",
			fields: fields{CustomFieldTypeConfig: &CustomFieldTypeConfig{IncludeGuests: nil}},
			want:   false,
		},
		{
			name:   "If IncludeGuests has a valid value, GetIncludeGuests should return it",
			fields: fields{CustomFieldTypeConfig: &CustomFieldTypeConfig{IncludeGuests: &boolValue}},
			want:   boolValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.CustomFieldTypeConfig.GetIncludeGuests())
		})
	}
}

func TestCustomFieldTypeConfig_GetIncludeTeamMembers(t *testing.T) {
	boolValue := true

	type fields struct {
		CustomFieldTypeConfig *CustomFieldTypeConfig
	}

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "If CustomFieldTypeConfig is nil, GetIncludeTeamMembers should return its zero value",
			fields: fields{CustomFieldTypeConfig: nil},
			want:   false,
		},
		{
			name:   "If IncludeTeamMembers is nil, GetIncludeTeamMembers should return its zero value",
			fields: fields{CustomFieldTypeConfig: &CustomFieldTypeConfig{IncludeTeamMembers: nil}},
			want:   false,
		},
		{
			name:   "If IncludeTeamMembers has a valid value, GetIncludeTeamMembers should return it",
			fields: fields{CustomFieldTypeConfig: &CustomFieldTypeConfig{IncludeTeamMembers: &boolValue}},
			want:   boolValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.CustomFieldTypeConfig.GetIncludeTeamMembers())
		})
	}
}
