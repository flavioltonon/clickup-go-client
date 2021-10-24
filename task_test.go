package clickup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask_GetID(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Task *Task
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Task is nil, GetID should return its zero value",
			fields: fields{Task: nil},
			want:   "",
		},
		{
			name:   "If ID is nil, GetID should return its zero value",
			fields: fields{Task: &Task{ID: nil}},
			want:   "",
		},
		{
			name:   "If ID has a valid value, GetID should return it",
			fields: fields{Task: &Task{ID: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Task.GetID())
		})
	}
}

func TestTask_GetCustomID(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Task *Task
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Task is nil, GetCustomID should return its zero value",
			fields: fields{Task: nil},
			want:   "",
		},
		{
			name:   "If CustomID is nil, GetCustomID should return its zero value",
			fields: fields{Task: &Task{CustomID: nil}},
			want:   "",
		},
		{
			name:   "If CustomID has a valid value, GetCustomID should return it",
			fields: fields{Task: &Task{CustomID: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Task.GetCustomID())
		})
	}
}

func TestTask_GetName(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Task *Task
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Task is nil, GetName should return its zero value",
			fields: fields{Task: nil},
			want:   "",
		},
		{
			name:   "If Name is nil, GetName should return its zero value",
			fields: fields{Task: &Task{Name: nil}},
			want:   "",
		},
		{
			name:   "If Name has a valid value, GetName should return it",
			fields: fields{Task: &Task{Name: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Task.GetName())
		})
	}
}

func TestTask_GetTextContext(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Task *Task
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Task is nil, GetTextContext should return its zero value",
			fields: fields{Task: nil},
			want:   "",
		},
		{
			name:   "If TextContext is nil, GetTextContext should return its zero value",
			fields: fields{Task: &Task{TextContext: nil}},
			want:   "",
		},
		{
			name:   "If TextContext has a valid value, GetTextContext should return it",
			fields: fields{Task: &Task{TextContext: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Task.GetTextContext())
		})
	}
}

func TestTask_GetDescription(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Task *Task
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Task is nil, GetDescription should return its zero value",
			fields: fields{Task: nil},
			want:   "",
		},
		{
			name:   "If Description is nil, GetDescription should return its zero value",
			fields: fields{Task: &Task{Description: nil}},
			want:   "",
		},
		{
			name:   "If Description has a valid value, GetDescription should return it",
			fields: fields{Task: &Task{Description: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Task.GetDescription())
		})
	}
}

func TestTask_GetStatus(t *testing.T) {
	stringValue := "1"

	taskStatusValue := Status{
		Status: &stringValue,
	}

	type fields struct {
		Task *Task
	}

	tests := []struct {
		name   string
		fields fields
		want   *Status
	}{
		{
			name:   "If Task is nil, GetStatus should return its zero value",
			fields: fields{Task: nil},
			want:   nil,
		},
		{
			name:   "If Status is nil, GetStatus should return its zero value",
			fields: fields{Task: &Task{Status: nil}},
			want:   nil,
		},
		{
			name:   "If Status has a valid value, GetStatus should return it",
			fields: fields{Task: &Task{Status: &taskStatusValue}},
			want:   &taskStatusValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Task.GetStatus())
		})
	}
}

func TestTask_GetOrderIndex(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Task *Task
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Task is nil, GetOrderIndex should return its zero value",
			fields: fields{Task: nil},
			want:   "",
		},
		{
			name:   "If OrderIndex is nil, GetOrderIndex should return its zero value",
			fields: fields{Task: &Task{OrderIndex: nil}},
			want:   "",
		},
		{
			name:   "If OrderIndex has a valid value, GetOrderIndex should return it",
			fields: fields{Task: &Task{OrderIndex: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Task.GetOrderIndex())
		})
	}
}

func TestTask_GetDateCreated(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Task *Task
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Task is nil, GetDateCreated should return its zero value",
			fields: fields{Task: nil},
			want:   "",
		},
		{
			name:   "If DateCreated is nil, GetDateCreated should return its zero value",
			fields: fields{Task: &Task{DateCreated: nil}},
			want:   "",
		},
		{
			name:   "If DateCreated has a valid value, GetDateCreated should return it",
			fields: fields{Task: &Task{DateCreated: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Task.GetDateCreated())
		})
	}
}

func TestTask_GetDateClosed(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Task *Task
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Task is nil, GetDateClosed should return its zero value",
			fields: fields{Task: nil},
			want:   "",
		},
		{
			name:   "If DateClosed is nil, GetDateClosed should return its zero value",
			fields: fields{Task: &Task{DateClosed: nil}},
			want:   "",
		},
		{
			name:   "If DateClosed has a valid value, GetDateClosed should return it",
			fields: fields{Task: &Task{DateClosed: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Task.GetDateClosed())
		})
	}
}

func TestTask_GetCreator(t *testing.T) {
	intValue := 1

	taskCreatorValue := Creator{
		ID: &intValue,
	}

	type fields struct {
		Task *Task
	}

	tests := []struct {
		name   string
		fields fields
		want   *Creator
	}{
		{
			name:   "If Task is nil, GetCreator should return its zero value",
			fields: fields{Task: nil},
			want:   nil,
		},
		{
			name:   "If Creator is nil, GetCreator should return its zero value",
			fields: fields{Task: &Task{Creator: nil}},
			want:   nil,
		},
		{
			name:   "If Creator has a valid value, GetCreator should return it",
			fields: fields{Task: &Task{Creator: &taskCreatorValue}},
			want:   &taskCreatorValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Task.GetCreator())
		})
	}
}

func TestTask_GetAssignees(t *testing.T) {
	taskAssigneesValue := []*TaskAssignee{}

	type fields struct {
		Task *Task
	}

	tests := []struct {
		name   string
		fields fields
		want   []*TaskAssignee
	}{
		{
			name:   "If Task is nil, GetAssignees should return its zero value",
			fields: fields{Task: nil},
			want:   nil,
		},
		{
			name:   "If Assignees is nil, GetAssignees should return its zero value",
			fields: fields{Task: &Task{Assignees: nil}},
			want:   nil,
		},
		{
			name:   "If Assignees has a valid value, GetAssignees should return it",
			fields: fields{Task: &Task{Assignees: taskAssigneesValue}},
			want:   taskAssigneesValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Task.GetAssignees())
		})
	}
}

func TestTask_GetChecklists(t *testing.T) {
	taskChecklistsValue := []*Checklist{}

	type fields struct {
		Task *Task
	}

	tests := []struct {
		name   string
		fields fields
		want   []*Checklist
	}{
		{
			name:   "If Task is nil, GetChecklists should return its zero value",
			fields: fields{Task: nil},
			want:   nil,
		},
		{
			name:   "If Checklists is nil, GetChecklists should return its zero value",
			fields: fields{Task: &Task{Checklists: nil}},
			want:   nil,
		},
		{
			name:   "If Checklists has a valid value, GetChecklists should return it",
			fields: fields{Task: &Task{Checklists: taskChecklistsValue}},
			want:   taskChecklistsValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Task.GetChecklists())
		})
	}
}

func TestTask_GetTags(t *testing.T) {
	taskTagsValue := []*Tag{}

	type fields struct {
		Task *Task
	}

	tests := []struct {
		name   string
		fields fields
		want   []*Tag
	}{
		{
			name:   "If Task is nil, GetTags should return its zero value",
			fields: fields{Task: nil},
			want:   nil,
		},
		{
			name:   "If Tags is nil, GetTags should return its zero value",
			fields: fields{Task: &Task{Tags: nil}},
			want:   nil,
		},
		{
			name:   "If Tags has a valid value, GetTags should return it",
			fields: fields{Task: &Task{Tags: taskTagsValue}},
			want:   taskTagsValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Task.GetTags())
		})
	}
}

func TestTask_GetParent(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Task *Task
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Task is nil, GetParent should return its zero value",
			fields: fields{Task: nil},
			want:   "",
		},
		{
			name:   "If Parent is nil, GetParent should return its zero value",
			fields: fields{Task: &Task{Parent: nil}},
			want:   "",
		},
		{
			name:   "If Parent has a valid value, GetParent should return it",
			fields: fields{Task: &Task{Parent: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Task.GetParent())
		})
	}
}

func TestTask_GetPriority(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Task *Task
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Task is nil, GetPriority should return its zero value",
			fields: fields{Task: nil},
			want:   "",
		},
		{
			name:   "If Priority is nil, GetPriority should return its zero value",
			fields: fields{Task: &Task{Priority: nil}},
			want:   "",
		},
		{
			name:   "If Priority has a valid value, GetPriority should return it",
			fields: fields{Task: &Task{Priority: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Task.GetPriority())
		})
	}
}

func TestTask_GetDueDate(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Task *Task
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Task is nil, GetDueDate should return its zero value",
			fields: fields{Task: nil},
			want:   "",
		},
		{
			name:   "If DueDate is nil, GetDueDate should return its zero value",
			fields: fields{Task: &Task{DueDate: nil}},
			want:   "",
		},
		{
			name:   "If DueDate has a valid value, GetDueDate should return it",
			fields: fields{Task: &Task{DueDate: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Task.GetDueDate())
		})
	}
}

func TestTask_GetStartDate(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Task *Task
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Task is nil, GetStartDate should return its zero value",
			fields: fields{Task: nil},
			want:   "",
		},
		{
			name:   "If StartDate is nil, GetStartDate should return its zero value",
			fields: fields{Task: &Task{StartDate: nil}},
			want:   "",
		},
		{
			name:   "If StartDate has a valid value, GetStartDate should return it",
			fields: fields{Task: &Task{StartDate: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Task.GetStartDate())
		})
	}
}

func TestTask_GetTimeEstimate(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Task *Task
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Task is nil, GetTimeEstimate should return its zero value",
			fields: fields{Task: nil},
			want:   "",
		},
		{
			name:   "If TimeEstimate is nil, GetTimeEstimate should return its zero value",
			fields: fields{Task: &Task{TimeEstimate: nil}},
			want:   "",
		},
		{
			name:   "If TimeEstimate has a valid value, GetTimeEstimate should return it",
			fields: fields{Task: &Task{TimeEstimate: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Task.GetTimeEstimate())
		})
	}
}

func TestTask_GetTimeSpent(t *testing.T) {
	intValue := 1

	type fields struct {
		Task *Task
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "If Task is nil, GetTimeSpent should return its zero value",
			fields: fields{Task: nil},
			want:   0,
		},
		{
			name:   "If TimeSpent is nil, GetTimeSpent should return its zero value",
			fields: fields{Task: &Task{TimeSpent: nil}},
			want:   0,
		},
		{
			name:   "If TimeSpent has a valid value, GetTimeSpent should return it",
			fields: fields{Task: &Task{TimeSpent: &intValue}},
			want:   intValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Task.GetTimeSpent())
		})
	}
}

func TestTask_GetCustomFields(t *testing.T) {
	stringValue := "1"

	taskCustomFieldsValue := []*CustomField{
		{
			ID: &stringValue,
		},
	}

	type fields struct {
		Task *Task
	}

	tests := []struct {
		name   string
		fields fields
		want   []*CustomField
	}{
		{
			name:   "If Task is nil, GetCustomFields should return its zero value",
			fields: fields{Task: nil},
			want:   nil,
		},
		{
			name:   "If CustomFields is nil, GetCustomFields should return its zero value",
			fields: fields{Task: &Task{CustomFields: nil}},
			want:   nil,
		},
		{
			name:   "If CustomFields has a valid value, GetCustomFields should return it",
			fields: fields{Task: &Task{CustomFields: taskCustomFieldsValue}},
			want:   taskCustomFieldsValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Task.GetCustomFields())
		})
	}
}

func TestTask_GetList(t *testing.T) {
	stringValue := "1"

	listValue := List{
		ID: &stringValue,
	}

	type fields struct {
		Task *Task
	}

	tests := []struct {
		name   string
		fields fields
		want   *List
	}{
		{
			name:   "If Task is nil, GetList should return its zero value",
			fields: fields{Task: nil},
			want:   nil,
		},
		{
			name:   "If List is nil, GetList should return its zero value",
			fields: fields{Task: &Task{List: nil}},
			want:   nil,
		},
		{
			name:   "If List has a valid value, GetList should return it",
			fields: fields{Task: &Task{List: &listValue}},
			want:   &listValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Task.GetList())
		})
	}
}

func TestTask_GetFolder(t *testing.T) {
	stringValue := "1"

	folder := Folder{
		ID: &stringValue,
	}

	type fields struct {
		Task *Task
	}

	tests := []struct {
		name   string
		fields fields
		want   *Folder
	}{
		{
			name:   "If Task is nil, GetFolder should return its zero value",
			fields: fields{Task: nil},
			want:   nil,
		},
		{
			name:   "If Folder is nil, GetFolder should return its zero value",
			fields: fields{Task: &Task{Folder: nil}},
			want:   nil,
		},
		{
			name:   "If Folder has a valid value, GetFolder should return it",
			fields: fields{Task: &Task{Folder: &folder}},
			want:   &folder,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Task.GetFolder())
		})
	}
}

func TestTask_GetSpace(t *testing.T) {
	stringValue := "1"

	spaceValue := Space{
		ID: &stringValue,
	}

	type fields struct {
		Task *Task
	}

	tests := []struct {
		name   string
		fields fields
		want   *Space
	}{
		{
			name:   "If Task is nil, GetSpace should return its zero value",
			fields: fields{Task: nil},
			want:   nil,
		},
		{
			name:   "If Space is nil, GetSpace should return its zero value",
			fields: fields{Task: &Task{Space: nil}},
			want:   nil,
		},
		{
			name:   "If Space has a valid value, GetSpace should return it",
			fields: fields{Task: &Task{Space: &spaceValue}},
			want:   &spaceValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Task.GetSpace())
		})
	}
}

func TestTask_GetURL(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Task *Task
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Task is nil, GetURL should return its zero value",
			fields: fields{Task: nil},
			want:   "",
		},
		{
			name:   "If URL is nil, GetURL should return its zero value",
			fields: fields{Task: &Task{URL: nil}},
			want:   "",
		},
		{
			name:   "If URL has a valid value, GetURL should return it",
			fields: fields{Task: &Task{URL: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Task.GetURL())
		})
	}
}
