package clickup

import (
	"net/url"
	"path"
	"strconv"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

// TeamsService is a service to ClickUp teams API
type TeamsService service

// ListTeams calls ClickUp teams API to fetch teams the authenticated user belongs to
func (s *TeamsService) ListTeams() ([]*Team, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   "api.clickup.com",
		Path:   path.Join("api", "v2", "team"),
	}

	response, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}

	result := new(struct {
		Teams []*Team `json:"teams"`
	})

	if err := response.Decode(result); err != nil {
		return nil, err
	}

	return result.Teams, nil
}

// CreateWebhookRequestBody is the request body required for CreateWebhook requests
type CreateWebhookRequestBody struct {
	Endpoint string   `json:"endpoint"`
	Events   []string `json:"events"`
}

// Validate validates a CreateWebhookRequestBody
func (b CreateWebhookRequestBody) Validate() error {
	return ozzo.ValidateStruct(&b,
		ozzo.Field(&b.Endpoint, ozzo.Required),
		ozzo.Field(&b.Events, ozzo.Required),
	)
}

// CreateWebhook calls ClickUp teams API to create a Webhook with a given set of options
func (s *TeamsService) CreateWebhook(teamID int, body *CreateWebhookRequestBody) (*Webhook, error) {
	if err := body.Validate(); err != nil {
		return nil, err
	}

	url := &url.URL{
		Scheme: "https",
		Host:   "api.clickup.com",
		Path:   path.Join("api", "v2", "team", strconv.Itoa(teamID), "webhook"),
	}

	response, err := s.client.Post(url, body)
	if err != nil {
		return nil, err
	}

	result := new(struct {
		Webhook *Webhook `json:"webhook"`
	})

	if err := response.Decode(result); err != nil {
		return nil, err
	}

	return result.Webhook, nil
}

// ListWebhooks calls ClickUp teams API to list Webhooks for a given team ID
func (s *TeamsService) ListWebhooks(teamID int) ([]*Webhook, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   "api.clickup.com",
		Path:   path.Join("api", "v2", "team", strconv.Itoa(teamID), "webhook"),
	}

	response, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}

	result := new(struct {
		Webhooks []*Webhook `json:"webhooks"`
	})

	if err := response.Decode(result); err != nil {
		return nil, err
	}

	return result.Webhooks, nil
}

// DeleteWebhook calls ClickUp teams API to delete a Webhook for a given webhook ID
func (s *TeamsService) DeleteWebhook(webhookID string) error {
	url := &url.URL{
		Scheme: "https",
		Host:   "api.clickup.com",
		Path:   path.Join("api", "v2", "webhook", webhookID),
	}

	return s.client.Delete(url)
}
