package clickup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebhookEvent_GetWebhookID(t *testing.T) {
	stringValue := "1"

	type fields struct {
		WebhookEvent *WebhookEvent
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If WebhookEvent is nil, GetWebhookID should return its zero value",
			fields: fields{WebhookEvent: nil},
			want:   "",
		},
		{
			name:   "If WebhookID is nil, GetWebhookID should return its zero value",
			fields: fields{WebhookEvent: &WebhookEvent{WebhookID: nil}},
			want:   "",
		},
		{
			name:   "If WebhookID has a valid value, GetWebhookID should return it",
			fields: fields{WebhookEvent: &WebhookEvent{WebhookID: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.WebhookEvent.GetWebhookID())
		})
	}
}

func TestWebhookEvent_GetEvent(t *testing.T) {
	webhookEventKind := TaskCreated

	type fields struct {
		WebhookEvent *WebhookEvent
	}

	tests := []struct {
		name   string
		fields fields
		want   WebhookEventKind
	}{
		{
			name:   "If WebhookEvent is nil, GetEvent should return its zero value",
			fields: fields{WebhookEvent: nil},
			want:   "",
		},
		{
			name:   "If Event is nil, GetEvent should return its zero value",
			fields: fields{WebhookEvent: &WebhookEvent{Event: nil}},
			want:   "",
		},
		{
			name:   "If Event has a valid value, GetEvent should return it",
			fields: fields{WebhookEvent: &WebhookEvent{Event: &webhookEventKind}},
			want:   webhookEventKind,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.WebhookEvent.GetEvent())
		})
	}
}

func TestWebhookEvent_GetTaskID(t *testing.T) {
	stringValue := "1"

	type fields struct {
		WebhookEvent *WebhookEvent
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If WebhookEvent is nil, GetTaskID should return its zero value",
			fields: fields{WebhookEvent: nil},
			want:   "",
		},
		{
			name:   "If TaskID is nil, GetTaskID should return its zero value",
			fields: fields{WebhookEvent: &WebhookEvent{TaskID: nil}},
			want:   "",
		},
		{
			name:   "If TaskID has a valid value, GetTaskID should return it",
			fields: fields{WebhookEvent: &WebhookEvent{TaskID: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.WebhookEvent.GetTaskID())
		})
	}
}

func TestWebhookEvent_GetHistoryItems(t *testing.T) {
	intValue := 1

	historyItemsValue := []*HistoryItem{
		{
			User: &User{
				ID: &intValue,
			},
		},
	}

	type fields struct {
		WebhookEvent *WebhookEvent
	}

	tests := []struct {
		name   string
		fields fields
		want   []*HistoryItem
	}{
		{
			name:   "If WebhookEvent is nil, GetHistoryItems should return its zero value",
			fields: fields{WebhookEvent: nil},
			want:   nil,
		},
		{
			name:   "If HistoryItems is nil, GetHistoryItems should return its zero value",
			fields: fields{WebhookEvent: &WebhookEvent{HistoryItems: nil}},
			want:   nil,
		},
		{
			name:   "If HistoryItems has a valid value, GetHistoryItems should return it",
			fields: fields{WebhookEvent: &WebhookEvent{HistoryItems: historyItemsValue}},
			want:   historyItemsValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.WebhookEvent.GetHistoryItems())
		})
	}
}

func TestHistoryItem_GetID(t *testing.T) {
	stringValue := "1"

	type fields struct {
		HistoryItem *HistoryItem
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If HistoryItem is nil, GetID should return its zero value",
			fields: fields{HistoryItem: nil},
			want:   "",
		},
		{
			name:   "If ID is nil, GetID should return its zero value",
			fields: fields{HistoryItem: &HistoryItem{ID: nil}},
			want:   "",
		},
		{
			name:   "If ID has a valid value, GetID should return it",
			fields: fields{HistoryItem: &HistoryItem{ID: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.HistoryItem.GetID())
		})
	}
}

func TestHistoryItem_GetType(t *testing.T) {
	intValue := 1

	type fields struct {
		HistoryItem *HistoryItem
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "If HistoryItem is nil, GetType should return its zero value",
			fields: fields{HistoryItem: nil},
			want:   0,
		},
		{
			name:   "If Type is nil, GetType should return its zero value",
			fields: fields{HistoryItem: &HistoryItem{Type: nil}},
			want:   0,
		},
		{
			name:   "If Type has a valid value, GetType should return it",
			fields: fields{HistoryItem: &HistoryItem{Type: &intValue}},
			want:   intValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.HistoryItem.GetType())
		})
	}
}

func TestHistoryItem_GetDate(t *testing.T) {
	stringValue := "1"

	type fields struct {
		HistoryItem *HistoryItem
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If HistoryItem is nil, GetDate should return its zero value",
			fields: fields{HistoryItem: nil},
			want:   "",
		},
		{
			name:   "If Date is nil, GetDate should return its zero value",
			fields: fields{HistoryItem: &HistoryItem{Date: nil}},
			want:   "",
		},
		{
			name:   "If Date has a valid value, GetDate should return it",
			fields: fields{HistoryItem: &HistoryItem{Date: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.HistoryItem.GetDate())
		})
	}
}

func TestHistoryItem_GetField(t *testing.T) {
	stringValue := "1"

	type fields struct {
		HistoryItem *HistoryItem
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If HistoryItem is nil, GetField should return its zero value",
			fields: fields{HistoryItem: nil},
			want:   "",
		},
		{
			name:   "If Field is nil, GetField should return its zero value",
			fields: fields{HistoryItem: &HistoryItem{Field: nil}},
			want:   "",
		},
		{
			name:   "If Field has a valid value, GetField should return it",
			fields: fields{HistoryItem: &HistoryItem{Field: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.HistoryItem.GetField())
		})
	}
}

func TestHistoryItem_GetParentID(t *testing.T) {
	stringValue := "1"

	type fields struct {
		HistoryItem *HistoryItem
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If HistoryItem is nil, GetParentID should return its zero value",
			fields: fields{HistoryItem: nil},
			want:   "",
		},
		{
			name:   "If ParentID is nil, GetParentID should return its zero value",
			fields: fields{HistoryItem: &HistoryItem{ParentID: nil}},
			want:   "",
		},
		{
			name:   "If ParentID has a valid value, GetParentID should return it",
			fields: fields{HistoryItem: &HistoryItem{ParentID: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.HistoryItem.GetParentID())
		})
	}
}

func TestHistoryItem_GetData(t *testing.T) {
	stringValue := "1"

	historyItemDataValue := HistoryItemData{
		StatusType: &stringValue,
	}

	type fields struct {
		HistoryItem *HistoryItem
	}

	tests := []struct {
		name   string
		fields fields
		want   *HistoryItemData
	}{
		{
			name:   "If HistoryItem is nil, GetData should return its zero value",
			fields: fields{HistoryItem: nil},
			want:   nil,
		},
		{
			name:   "If Data is nil, GetData should return its zero value",
			fields: fields{HistoryItem: &HistoryItem{Data: nil}},
			want:   nil,
		},
		{
			name:   "If Data has a valid value, GetData should return it",
			fields: fields{HistoryItem: &HistoryItem{Data: &historyItemDataValue}},
			want:   &historyItemDataValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.HistoryItem.GetData())
		})
	}
}

func TestHistoryItem_GetSource(t *testing.T) {
	stringValue := "1"

	type fields struct {
		HistoryItem *HistoryItem
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If HistoryItem is nil, GetSource should return its zero value",
			fields: fields{HistoryItem: nil},
			want:   "",
		},
		{
			name:   "If Source is nil, GetSource should return its zero value",
			fields: fields{HistoryItem: &HistoryItem{Source: nil}},
			want:   "",
		},
		{
			name:   "If Source has a valid value, GetSource should return it",
			fields: fields{HistoryItem: &HistoryItem{Source: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.HistoryItem.GetSource())
		})
	}
}

func TestHistoryItem_GetUser(t *testing.T) {
	intValue := 1

	userValue := User{
		ID: &intValue,
	}

	type fields struct {
		HistoryItem *HistoryItem
	}

	tests := []struct {
		name   string
		fields fields
		want   *User
	}{
		{
			name:   "If HistoryItem is nil, GetUser should return its zero value",
			fields: fields{HistoryItem: nil},
			want:   nil,
		},
		{
			name:   "If User is nil, GetUser should return its zero value",
			fields: fields{HistoryItem: &HistoryItem{User: nil}},
			want:   nil,
		},
		{
			name:   "If User has a valid value, GetUser should return it",
			fields: fields{HistoryItem: &HistoryItem{User: &userValue}},
			want:   &userValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.HistoryItem.GetUser())
		})
	}
}

func TestHistoryItem_GetBefore(t *testing.T) {
	stringValue := "1"

	historyItemStatusValue := Status{
		Status: &stringValue,
	}

	type fields struct {
		HistoryItem *HistoryItem
	}

	tests := []struct {
		name   string
		fields fields
		want   *Status
	}{
		{
			name:   "If HistoryItem is nil, GetBefore should return its zero value",
			fields: fields{HistoryItem: nil},
			want:   nil,
		},
		{
			name:   "If Before is nil, GetBefore should return its zero value",
			fields: fields{HistoryItem: &HistoryItem{Before: nil}},
			want:   nil,
		},
		{
			name:   "If Before has a valid value, GetBefore should return it",
			fields: fields{HistoryItem: &HistoryItem{Before: &historyItemStatusValue}},
			want:   &historyItemStatusValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.HistoryItem.GetBefore())
		})
	}
}

func TestHistoryItem_GetAfter(t *testing.T) {
	stringValue := "1"

	historyItemStatusValue := Status{
		Status: &stringValue,
	}

	type fields struct {
		HistoryItem *HistoryItem
	}

	tests := []struct {
		name   string
		fields fields
		want   *Status
	}{
		{
			name:   "If HistoryItem is nil, GetAfter should return its zero value",
			fields: fields{HistoryItem: nil},
			want:   nil,
		},
		{
			name:   "If After is nil, GetAfter should return its zero value",
			fields: fields{HistoryItem: &HistoryItem{After: nil}},
			want:   nil,
		},
		{
			name:   "If After has a valid value, GetAfter should return it",
			fields: fields{HistoryItem: &HistoryItem{After: &historyItemStatusValue}},
			want:   &historyItemStatusValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.HistoryItem.GetAfter())
		})
	}
}

func TestHistoryItemData_GetStatusType(t *testing.T) {
	stringValue := "1"

	type fields struct {
		HistoryItemData *HistoryItemData
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If HistoryItemData is nil, GetStatusType should return its zero value",
			fields: fields{HistoryItemData: nil},
			want:   "",
		},
		{
			name:   "If StatusType is nil, GetStatusType should return its zero value",
			fields: fields{HistoryItemData: &HistoryItemData{StatusType: nil}},
			want:   "",
		},
		{
			name:   "If StatusType has a valid value, GetStatusType should return it",
			fields: fields{HistoryItemData: &HistoryItemData{StatusType: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.HistoryItemData.GetStatusType())
		})
	}
}
