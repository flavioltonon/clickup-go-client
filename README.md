# clickup-go-client

clickup-go-client is a client for ClickUp's API written in Go.

### Usage

#### Managing webhooks
```
func main() {
	client, err := clickup.NewClient(nil)
	if err != nil {
		log.Fatal(err)
	}

	client.SetAuthorizationMethod(clickup.NewPersonalTokenAuthorization(PERSONAL_TOKEN))

	webhooks, err := client.Teams.ListWebhooks(TEAM_ID)
	if err != nil {
		log.Fatal(err)
	}

	if len(webhooks) == 0 {
		webhook, err := client.Teams.CreateWebhook(TEAM_ID, &clickup.CreateWebhookRequestBody{
			Endpoint: "https://my.endpoint.com/webhook",
			Events: []clickup.WebhookEventKind{
				clickup.TaskCreated,
				clickup.TaskStatusUpdated,
				clickup.TaskMoved,
				clickup.TaskDeleted,
				clickup.FolderCreated,
				clickup.ListCreated,
			},
		})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Webhook %s created successfully. Secret: %s\n", webhook.GetID(), webhook.GetSecret())
	}

	for _, webhook := range webhooks {
		if err := client.Teams.DeleteWebhook(webhook.GetID()); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Webhook %s deleted successfully\n", webhook.GetID())
	}
}
```

#### Handling webhook events
```
func main() {
	router := mux.NewRouter()

	parser := clickup.NewWebhookEventParser(CLICKUP_WEBHOOK_SECRET)

	router.HandleFunc("/webhook", func(rw http.ResponseWriter, r *http.Request) {
		event, err := parser.ParseWebhookEvent(r)
		if errors.Is(err, ErrInvalidRequestBody) {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		if errors.Is(err, ErrInvalidRequestSignature) {
			http.Error(rw, err.Error(), http.StatusUnauthorized)
			return
		}

		switch event.GetEvent() {
		case clickup.TaskCreated:
			fmt.Printf("Task %s created\n", event.GetTaskID())
		case clickup.TaskStatusUpdated:
			fmt.Printf("Task %s status updated\n", event.GetTaskID())

			for _, item := range event.GetHistoryItems() {
				fmt.Printf("- From: %s\n", item.GetBefore().GetStatus())
				fmt.Printf("- To: %s\n", item.GetAfter().GetStatus())
				fmt.Printf("- Author: %s\n", item.GetUser().GetUsername())
			}
		}

		rw.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":8080", router)
}
```