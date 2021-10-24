package clickup

// Team is a ClickUp team
type Team struct {
	ID      *string       `json:"id"`
	Name    *string       `json:"name"`
	Color   *string       `json:"color"`
	Avatar  *string       `json:"avatar"`
	Members []*TeamMember `json:"members"`
}

func (t *Team) GetID() string {
	if t == nil || t.ID == nil {
		return ""
	}

	return *t.ID
}

func (t *Team) GetName() string {
	if t == nil || t.Name == nil {
		return ""
	}

	return *t.Name
}

func (t *Team) GetColor() string {
	if t == nil || t.Color == nil {
		return ""
	}

	return *t.Color
}

func (t *Team) GetAvatar() string {
	if t == nil || t.Avatar == nil {
		return ""
	}

	return *t.Avatar
}

func (t *Team) GetMembers() []*TeamMember {
	if t == nil {
		return nil
	}

	return t.Members
}

// TeamMember is a team member user
type TeamMember struct {
	User      *User `json:"user"`
	InvitedBy *User `json:"invited_by"`
}

func (m *TeamMember) GetUser() *User {
	if m == nil {
		return nil
	}

	return m.User
}

func (m *TeamMember) GetInvitedBy() *User {
	if m == nil {
		return nil
	}

	return m.InvitedBy
}
