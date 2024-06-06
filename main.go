package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
)

// this was used as  daemon to redeploy my project on a vps

type WebhookPayload struct {
	Ref        string `json:"ref"`
	Repository struct {
		CloneURL string `json:"clone_url"`
	} `json:"repository"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var payload WebhookPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if payload.Ref == "refs/heads/main" || payload.Ref == "ref/heads/main/dockerized" || payload.Ref == "refs/heads/release" {
		go pullAndDeploy(payload.Repository.CloneURL)
	}

	w.WriteHeader(http.StatusOK)
}

func pullAndDeploy(_ string) {
	cmd := exec.Command("/bin/sh", "-c", "cd ~/repo/boymezobo && ./deploy.sh") //i used this as a daemon for redeploying my docker containers on a remote server
	err := cmd.Run()
	if err != nil {
		log.Printf("Deployment error: %v", err)
	} else {
		log.Println("Deployment successful")
	}
}

func main() {
	http.HandleFunc("/webhook", handler)

	port := ":8080"

	log.Printf("Server listening on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
