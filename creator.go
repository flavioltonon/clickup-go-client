package clickup

// Creator is the data about the creator of a Task
type Creator struct {
	ID             *int    `json:"id,omitempty"`
	Username       *string `json:"username,omitempty"`
	Color          *string `json:"color,omitempty"`
	ProfilePicture *string `json:"profilePicture,omitempty"`
}

func (c *Creator) GetID() int {
	if c == nil || c.ID == nil {
		return 0
	}

	return *c.ID
}

func (c *Creator) GetUsername() string {
	if c == nil || c.Username == nil {
		return ""
	}

	return *c.Username
}

func (c *Creator) GetColor() string {
	if c == nil || c.Color == nil {
		return ""
	}

	return *c.Color
}

func (c *Creator) GetProfilePicture() string {
	if c == nil || c.ProfilePicture == nil {
		return ""
	}

	return *c.ProfilePicture
}
