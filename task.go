package clickup

// Task is a ClickUp task
type Task struct {
	ID           *string         `json:"id,omitempty"`
	CustomID     *string         `json:"custom_id,omitempty"`
	Name         *string         `json:"name,omitempty"`
	TextContext  *string         `json:"text_context,omitempty"`
	Description  *string         `json:"description,omitempty"`
	Status       *Status         `json:"status,omitempty"`
	OrderIndex   *string         `json:"orderindex,omitempty"`
	DateCreated  *string         `json:"date_created,omitempty"`
	DateClosed   *string         `json:"date_closed,omitempty"`
	Creator      *Creator        `json:"creator,omitempty"`
	Assignees    []*TaskAssignee `json:"assignees,omitempty"`
	Checklists   []*Checklist    `json:"checklists,omitempty"`
	Tags         []*Tag          `json:"tags,omitempty"`
	Parent       *string         `json:"parent,omitempty"`
	Priority     *string         `json:"priority,omitempty"`
	DueDate      *string         `json:"due_date,omitempty"`
	StartDate    *string         `json:"start_date,omitempty"`
	TimeEstimate *string         `json:"time_estimate,omitempty"`
	TimeSpent    *int            `json:"time_spent,omitempty"`
	CustomFields []*CustomField  `json:"custom_fields,omitempty"`
	List         *List           `json:"list,omitempty"`
	Folder       *Folder         `json:"folder,omitempty"`
	Space        *Space          `json:"space,omitempty"`
	URL          *string         `json:"url,omitempty"`
}

func (t *Task) GetID() string {
	if t == nil || t.ID == nil {
		return ""
	}

	return *t.ID
}

func (t *Task) GetCustomID() string {
	if t == nil || t.CustomID == nil {
		return ""
	}

	return *t.CustomID
}

func (t *Task) GetName() string {
	if t == nil || t.Name == nil {
		return ""
	}

	return *t.Name
}

func (t *Task) GetTextContext() string {
	if t == nil || t.TextContext == nil {
		return ""
	}

	return *t.TextContext
}

func (t *Task) GetDescription() string {
	if t == nil || t.Description == nil {
		return ""
	}

	return *t.Description
}

func (t *Task) GetStatus() *Status {
	if t == nil {
		return nil
	}

	return t.Status
}
func (t *Task) GetOrderIndex() string {
	if t == nil || t.OrderIndex == nil {
		return ""
	}

	return *t.OrderIndex
}

func (t *Task) GetDateCreated() string {
	if t == nil || t.DateCreated == nil {
		return ""
	}

	return *t.DateCreated
}

func (t *Task) GetDateClosed() string {
	if t == nil || t.DateClosed == nil {
		return ""
	}

	return *t.DateClosed
}

func (t *Task) GetCreator() *Creator {
	if t == nil {
		return nil
	}

	return t.Creator
}
func (t *Task) GetAssignees() []*TaskAssignee {
	if t == nil {
		return nil
	}

	return t.Assignees
}

func (t *Task) GetChecklists() []*Checklist {
	if t == nil {
		return nil
	}

	return t.Checklists
}

func (t *Task) GetTags() []*Tag {
	if t == nil {
		return nil
	}

	return t.Tags
}

func (t *Task) GetParent() string {
	if t == nil || t.Parent == nil {
		return ""
	}

	return *t.Parent
}

func (t *Task) GetPriority() string {
	if t == nil || t.Priority == nil {
		return ""
	}

	return *t.Priority
}

func (t *Task) GetDueDate() string {
	if t == nil || t.DueDate == nil {
		return ""
	}

	return *t.DueDate
}

func (t *Task) GetStartDate() string {
	if t == nil || t.StartDate == nil {
		return ""
	}

	return *t.StartDate
}

func (t *Task) GetTimeEstimate() string {
	if t == nil || t.TimeEstimate == nil {
		return ""
	}

	return *t.TimeEstimate
}

func (t *Task) GetTimeSpent() int {
	if t == nil || t.TimeSpent == nil {
		return 0
	}

	return *t.TimeSpent
}

func (t *Task) GetCustomFields() []*CustomField {
	if t == nil {
		return nil
	}

	return t.CustomFields
}

func (t *Task) GetList() *List {
	if t == nil {
		return nil
	}

	return t.List
}
func (t *Task) GetFolder() *Folder {
	if t == nil {
		return nil
	}

	return t.Folder
}
func (t *Task) GetSpace() *Space {
	if t == nil {
		return nil
	}

	return t.Space
}
func (t *Task) GetURL() string {
	if t == nil || t.URL == nil {
		return ""
	}

	return *t.URL
}

// TaskAssignee holds the data of an assignee
type TaskAssignee struct{}
