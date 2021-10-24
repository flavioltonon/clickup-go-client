package clickup

// Status holds data about the status an entity
type Status struct {
	Status     *string `json:"status,omitempty"`
	Color      *string `json:"color,omitempty"`
	OrderIndex *int    `json:"orderindex,omitempty"`
	Type       *string `json:"type,omitempty"`
}

func (s *Status) GetStatus() string {
	if s == nil || s.Status == nil {
		return ""
	}

	return *s.Status
}

func (s *Status) GetColor() string {
	if s == nil || s.Color == nil {
		return ""
	}

	return *s.Color
}

func (s *Status) GetOrderIndex() int {
	if s == nil || s.OrderIndex == nil {
		return 0
	}

	return *s.OrderIndex
}

func (s *Status) GetType() string {
	if s == nil || s.Type == nil {
		return ""
	}

	return *s.Type
}
