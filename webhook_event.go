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

// WebhookEvent is an event sent by ClickUp to webhooks
type WebhookEvent struct {
	WebhookID    *string           `json:"webhook_id,omitempty"`
	Event        *WebhookEventKind `json:"event,omitempty"`
	TaskID       *string           `json:"task_id,omitempty"`
	HistoryItems []*HistoryItem    `json:"history_items,omitempty"`
}

// WebhookEventParser is a parser for WebhookEvents
type WebhookEventParser struct {
	secret string
}

// NewWebhookEventParser returns a new WebhookEventParser with a given secret
func NewWebhookEventParser(secret string) *WebhookEventParser {
	return &WebhookEventParser{secret: secret}
}

// ParseWebhookEvent validates the signature of a request from ClickUp's webhook and decodes its body into a
// WebhookEvent, if it is valid.
//
// - If the signature is not valid, an ErrInvalidRequestSignature will be returned.
// - If the request body cannot be read or decoded, an ErrInvalidRequestBody will be returned
func (p *WebhookEventParser) ParseWebhookEvent(r *http.Request) (*WebhookEvent, error) {
	signature := r.Header.Get(header_xSignature)

	if err := ozzo.Required.Validate(signature); err != nil {
		return nil, ErrInvalidRequestSignature
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, ErrInvalidRequestBody
	}

	hash := hmac.New(sha256.New, []byte(p.secret))

	hash.Write(body)

	if hex.EncodeToString(hash.Sum(nil)) != signature {
		return nil, ErrInvalidRequestSignature
	}

	event := new(WebhookEvent)

	if err := json.Unmarshal(body, event); err != nil {
		return nil, ErrInvalidRequestBody
	}

	return event, nil
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

// HistoryItem is a item in the WebhookEvent history
type HistoryItem struct {
	ID       *string          `json:"id,omitempty"`
	Type     *int             `json:"type,omitempty"`
	Date     *string          `json:"date,omitempty"`
	Field    *string          `json:"field,omitempty"`
	ParentID *string          `json:"parent_id,omitempty"`
	Data     *HistoryItemData `json:"data,omitempty"`
	Source   *string          `json:"source,omitempty"`
	User     *User            `json:"user,omitempty"`
	Before   *Status          `json:"before,omitempty"`
	After    *Status          `json:"after,omitempty"`
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

func (i *HistoryItem) GetBefore() *Status {
	if i == nil {
		return nil
	}

	return i.Before
}

func (i *HistoryItem) GetAfter() *Status {
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
