package clickup

// Team is a ClickUp team
type Team struct {
	ID      *string   `json:"id"`
	Name    *string   `json:"name"`
	Color   *string   `json:"color"`
	Avatar  *string   `json:"avatar"`
	Members []*Member `json:"members"`
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

func (t *Team) GetMembers() []*Member {
	if t == nil {
		return nil
	}

	return t.Members
}
