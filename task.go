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

func (t *Task) GetStatus() *TaskStatus {
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

func (t *Task) GetCreator() *TaskCreator {
	if t == nil {
		return nil
	}

	return t.Creator
}
func (t *Task) GetAssignees() []struct{}  { return t.Assignees }
func (t *Task) GetChecklists() []struct{} { return t.Checklists }
func (t *Task) GetTags() []struct{}       { return t.Tags }

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

func (t *Task) GetCustomFields() []*TaskCustomField { return t.CustomFields }
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

// TaskStatus is the status of a Task
type TaskStatus struct {
	Status     *string `json:"status,omitempty"`
	Color      *string `json:"color,omitempty"`
	OrderIndex *int    `json:"orderindex,omitempty"`
	Type       *string `json:"type,omitempty"`
}

func (s *TaskStatus) GetStatus() string {
	if s == nil || s.Status == nil {
		return ""
	}

	return *s.Status
}

func (s *TaskStatus) GetColor() string {
	if s == nil || s.Color == nil {
		return ""
	}

	return *s.Color
}

func (s *TaskStatus) GetOrderIndex() int {
	if s == nil || s.OrderIndex == nil {
		return 0
	}

	return *s.OrderIndex
}

func (s *TaskStatus) GetType() string {
	if s == nil || s.Type == nil {
		return ""
	}

	return *s.Type
}

// TaskCreator is the data about the creator of a Task
type TaskCreator struct {
	ID             *int    `json:"id,omitempty"`
	Username       *string `json:"username,omitempty"`
	Color          *string `json:"color,omitempty"`
	ProfilePicture *string `json:"profilePicture,omitempty"`
}

func (c *TaskCreator) GetID() int {
	if c == nil || c.ID == nil {
		return 0
	}

	return *c.ID
}

func (c *TaskCreator) GetUsername() string {
	if c == nil || c.Username == nil {
		return ""
	}

	return *c.Username
}

func (c *TaskCreator) GetColor() string {
	if c == nil || c.Color == nil {
		return ""
	}

	return *c.Color
}

func (c *TaskCreator) GetProfilePicture() string {
	if c == nil || c.ProfilePicture == nil {
		return ""
	}

	return *c.ProfilePicture
}

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

func (f *TaskCustomField) GetID() string {
	if f == nil || f.ID == nil {
		return ""
	}

	return *f.ID
}

func (f *TaskCustomField) GetName() string {
	if f == nil || f.Name == nil {
		return ""
	}

	return *f.Name
}

func (f *TaskCustomField) GetType() string {
	if f == nil || f.Type == nil {
		return ""
	}

	return *f.Type
}

func (f *TaskCustomField) GetTypeConfig() *TaskCustomFieldTypeConfig {
	if f == nil {
		return nil
	}

	return f.TypeConfig
}

func (f *TaskCustomField) GetValue() string {
	if f == nil || f.Value == nil {
		return ""
	}

	return *f.Value
}

func (f *TaskCustomField) GetDateCreated() string {
	if f == nil || f.DateCreated == nil {
		return ""
	}

	return *f.DateCreated
}

func (f *TaskCustomField) GetHideFromGuests() bool {
	if f == nil || f.HideFromGuests == nil {
		return false
	}

	return *f.HideFromGuests
}
func (f *TaskCustomField) GetRequired() bool {
	if f == nil || f.Required == nil {
		return false
	}

	return *f.Required
}

// TaskCustomFieldTypeConfig are special configs for TaskCustomFields
type TaskCustomFieldTypeConfig struct {
	IncludeGuests      *bool `json:"include_guests,omitempty"`
	IncludeTeamMembers *bool `json:"include_team_members,omitempty"`
}

func (c *TaskCustomFieldTypeConfig) GetIncludeGuests() bool {
	if c == nil || c.IncludeGuests == nil {
		return false
	}

	return *c.IncludeGuests
}
func (c *TaskCustomFieldTypeConfig) GetIncludeTeamMembers() bool {
	if c == nil || c.IncludeTeamMembers == nil {
		return false
	}

	return *c.IncludeTeamMembers
}

type List struct {
	ID *string `json:"id,omitempty"`
}

func (l *List) GetID() string {
	if l == nil || l.ID == nil {
		return ""
	}

	return *l.ID
}

type Folder struct {
	ID *string `json:"id,omitempty"`
}

func (f *Folder) GetID() string {
	if f == nil || f.ID == nil {
		return ""
	}

	return *f.ID
}

type Space struct {
	ID *string `json:"id,omitempty"`
}

func (s *Space) GetID() string {
	if s == nil || s.ID == nil {
		return ""
	}

	return *s.ID
}

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
