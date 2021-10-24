package clickup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMember_GetUser(t *testing.T) {
	intValue := 1

	userValue := &User{
		ID: &intValue,
	}

	type fields struct {
		Member *Member
	}

	tests := []struct {
		name   string
		fields fields
		want   *User
	}{
		{
			name:   "If Member is nil, GetUser should return its zero value",
			fields: fields{Member: nil},
			want:   nil,
		},
		{
			name:   "If User is nil, GetUser should return its zero value",
			fields: fields{Member: &Member{User: nil}},
			want:   nil,
		},
		{
			name:   "If User has a valid value, GetUser should return it",
			fields: fields{Member: &Member{User: userValue}},
			want:   userValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Member.GetUser())
		})
	}
}

func TestMember_GetInvitedBy(t *testing.T) {
	intValue := 1

	userValue := &User{
		ID: &intValue,
	}

	type fields struct {
		Member *Member
	}

	tests := []struct {
		name   string
		fields fields
		want   *User
	}{
		{
			name:   "If Member is nil, GetInvitedBy should return its zero value",
			fields: fields{Member: nil},
			want:   nil,
		},
		{
			name:   "If InvitedBy is nil, GetInvitedBy should return its zero value",
			fields: fields{Member: &Member{InvitedBy: nil}},
			want:   nil,
		},
		{
			name:   "If InvitedBy has a valid value, GetInvitedBy should return it",
			fields: fields{Member: &Member{InvitedBy: userValue}},
			want:   userValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Member.GetInvitedBy())
		})
	}
}
