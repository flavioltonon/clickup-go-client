package clickup

// Member is a team member user
type Member struct {
	User      *User `json:"user"`
	InvitedBy *User `json:"invited_by"`
}

func (m *Member) GetUser() *User {
	if m == nil {
		return nil
	}

	return m.User
}

func (m *Member) GetInvitedBy() *User {
	if m == nil {
		return nil
	}

	return m.InvitedBy
}
