package clickup

import (
	"net/url"
	"path"
)

// TasksService is a service to ClickUp tasks API
type TasksService service

// Task is a ClickUp task
type Task struct {
	ID           *string            `json:"id,omitempty"`
	CustomID     *string            `json:"custom_id,omitempty"`
	Name         *string            `json:"name,omitempty"`
	TextContext  *string            `json:"text_context,omitempty"`
	Description  *string            `json:"description,omitempty"`
	Status       *TaskStatus        `json:"status,omitempty"`
	OrderIndex   *string            `json:"orderindex,omitempty"`
	DateCreated  *string            `json:"date_created,omitempty"`
	DateClosed   *string            `json:"date_closed,omitempty"`
	Creator      *TaskCreator       `json:"creator,omitempty"`
	Assignees    []struct{}         `json:"assignees,omitempty"`
	Checklists   []struct{}         `json:"checklists,omitempty"`
	Tags         []struct{}         `json:"tags,omitempty"`
	Parent       *string            `json:"parent,omitempty"`
	Priority     *string            `json:"priority,omitempty"`
	DueDate      *string            `json:"due_date,omitempty"`
	StartDate    *string            `json:"start_date,omitempty"`
	TimeEstimate *string            `json:"time_estimate,omitempty"`
	TimeSpent    *int               `json:"time_spent,omitempty"`
	CustomFields []*TaskCustomField `json:"custom_fields,omitempty"`
	List         *List              `json:"list,omitempty"`
	Folder       *Folder            `json:"folder,omitempty"`
	Space        *Space             `json:"space,omitempty"`
	URL          *string            `json:"url,omitempty"`
}

func (t *Task) GetID() string                       { return dereferenceString(t.ID) }
func (t *Task) GetCustomID() string                 { return dereferenceString(t.CustomID) }
func (t *Task) GetName() string                     { return dereferenceString(t.Name) }
func (t *Task) GetTextContext() string              { return dereferenceString(t.TextContext) }
func (t *Task) GetDescription() string              { return dereferenceString(t.Description) }
func (t *Task) GetStatus() *TaskStatus              { return t.Status }
func (t *Task) GetOrderIndex() string               { return dereferenceString(t.OrderIndex) }
func (t *Task) GetDateCreated() string              { return dereferenceString(t.DateCreated) }
func (t *Task) GetDateClosed() string               { return dereferenceString(t.DateClosed) }
func (t *Task) GetCreator() *TaskCreator            { return t.Creator }
func (t *Task) GetAssignees() []struct{}            { return t.Assignees }
func (t *Task) GetChecklists() []struct{}           { return t.Checklists }
func (t *Task) GetTags() []struct{}                 { return t.Tags }
func (t *Task) GetParent() string                   { return dereferenceString(t.Parent) }
func (t *Task) GetPriority() string                 { return dereferenceString(t.Priority) }
func (t *Task) GetDueDate() string                  { return dereferenceString(t.DueDate) }
func (t *Task) GetStartDate() string                { return dereferenceString(t.StartDate) }
func (t *Task) GetTimeEstimate() string             { return dereferenceString(t.TimeEstimate) }
func (t *Task) GetTimeSpent() int                   { return dereferenceInt(t.TimeSpent) }
func (t *Task) GetCustomFields() []*TaskCustomField { return t.CustomFields }
func (t *Task) GetList() *List                      { return t.List }
func (t *Task) GetFolder() *Folder                  { return t.Folder }
func (t *Task) GetSpace() *Space                    { return t.Space }
func (t *Task) GetURL() string                      { return dereferenceString(t.URL) }

// TaskStatus is the status of a Task
type TaskStatus struct {
	Status     *string `json:"status,omitempty"`
	Color      *string `json:"color,omitempty"`
	OrderIndex *int    `json:"orderindex,omitempty"`
	Type       *string `json:"type,omitempty"`
}

func (s *TaskStatus) GetStatus() string  { return dereferenceString(s.Status) }
func (s *TaskStatus) GetColor() string   { return dereferenceString(s.Color) }
func (s *TaskStatus) GetOrderIndex() int { return dereferenceInt(s.OrderIndex) }
func (s *TaskStatus) GetType() string    { return dereferenceString(s.Type) }

// TaskCreator is the data about the creator of a Task
type TaskCreator struct {
	ID             *int    `json:"id,omitempty"`
	Username       *string `json:"username,omitempty"`
	Color          *string `json:"color,omitempty"`
	ProfilePicture *string `json:"profilePicture,omitempty"`
}

func (c *TaskCreator) GetID() int                { return dereferenceInt(c.ID) }
func (c *TaskCreator) GetUsername() string       { return dereferenceString(c.Username) }
func (c *TaskCreator) GetColor() string          { return dereferenceString(c.Color) }
func (c *TaskCreator) GetProfilePicture() string { return dereferenceString(c.ProfilePicture) }

// TaskCustomField is a customized field that can be added to Tasks
type TaskCustomField struct {
	ID             *string                    `json:"id,omitempty"`
	Name           *string                    `json:"name,omitempty"`
	Type           *string                    `json:"type,omitempty"`
	TypeConfig     *TaskCustomFieldTypeConfig `json:"type_config,omitempty"`
	Value          *string                    `json:"value,omitempty"`
	DateCreated    *string                    `json:"date_created,omitempty"`
	HideFromGuests *bool                      `json:"hide_from_guests,omitempty"`
	Required       *bool                      `json:"required,omitempty"`
}

func (f *TaskCustomField) GetID() string                             { return dereferenceString(f.ID) }
func (f *TaskCustomField) GetName() string                           { return dereferenceString(f.Name) }
func (f *TaskCustomField) GetType() string                           { return dereferenceString(f.Type) }
func (f *TaskCustomField) GetTypeConfig() *TaskCustomFieldTypeConfig { return f.TypeConfig }
func (f *TaskCustomField) GetValue() string                          { return dereferenceString(f.Value) }
func (f *TaskCustomField) GetDateCreated() string                    { return dereferenceString(f.DateCreated) }
func (f *TaskCustomField) GetHideFromGuests() bool                   { return dereferenceBool(f.HideFromGuests) }
func (f *TaskCustomField) GetRequired() bool                         { return dereferenceBool(f.Required) }

// TaskCustomFieldTypeConfig are special configs for TaskCustomFields
type TaskCustomFieldTypeConfig struct {
	IncludeGuests      *bool `json:"include_guests,omitempty"`
	IncludeTeamMembers *bool `json:"include_team_members,omitempty"`
}

func (c *TaskCustomFieldTypeConfig) GetIncludeGuests() bool { return dereferenceBool(c.IncludeGuests) }
func (c *TaskCustomFieldTypeConfig) GetIncludeTeamMembers() bool {
	return dereferenceBool(c.IncludeTeamMembers)
}

type List struct {
	ID *string `json:"id,omitempty"`
}

func (l *List) GetID() string { return dereferenceString(l.ID) }

type Folder struct {
	ID *string `json:"id,omitempty"`
}

func (f *Folder) GetID() string { return dereferenceString(f.ID) }

type Space struct {
	ID *string `json:"id,omitempty"`
}

func (s *Space) GetID() string { return dereferenceString(s.ID) }

// GetTask calls ClickUp tasks API to fetch a Task with a given ID
func (s *TasksService) GetTask(taskID string) (*Task, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   "api.clickup.com",
		Path:   path.Join("api", "v2", "task", taskID),
	}

	response, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}

	result := new(Task)

	if err := response.Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}
