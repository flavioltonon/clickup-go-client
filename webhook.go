package clickup

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

// Webhook is a ClickUp webhook data
type Webhook struct {
	ID       *string            `json:"id,omitempty"`
	UserID   *int               `json:"userid,omitempty"`
	TeamID   *int               `json:"team_id,omitempty"`
	Endpoint *string            `json:"endpoint,omitempty"`
	ClientID *string            `json:"client_id,omitempty"`
	Events   []WebhookEventKind `json:"events,omitempty"`
	TaskID   *string            `json:"task_id,omitempty"`
	ListID   *string            `json:"list_id,omitempty"`
	FolderID *string            `json:"folder_id,omitempty"`
	SpaceID  *string            `json:"space_id,omitempty"`
	Health   *WebhookHealth     `json:"health,omitempty"`
	Secret   *string            `json:"secret,omitempty"`
}

func (w *Webhook) GetID() string                 { return dereferenceString(w.ID) }
func (w *Webhook) GetUserID() int                { return dereferenceInt(w.UserID) }
func (w *Webhook) GetTeamID() int                { return dereferenceInt(w.TeamID) }
func (w *Webhook) GetEndpoint() string           { return dereferenceString(w.Endpoint) }
func (w *Webhook) GetClientID() string           { return dereferenceString(w.ClientID) }
func (w *Webhook) GetEvents() []WebhookEventKind { return w.Events }
func (w *Webhook) GetTaskID() string             { return dereferenceString(w.TaskID) }
func (w *Webhook) GetListID() string             { return dereferenceString(w.ListID) }
func (w *Webhook) GetFolderID() string           { return dereferenceString(w.FolderID) }
func (w *Webhook) GetSpaceID() string            { return dereferenceString(w.SpaceID) }
func (w *Webhook) GetHealth() *WebhookHealth     { return w.Health }
func (w *Webhook) GetSecret() string             { return dereferenceString(w.Secret) }

// WebhookEventKind is a webhook event kind
type WebhookEventKind string

func (e WebhookEventKind) String() string { return string(e) }

const (
	TaskCreated             WebhookEventKind = "taskCreated"
	TaskUpdated             WebhookEventKind = "taskUpdated"
	TaskDeleted             WebhookEventKind = "taskDeleted"
	TaskPriorityUpdated     WebhookEventKind = "taskPriorityUpdated"
	TaskStatusUpdated       WebhookEventKind = "taskStatusUpdated"
	TaskAssigneeUpdated     WebhookEventKind = "taskAssigneeUpdated"
	TaskDueDateUpdated      WebhookEventKind = "taskDueDateUpdated"
	TaskTagUpdated          WebhookEventKind = "taskTagUpdated"
	TaskMoved               WebhookEventKind = "taskMoved"
	TaskCommentPosted       WebhookEventKind = "taskCommentPosted"
	TaskCommentUpdated      WebhookEventKind = "taskCommentUpdated"
	TaskTimeEstimateUpdated WebhookEventKind = "taskTimeEstimateUpdated"
	TaskTimeTrackedUpdated  WebhookEventKind = "taskTimeTrackedUpdated"
	ListCreated             WebhookEventKind = "listCreated"
	ListUpdated             WebhookEventKind = "listUpdated"
	ListDeleted             WebhookEventKind = "listDeleted"
	FolderCreated           WebhookEventKind = "folderCreated"
	FolderUpdated           WebhookEventKind = "folderUpdated"
	FolderDeleted           WebhookEventKind = "folderDeleted"
	SpaceCreated            WebhookEventKind = "spaceCreated"
	SpaceUpdated            WebhookEventKind = "spaceUpdated"
	SpaceDeleted            WebhookEventKind = "spaceDeleted"
	GoalCreated             WebhookEventKind = "goalCreated"
	GoalUpdated             WebhookEventKind = "goalUpdated"
	GoalDeleted             WebhookEventKind = "goalDeleted"
	KeyResultCreated        WebhookEventKind = "keyResultCreated"
	KeyResultUpdated        WebhookEventKind = "keyResultUpdated"
	KeyResultDeleted        WebhookEventKind = "keyResultDeleted"
)

// WebhookHealth is the health status of a Webhook
type WebhookHealth struct {
	Status    *string `json:"status,omitempty"`
	FailCount *int    `json:"fail_count,omitempty"`
}

func (h *WebhookHealth) GetStatus() string { return dereferenceString(h.Status) }
func (h *WebhookHealth) GetFailCount() int { return dereferenceInt(h.FailCount) }

// HistoryItem is a item in the WebhookEvent history
type HistoryItem struct {
	ID       *string            `json:"id,omitempty"`
	Type     *int               `json:"type,omitempty"`
	Date     *string            `json:"date,omitempty"`
	Field    *string            `json:"field,omitempty"`
	ParentID *string            `json:"parent_id,omitempty"`
	Data     *HistoryItemData   `json:"data,omitempty"`
	Source   *string            `json:"source,omitempty"`
	User     *User              `json:"user,omitempty"`
	Before   *HistoryItemStatus `json:"before,omitempty"`
	After    *HistoryItemStatus `json:"after,omitempty"`
}

func (i *HistoryItem) GetID() string                 { return dereferenceString(i.ID) }
func (i *HistoryItem) GetType() int                  { return dereferenceInt(i.Type) }
func (i *HistoryItem) GetDate() string               { return dereferenceString(i.Date) }
func (i *HistoryItem) GetField() string              { return dereferenceString(i.Field) }
func (i *HistoryItem) GetParentID() string           { return dereferenceString(i.ParentID) }
func (i *HistoryItem) GetData() *HistoryItemData     { return i.Data }
func (i *HistoryItem) GetSource() string             { return dereferenceString(i.Source) }
func (i *HistoryItem) GetUser() *User                { return i.User }
func (i *HistoryItem) GetBefore() *HistoryItemStatus { return i.Before }
func (i *HistoryItem) GetAfter() *HistoryItemStatus  { return i.After }

// HistoryItemData holds additional data about a HistoryItem
type HistoryItemData struct {
	StatusType *string `json:"status_type,omitempty"`
}

func (s *HistoryItemData) GetStatusType() string { return dereferenceString(s.StatusType) }

// HistoryItemStatus holds data about the status of HistoryItem
type HistoryItemStatus struct {
	Status     *string `json:"status,omitempty"`
	Color      *string `json:"color,omitempty"`
	OrderIndex *int    `json:"orderindex,omitempty"`
	Type       *string `json:"type,omitempty"`
}

func (s *HistoryItemStatus) GetStatus() string  { return dereferenceString(s.Status) }
func (s *HistoryItemStatus) GetColor() string   { return dereferenceString(s.Color) }
func (s *HistoryItemStatus) GetOrderIndex() int { return dereferenceInt(s.OrderIndex) }
func (s *HistoryItemStatus) GetType() string    { return dereferenceString(s.Type) }

// WebhookEvent is an event sent by ClickUp to webhooks
type WebhookEvent struct {
	WebhookID    *string          `json:"webhook_id,omitempty"`
	Event        WebhookEventKind `json:"event,omitempty"`
	TaskID       *string          `json:"task_id,omitempty"`
	HistoryItems []*HistoryItem   `json:"history_items,omitempty"`
}

func (e *WebhookEvent) GetWebhookID() string            { return dereferenceString(e.WebhookID) }
func (e *WebhookEvent) GetEvent() WebhookEventKind      { return e.Event }
func (e *WebhookEvent) GetTaskID() string               { return dereferenceString(e.TaskID) }
func (e *WebhookEvent) GetHistoryItems() []*HistoryItem { return e.HistoryItems }

// ParseWebhookEvent validates the signature of a request from ClickUp's webhook and decodes its body into a
// WebhookEvent, if it is valid.
func ParseWebhookEvent(r *http.Request, webhookSecret string) (*WebhookEvent, error) {
	signature := r.Header.Get(_xSignature.String())

	if err := ozzo.Required.Validate(signature); err != nil {
		return nil, ErrInvalidRequestSignature
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	hash := hmac.New(sha256.New, []byte(webhookSecret))

	hash.Write(body)

	if hex.EncodeToString(hash.Sum(nil)) != signature {
		return nil, ErrInvalidRequestSignature
	}

	var event WebhookEvent

	if err := json.Unmarshal(body, &event); err != nil {
		return nil, err
	}

	return &event, nil
}
