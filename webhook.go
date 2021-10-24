package clickup

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
