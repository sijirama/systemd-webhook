
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
)

// WebhookPayload represents the payload structure received from the webhook
type WebhookPayload struct {
	Ref        string `json:"ref"`
	Repository struct {
		CloneURL string `json:"clone_url"`
	} `json:"repository"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received a new webhook request")

	var payload WebhookPayload
	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		log.Printf("Error decoding webhook payload: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	log.Printf("Received webhook payload: %+v", payload)

	if payload.Ref == "refs/heads/main" || payload.Ref == "ref/heads/main/dockerized" || payload.Ref == "refs/heads/release" {
		log.Printf("Triggering deployment for ref: %s", payload.Ref)
		go pullAndDeploy(payload.Repository.CloneURL)
	} else {
		log.Printf("Ignoring ref: %s", payload.Ref)
	}

	w.WriteHeader(http.StatusOK)
}

func checkhealth(w http.ResponseWriter, r *http.Request) {
	log.Println("Received a health check request")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func pullAndDeploy(cloneURL string) {
	log.Printf("Starting deployment for repository: %s", cloneURL)
	cmd := exec.Command("/bin/sh", "-c", "cd ~/repo/boymezobo && ./deploy.sh")
	err := cmd.Run()
	if err != nil {
		log.Printf("Deployment error: %v", err)
	} else {
		log.Println("Deployment successful")
	}
}

func main() {
	log.Println("Starting server...")

	http.HandleFunc("/webhook", handler)
	http.HandleFunc("/checkhealth", checkhealth)
	port := ":8080"
	log.Printf("Server listening on %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
