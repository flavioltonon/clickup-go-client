package clickup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_GetID(t *testing.T) {
	intValue := 1

	type fields struct {
		User *User
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "If User is nil, GetID should return its zero value",
			fields: fields{User: nil},
			want:   0,
		},
		{
			name:   "If ID is nil, GetID should return its zero value",
			fields: fields{User: &User{ID: nil}},
			want:   0,
		},
		{
			name:   "If ID has a valid value, GetID should return it",
			fields: fields{User: &User{ID: &intValue}},
			want:   intValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.User.GetID())
		})
	}
}

func TestUser_GetUsername(t *testing.T) {
	stringValue := "1"

	type fields struct {
		User *User
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If User is nil, GetUsername should return its zero value",
			fields: fields{User: nil},
			want:   "",
		},
		{
			name:   "If Username is nil, GetUsername should return its zero value",
			fields: fields{User: &User{Username: nil}},
			want:   "",
		},
		{
			name:   "If Username has a valid value, GetUsername should return it",
			fields: fields{User: &User{Username: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.User.GetUsername())
		})
	}
}

func TestUser_GetEmail(t *testing.T) {
	stringValue := "1"

	type fields struct {
		User *User
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If User is nil, GetEmail should return its zero value",
			fields: fields{User: nil},
			want:   "",
		},
		{
			name:   "If Email is nil, GetEmail should return its zero value",
			fields: fields{User: &User{Email: nil}},
			want:   "",
		},
		{
			name:   "If Email has a valid value, GetEmail should return it",
			fields: fields{User: &User{Email: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.User.GetEmail())
		})
	}
}

func TestUser_GetColor(t *testing.T) {
	stringValue := "1"

	type fields struct {
		User *User
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If User is nil, GetColor should return its zero value",
			fields: fields{User: nil},
			want:   "",
		},
		{
			name:   "If Color is nil, GetColor should return its zero value",
			fields: fields{User: &User{Color: nil}},
			want:   "",
		},
		{
			name:   "If Color has a valid value, GetColor should return it",
			fields: fields{User: &User{Color: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.User.GetColor())
		})
	}
}

func TestUser_GetProfilePicture(t *testing.T) {
	stringValue := "1"

	type fields struct {
		User *User
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If User is nil, GetProfilePicture should return its zero value",
			fields: fields{User: nil},
			want:   "",
		},
		{
			name:   "If ProfilePicture is nil, GetProfilePicture should return its zero value",
			fields: fields{User: &User{ProfilePicture: nil}},
			want:   "",
		},
		{
			name:   "If ProfilePicture has a valid value, GetProfilePicture should return it",
			fields: fields{User: &User{ProfilePicture: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.User.GetProfilePicture())
		})
	}
}

func TestUser_GetInitials(t *testing.T) {
	stringValue := "1"

	type fields struct {
		User *User
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If User is nil, GetInitials should return its zero value",
			fields: fields{User: nil},
			want:   "",
		},
		{
			name:   "If Initials is nil, GetInitials should return its zero value",
			fields: fields{User: &User{Initials: nil}},
			want:   "",
		},
		{
			name:   "If Initials has a valid value, GetInitials should return it",
			fields: fields{User: &User{Initials: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.User.GetInitials())
		})
	}
}

func TestUser_GetRole(t *testing.T) {
	intValue := 1

	type fields struct {
		User *User
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "If User is nil, GetRole should return its zero value",
			fields: fields{User: nil},
			want:   0,
		},
		{
			name:   "If Role is nil, GetRole should return its zero value",
			fields: fields{User: &User{Role: nil}},
			want:   0,
		},
		{
			name:   "If Role has a valid value, GetRole should return it",
			fields: fields{User: &User{Role: &intValue}},
			want:   intValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.User.GetRole())
		})
	}
}

func TestUser_GetCustomRole(t *testing.T) {
	intValue := 1

	type fields struct {
		User *User
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "If User is nil, GetCustomRole should return its zero value",
			fields: fields{User: nil},
			want:   0,
		},
		{
			name:   "If CustomRole is nil, GetCustomRole should return its zero value",
			fields: fields{User: &User{CustomRole: nil}},
			want:   0,
		},
		{
			name:   "If CustomRole has a valid value, GetCustomRole should return it",
			fields: fields{User: &User{CustomRole: &intValue}},
			want:   intValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.User.GetCustomRole())
		})
	}
}

func TestUser_GetLastActive(t *testing.T) {
	stringValue := "1"

	type fields struct {
		User *User
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If User is nil, GetLastActive should return its zero value",
			fields: fields{User: nil},
			want:   "",
		},
		{
			name:   "If LastActive is nil, GetLastActive should return its zero value",
			fields: fields{User: &User{LastActive: nil}},
			want:   "",
		},
		{
			name:   "If LastActive has a valid value, GetLastActive should return it",
			fields: fields{User: &User{LastActive: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.User.GetLastActive())
		})
	}
}

func TestUser_GetDateJoined(t *testing.T) {
	stringValue := "1"

	type fields struct {
		User *User
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If User is nil, GetDateJoined should return its zero value",
			fields: fields{User: nil},
			want:   "",
		},
		{
			name:   "If DateJoined is nil, GetDateJoined should return its zero value",
			fields: fields{User: &User{DateJoined: nil}},
			want:   "",
		},
		{
			name:   "If DateJoined has a valid value, GetDateJoined should return it",
			fields: fields{User: &User{DateJoined: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.User.GetDateJoined())
		})
	}
}

func TestUser_GetDateInvited(t *testing.T) {
	stringValue := "1"

	type fields struct {
		User *User
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If User is nil, GetDateInvited should return its zero value",
			fields: fields{User: nil},
			want:   "",
		},
		{
			name:   "If DateInvited is nil, GetDateInvited should return its zero value",
			fields: fields{User: &User{DateInvited: nil}},
			want:   "",
		},
		{
			name:   "If DateInvited has a valid value, GetDateInvited should return it",
			fields: fields{User: &User{DateInvited: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.User.GetDateInvited())
		})
	}
}
