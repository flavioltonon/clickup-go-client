package clickup

// User is a ClickUp user
type User struct {
	ID             *int    `json:"id,omitempty"`
	Username       *string `json:"username,omitempty"`
	Email          *string `json:"email,omitempty"`
	Color          *string `json:"color,omitempty"`
	ProfilePicture *string `json:"profilePicture,omitempty"`
	Initials       *string `json:"initials,omitempty"`
	Role           *int    `json:"role,omitempty"`
	CustomRole     *int    `json:"custom_role,omitempty"`
	LastActive     *string `json:"last_active,omitempty"`
	DateJoined     *string `json:"date_joined,omitempty"`
	DateInvited    *string `json:"date_invited,omitempty"`
}

func (u *User) GetID() int {
	if u == nil || u.ID == nil {
		return 0
	}

	return *u.ID
}

func (u *User) GetUsername() string {
	if u == nil || u.Username == nil {
		return ""
	}

	return *u.Username
}

func (u *User) GetEmail() string {
	if u == nil || u.Email == nil {
		return ""
	}

	return *u.Email
}

func (u *User) GetColor() string {
	if u == nil || u.Color == nil {
		return ""
	}

	return *u.Color
}

func (u *User) GetProfilePicture() string {
	if u == nil || u.ProfilePicture == nil {
		return ""
	}

	return *u.ProfilePicture
}

func (u *User) GetInitials() string {
	if u == nil || u.Initials == nil {
		return ""
	}

	return *u.Initials
}

func (u *User) GetRole() int {
	if u == nil || u.Role == nil {
		return 0
	}

	return *u.Role
}

func (u *User) GetCustomRole() int {
	if u == nil || u.CustomRole == nil {
		return 0
	}

	return *u.CustomRole
}

func (u *User) GetLastActive() string {
	if u == nil || u.LastActive == nil {
		return ""
	}

	return *u.LastActive
}

func (u *User) GetDateJoined() string {
	if u == nil || u.DateJoined == nil {
		return ""
	}

	return *u.DateJoined
}

func (u *User) GetDateInvited() string {
	if u == nil || u.DateInvited == nil {
		return ""
	}

	return *u.DateInvited
}
