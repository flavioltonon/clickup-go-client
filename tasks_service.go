package clickup

import (
	"net/url"
	"path"
)

// TasksService is a service to ClickUp tasks API
type TasksService service

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
