package clickup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebhook_GetID(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Webhook *Webhook
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Webhook is nil, GetID should return its zero value",
			fields: fields{Webhook: nil},
			want:   "",
		},
		{
			name:   "If ID is nil, GetID should return its zero value",
			fields: fields{Webhook: &Webhook{ID: nil}},
			want:   "",
		},
		{
			name:   "If ID has a valid value, GetID should return it",
			fields: fields{Webhook: &Webhook{ID: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Webhook.GetID())
		})
	}
}

func TestWebhook_GetUserID(t *testing.T) {
	intValue := 1

	type fields struct {
		Webhook *Webhook
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "If Webhook is nil, GetUserID should return its zero value",
			fields: fields{Webhook: nil},
			want:   0,
		},
		{
			name:   "If UserID is nil, GetUserID should return its zero value",
			fields: fields{Webhook: &Webhook{UserID: nil}},
			want:   0,
		},
		{
			name:   "If UserID has a valid value, GetUserID should return it",
			fields: fields{Webhook: &Webhook{UserID: &intValue}},
			want:   intValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Webhook.GetUserID())
		})
	}
}

func TestWebhook_GetTeamID(t *testing.T) {
	intValue := 1

	type fields struct {
		Webhook *Webhook
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "If Webhook is nil, GetTeamID should return its zero value",
			fields: fields{Webhook: nil},
			want:   0,
		},
		{
			name:   "If TeamID is nil, GetTeamID should return its zero value",
			fields: fields{Webhook: &Webhook{TeamID: nil}},
			want:   0,
		},
		{
			name:   "If TeamID has a valid value, GetTeamID should return it",
			fields: fields{Webhook: &Webhook{TeamID: &intValue}},
			want:   intValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Webhook.GetTeamID())
		})
	}
}

func TestWebhook_GetEndpoint(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Webhook *Webhook
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Webhook is nil, GetEndpoint should return its zero value",
			fields: fields{Webhook: nil},
			want:   "",
		},
		{
			name:   "If Endpoint is nil, GetEndpoint should return its zero value",
			fields: fields{Webhook: &Webhook{Endpoint: nil}},
			want:   "",
		},
		{
			name:   "If Endpoint has a valid value, GetEndpoint should return it",
			fields: fields{Webhook: &Webhook{Endpoint: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Webhook.GetEndpoint())
		})
	}
}

func TestWebhook_GetClientID(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Webhook *Webhook
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Webhook is nil, GetClientID should return its zero value",
			fields: fields{Webhook: nil},
			want:   "",
		},
		{
			name:   "If ClientID is nil, GetClientID should return its zero value",
			fields: fields{Webhook: &Webhook{ClientID: nil}},
			want:   "",
		},
		{
			name:   "If ClientID has a valid value, GetClientID should return it",
			fields: fields{Webhook: &Webhook{ClientID: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Webhook.GetClientID())
		})
	}
}

func TestWebhook_GetEvents(t *testing.T) {
	eventsValue := []WebhookEventKind{
		TaskCreated,
	}

	type fields struct {
		Webhook *Webhook
	}

	tests := []struct {
		name   string
		fields fields
		want   []WebhookEventKind
	}{
		{
			name:   "If Webhook is nil, GetEvents should return its zero value",
			fields: fields{Webhook: nil},
			want:   nil,
		},
		{
			name:   "If Events is nil, GetEvents should return its zero value",
			fields: fields{Webhook: &Webhook{Events: nil}},
			want:   nil,
		},
		{
			name:   "If Events has a valid value, GetEvents should return it",
			fields: fields{Webhook: &Webhook{Events: eventsValue}},
			want:   eventsValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Webhook.GetEvents())
		})
	}
}

func TestWebhook_GetTaskID(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Webhook *Webhook
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Webhook is nil, GetTaskID should return its zero value",
			fields: fields{Webhook: nil},
			want:   "",
		},
		{
			name:   "If TaskID is nil, GetTaskID should return its zero value",
			fields: fields{Webhook: &Webhook{TaskID: nil}},
			want:   "",
		},
		{
			name:   "If TaskID has a valid value, GetTaskID should return it",
			fields: fields{Webhook: &Webhook{TaskID: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Webhook.GetTaskID())
		})
	}
}

func TestWebhook_GetListID(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Webhook *Webhook
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Webhook is nil, GetListID should return its zero value",
			fields: fields{Webhook: nil},
			want:   "",
		},
		{
			name:   "If ListID is nil, GetListID should return its zero value",
			fields: fields{Webhook: &Webhook{ListID: nil}},
			want:   "",
		},
		{
			name:   "If ListID has a valid value, GetListID should return it",
			fields: fields{Webhook: &Webhook{ListID: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Webhook.GetListID())
		})
	}
}

func TestWebhook_GetFolderID(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Webhook *Webhook
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Webhook is nil, GetFolderID should return its zero value",
			fields: fields{Webhook: nil},
			want:   "",
		},
		{
			name:   "If FolderID is nil, GetFolderID should return its zero value",
			fields: fields{Webhook: &Webhook{FolderID: nil}},
			want:   "",
		},
		{
			name:   "If FolderID has a valid value, GetFolderID should return it",
			fields: fields{Webhook: &Webhook{FolderID: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Webhook.GetFolderID())
		})
	}
}

func TestWebhook_GetSpaceID(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Webhook *Webhook
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Webhook is nil, GetSpaceID should return its zero value",
			fields: fields{Webhook: nil},
			want:   "",
		},
		{
			name:   "If SpaceID is nil, GetSpaceID should return its zero value",
			fields: fields{Webhook: &Webhook{SpaceID: nil}},
			want:   "",
		},
		{
			name:   "If SpaceID has a valid value, GetSpaceID should return it",
			fields: fields{Webhook: &Webhook{SpaceID: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Webhook.GetSpaceID())
		})
	}
}

func TestWebhook_GetHealth(t *testing.T) {
	stringValue := "1"

	healthValue := &WebhookHealth{
		Status: &stringValue,
	}

	type fields struct {
		Webhook *Webhook
	}

	tests := []struct {
		name   string
		fields fields
		want   *WebhookHealth
	}{
		{
			name:   "If Webhook is nil, GetHealth should return its zero value",
			fields: fields{Webhook: nil},
			want:   nil,
		},
		{
			name:   "If Health is nil, GetHealth should return its zero value",
			fields: fields{Webhook: &Webhook{Health: nil}},
			want:   nil,
		},
		{
			name:   "If Health has a valid value, GetHealth should return it",
			fields: fields{Webhook: &Webhook{Health: healthValue}},
			want:   healthValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Webhook.GetHealth())
		})
	}
}

func TestWebhook_GetSecret(t *testing.T) {
	stringValue := "1"

	type fields struct {
		Webhook *Webhook
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If Webhook is nil, GetSecret should return its zero value",
			fields: fields{Webhook: nil},
			want:   "",
		},
		{
			name:   "If Secret is nil, GetSecret should return its zero value",
			fields: fields{Webhook: &Webhook{Secret: nil}},
			want:   "",
		},
		{
			name:   "If Secret has a valid value, GetSecret should return it",
			fields: fields{Webhook: &Webhook{Secret: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.Webhook.GetSecret())
		})
	}
}

func TestWebhookHealth_GetStatus(t *testing.T) {
	stringValue := "1"

	type fields struct {
		WebhookHealth *WebhookHealth
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "If WebhookHealth is nil, GetStatus should return its zero value",
			fields: fields{WebhookHealth: nil},
			want:   "",
		},
		{
			name:   "If Status is nil, GetStatus should return its zero value",
			fields: fields{WebhookHealth: &WebhookHealth{Status: nil}},
			want:   "",
		},
		{
			name:   "If Status has a valid value, GetStatus should return it",
			fields: fields{WebhookHealth: &WebhookHealth{Status: &stringValue}},
			want:   stringValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.WebhookHealth.GetStatus())
		})
	}
}

func TestWebhookHealth_GetFailCount(t *testing.T) {
	intValue := 1

	type fields struct {
		WebhookHealth *WebhookHealth
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "If WebhookHealth is nil, GetFailCount should return its zero value",
			fields: fields{WebhookHealth: nil},
			want:   0,
		},
		{
			name:   "If FailCount is nil, GetFailCount should return its zero value",
			fields: fields{WebhookHealth: &WebhookHealth{FailCount: nil}},
			want:   0,
		},
		{
			name:   "If FailCount has a valid value, GetFailCount should return it",
			fields: fields{WebhookHealth: &WebhookHealth{FailCount: &intValue}},
			want:   intValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.WebhookHealth.GetFailCount())
		})
	}
}
