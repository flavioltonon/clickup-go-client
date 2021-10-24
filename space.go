package clickup

type Space struct {
	ID *string `json:"id,omitempty"`
}

func (s *Space) GetID() string {
	if s == nil || s.ID == nil {
		return ""
	}

	return *s.ID
}
