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

func (w *Webhook) GetID() string {
	if w == nil || w.ID == nil {
		return ""
	}

	return *w.ID
}

func (w *Webhook) GetUserID() int {
	if w == nil || w.UserID == nil {
		return 0
	}

	return *w.UserID
}

func (w *Webhook) GetTeamID() int {
	if w == nil || w.TeamID == nil {
		return 0
	}

	return *w.TeamID
}

func (w *Webhook) GetEndpoint() string {
	if w == nil || w.Endpoint == nil {
		return ""
	}

	return *w.Endpoint
}

func (w *Webhook) GetClientID() string {
	if w == nil || w.ClientID == nil {
		return ""
	}

	return *w.ClientID
}

func (w *Webhook) GetEvents() []WebhookEventKind {
	if w == nil {
		return nil
	}

	return w.Events
}

func (w *Webhook) GetTaskID() string {
	if w == nil || w.TaskID == nil {
		return ""
	}

	return *w.TaskID
}

func (w *Webhook) GetListID() string {
	if w == nil || w.ListID == nil {
		return ""
	}

	return *w.ListID
}

func (w *Webhook) GetFolderID() string {
	if w == nil || w.FolderID == nil {
		return ""
	}

	return *w.FolderID
}

func (w *Webhook) GetSpaceID() string {
	if w == nil || w.SpaceID == nil {
		return ""
	}

	return *w.SpaceID
}

func (w *Webhook) GetHealth() *WebhookHealth {
	if w == nil {
		return nil
	}

	return w.Health
}

func (w *Webhook) GetSecret() string {
	if w == nil || w.Secret == nil {
		return ""
	}

	return *w.Secret
}

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

func (h *WebhookHealth) GetStatus() string {
	if h == nil || h.Status == nil {
		return ""
	}

	return *h.Status
}

func (h *WebhookHealth) GetFailCount() int {
	if h == nil || h.FailCount == nil {
		return 0
	}

	return *h.FailCount
}

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

func (i *HistoryItem) GetID() string {
	if i == nil || i.ID == nil {
		return ""
	}

	return *i.ID
}

func (i *HistoryItem) GetType() int {
	if i == nil || i.Type == nil {
		return 0
	}

	return *i.Type
}

func (i *HistoryItem) GetDate() string {
	if i == nil || i.Date == nil {
		return ""
	}

	return *i.Date
}

func (i *HistoryItem) GetField() string {
	if i == nil || i.Field == nil {
		return ""
	}

	return *i.Field
}

func (i *HistoryItem) GetParentID() string {
	if i == nil || i.ParentID == nil {
		return ""
	}

	return *i.ParentID
}

func (i *HistoryItem) GetData() *HistoryItemData {
	if i == nil {
		return nil
	}

	return i.Data
}

func (i *HistoryItem) GetSource() string {
	if i == nil || i.Source == nil {
		return ""
	}

	return *i.Source
}

func (i *HistoryItem) GetUser() *User {
	if i == nil {
		return nil
	}

	return i.User
}

func (i *HistoryItem) GetBefore() *HistoryItemStatus {
	if i == nil {
		return nil
	}

	return i.Before
}

func (i *HistoryItem) GetAfter() *HistoryItemStatus {
	if i == nil {
		return nil
	}

	return i.After
}

// HistoryItemData holds additional data about a HistoryItem
type HistoryItemData struct {
	StatusType *string `json:"status_type,omitempty"`
}

func (s *HistoryItemData) GetStatusType() string {
	if s == nil || s.StatusType == nil {
		return ""
	}

	return *s.StatusType
}

// HistoryItemStatus holds data about the status of HistoryItem
type HistoryItemStatus struct {
	Status     *string `json:"status,omitempty"`
	Color      *string `json:"color,omitempty"`
	OrderIndex *int    `json:"orderindex,omitempty"`
	Type       *string `json:"type,omitempty"`
}

func (s *HistoryItemStatus) GetStatus() string {
	if s == nil || s.Status == nil {
		return ""
	}

	return *s.Status
}

func (s *HistoryItemStatus) GetColor() string {
	if s == nil || s.Color == nil {
		return ""
	}

	return *s.Color
}

func (s *HistoryItemStatus) GetOrderIndex() int {
	if s == nil || s.OrderIndex == nil {
		return 0
	}

	return *s.OrderIndex
}

func (s *HistoryItemStatus) GetType() string {
	if s == nil || s.Type == nil {
		return ""
	}

	return *s.Type
}

// WebhookEvent is an event sent by ClickUp to webhooks
type WebhookEvent struct {
	WebhookID    *string           `json:"webhook_id,omitempty"`
	Event        *WebhookEventKind `json:"event,omitempty"`
	TaskID       *string           `json:"task_id,omitempty"`
	HistoryItems []*HistoryItem    `json:"history_items,omitempty"`
}

func (e *WebhookEvent) GetWebhookID() string {
	if e == nil || e.WebhookID == nil {
		return ""
	}

	return *e.WebhookID
}

func (e *WebhookEvent) GetEvent() WebhookEventKind {
	if e == nil || e.Event == nil {
		return ""
	}

	return *e.Event
}

func (e *WebhookEvent) GetTaskID() string {
	if e == nil || e.TaskID == nil {
		return ""
	}

	return *e.TaskID
}

func (e *WebhookEvent) GetHistoryItems() []*HistoryItem {
	if e == nil {
		return nil
	}

	return e.HistoryItems
}

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
