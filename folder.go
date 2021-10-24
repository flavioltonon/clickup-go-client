package clickup

type Folder struct {
	ID *string `json:"id,omitempty"`
}

func (f *Folder) GetID() string {
	if f == nil || f.ID == nil {
		return ""
	}

	return *f.ID
}
