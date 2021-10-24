package clickup

type List struct {
	ID *string `json:"id,omitempty"`
}

func (l *List) GetID() string {
	if l == nil || l.ID == nil {
		return ""
	}

	return *l.ID
}
