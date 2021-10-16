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

func (u *User) GetID() int                { return dereferenceInt(u.ID) }
func (u *User) GetUsername() string       { return dereferenceString(u.Username) }
func (u *User) GetEmail() string          { return dereferenceString(u.Email) }
func (u *User) GetColor() string          { return dereferenceString(u.Color) }
func (u *User) GetProfilePicture() string { return dereferenceString(u.ProfilePicture) }
func (u *User) GetInitials() string       { return dereferenceString(u.Initials) }
func (u *User) GetRole() int              { return dereferenceInt(u.Role) }
func (u *User) GetCustomRole() int        { return dereferenceInt(u.CustomRole) }
func (u *User) GetLastActive() string     { return dereferenceString(u.LastActive) }
func (u *User) GetDateJoined() string     { return dereferenceString(u.DateJoined) }
func (u *User) GetDateInvited() string    { return dereferenceString(u.DateInvited) }
