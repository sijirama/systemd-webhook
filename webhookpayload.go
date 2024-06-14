package main

type WebhookPayload struct {
	Ref        string `json:"ref"`
	Before     string `json:"before"`
	After      string `json:"after"`
	Repository struct {
		ID       int    `json:"id"`
		NodeID   string `json:"node_id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
		Owner    struct {
			Name      string `json:"name"`
			Email     string `json:"email"`
			Login     string `json:"login"`
			ID        int    `json:"id"`
			NodeID    string `json:"node_id"`
			AvatarURL string `json:"avatar_url"`
			URL       string `json:"url"`
			HTMLURL   string `json:"html_url"`
		} `json:"owner"`
		CloneURL string `json:"clone_url"`
	} `json:"repository"`
	Pusher struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"pusher"`
	Sender struct {
		Login     string `json:"login"`
		ID        int    `json:"id"`
		NodeID    string `json:"node_id"`
		AvatarURL string `json:"avatar_url"`
		URL       string `json:"url"`
		HTMLURL   string `json:"html_url"`
	} `json:"sender"`
}
