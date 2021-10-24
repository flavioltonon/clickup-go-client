package clickup

// CustomField is a customized field that can be added to Tasks
type CustomField struct {
	ID             *string                `json:"id,omitempty"`
	Name           *string                `json:"name,omitempty"`
	Type           *string                `json:"type,omitempty"`
	TypeConfig     *CustomFieldTypeConfig `json:"type_config,omitempty"`
	Value          *string                `json:"value,omitempty"`
	DateCreated    *string                `json:"date_created,omitempty"`
	HideFromGuests *bool                  `json:"hide_from_guests,omitempty"`
	Required       *bool                  `json:"required,omitempty"`
}

func (f *CustomField) GetID() string {
	if f == nil || f.ID == nil {
		return ""
	}

	return *f.ID
}

func (f *CustomField) GetName() string {
	if f == nil || f.Name == nil {
		return ""
	}

	return *f.Name
}

func (f *CustomField) GetType() string {
	if f == nil || f.Type == nil {
		return ""
	}

	return *f.Type
}

func (f *CustomField) GetTypeConfig() *CustomFieldTypeConfig {
	if f == nil {
		return nil
	}

	return f.TypeConfig
}

func (f *CustomField) GetValue() string {
	if f == nil || f.Value == nil {
		return ""
	}

	return *f.Value
}

func (f *CustomField) GetDateCreated() string {
	if f == nil || f.DateCreated == nil {
		return ""
	}

	return *f.DateCreated
}

func (f *CustomField) GetHideFromGuests() bool {
	if f == nil || f.HideFromGuests == nil {
		return false
	}

	return *f.HideFromGuests
}
func (f *CustomField) GetRequired() bool {
	if f == nil || f.Required == nil {
		return false
	}

	return *f.Required
}

// CustomFieldTypeConfig are special configs for CustomFields
type CustomFieldTypeConfig struct {
	IncludeGuests      *bool `json:"include_guests,omitempty"`
	IncludeTeamMembers *bool `json:"include_team_members,omitempty"`
}

func (c *CustomFieldTypeConfig) GetIncludeGuests() bool {
	if c == nil || c.IncludeGuests == nil {
		return false
	}

	return *c.IncludeGuests
}
func (c *CustomFieldTypeConfig) GetIncludeTeamMembers() bool {
	if c == nil || c.IncludeTeamMembers == nil {
		return false
	}

	return *c.IncludeTeamMembers
}
