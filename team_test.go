package clickup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeam_GetID(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Team *Team
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Team is nil, GetID should return its zero value",
			fields: fields{Team: nil},
			want:   "",
		},
		{
			name:   "If ID is nil, GetID should return its zero value",
			fields: fields{Team: &Team{ID: nil}},
			want:   "",
		},
		{
			name:   "If ID has a valid value, GetID should return it",
			fields: fields{Team: &Team{ID: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Team.GetID())
		})
	}
}

func TestTeam_GetName(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Team *Team
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Team is nil, GetName should return its zero value",
			fields: fields{Team: nil},
			want:   "",
		},
		{
			name:   "If Name is nil, GetName should return its zero value",
			fields: fields{Team: &Team{Name: nil}},
			want:   "",
		},
		{
			name:   "If Name has a valid value, GetName should return it",
			fields: fields{Team: &Team{Name: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Team.GetName())
		})
	}
}

func TestTeam_GetColor(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Team *Team
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Team is nil, GetColor should return its zero value",
			fields: fields{Team: nil},
			want:   "",
		},
		{
			name:   "If Color is nil, GetColor should return its zero value",
			fields: fields{Team: &Team{Color: nil}},
			want:   "",
		},
		{
			name:   "If Color has a valid value, GetColor should return it",
			fields: fields{Team: &Team{Color: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Team.GetColor())
		})
	}
}

func TestTeam_GetAvatar(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Team *Team
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Team is nil, GetAvatar should return its zero value",
			fields: fields{Team: nil},
			want:   "",
		},
		{
			name:   "If Avatar is nil, GetAvatar should return its zero value",
			fields: fields{Team: &Team{Avatar: nil}},
			want:   "",
		},
		{
			name:   "If Avatar has a valid value, GetAvatar should return it",
			fields: fields{Team: &Team{Avatar: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Team.GetAvatar())
		})
	}
}

func TestTeam_GetMembers(t *testing.T) {
	intValue := 1

	membersValue := []*TeamMember{
		{
			User: &User{
				ID: &intValue,
			},
		},
	}

	type fields struct {
		Team *Team
	}

	tests := []struct {
		name   string
		fields fields
		want   []*TeamMember
	}{
		{
			name:   "If Team is nil, GetMembers should return its zero value",
			fields: fields{Team: nil},
			want:   nil,
		},
		{
			name:   "If Members is nil, GetMembers should return its zero value",
			fields: fields{Team: &Team{Members: nil}},
			want:   nil,
		},
		{
			name:   "If Members has a valid value, GetMembers should return it",
			fields: fields{Team: &Team{Members: membersValue}},
			want:   membersValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Team.GetMembers())
		})
	}
}

func TestTeamMember_GetUser(t *testing.T) {
	intValue := 1

	userValue := &User{
		ID: &intValue,
	}

	type fields struct {
		TeamMember *TeamMember
	}

	tests := []struct {
		name   string
		fields fields
		want   *User
	}{
		{
			name:   "If TeamMember is nil, GetUser should return its zero value",
			fields: fields{TeamMember: nil},
			want:   nil,
		},
		{
			name:   "If User is nil, GetUser should return its zero value",
			fields: fields{TeamMember: &TeamMember{User: nil}},
			want:   nil,
		},
		{
			name:   "If User has a valid value, GetUser should return it",
			fields: fields{TeamMember: &TeamMember{User: userValue}},
			want:   userValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.TeamMember.GetUser())
		})
	}
}

func TestTeamMember_GetInvitedBy(t *testing.T) {
	intValue := 1

	userValue := &User{
		ID: &intValue,
	}

	type fields struct {
		TeamMember *TeamMember
	}

	tests := []struct {
		name   string
		fields fields
		want   *User
	}{
		{
			name:   "If TeamMember is nil, GetInvitedBy should return its zero value",
			fields: fields{TeamMember: nil},
			want:   nil,
		},
		{
			name:   "If InvitedBy is nil, GetInvitedBy should return its zero value",
			fields: fields{TeamMember: &TeamMember{InvitedBy: nil}},
			want:   nil,
		},
		{
			name:   "If InvitedBy has a valid value, GetInvitedBy should return it",
			fields: fields{TeamMember: &TeamMember{InvitedBy: userValue}},
			want:   userValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.TeamMember.GetInvitedBy())
		})
	}
}
