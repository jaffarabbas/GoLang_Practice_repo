package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
)

// RepoStatus represents the status of new, updated, and deleted files
type RepoStatus struct {
	New     []string `json:"new"`
	Updated []string `json:"updated"`
	Deleted []string `json:"deleted"`
}

func main() {
	// Define HTTP endpoints
	http.HandleFunc("/status", getStatusHandler)
	http.HandleFunc("/commit-push", commitPushHandler)

	// Start HTTP server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handler for /status endpoint
func getStatusHandler(w http.ResponseWriter, r *http.Request) {
	// Parse repository path from request query
	repoPath := r.URL.Query().Get("repo")
	log.Print(repoPath)
	// Clone the repository
	if err := cloneRepo(repoPath); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get status of files in repository
	status, err := getRepoStatus(repoPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Marshal status to JSON and send response
	res, err := json.Marshal(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Handler for /commit-push endpoint
func commitPushHandler(w http.ResponseWriter, r *http.Request) {
	// Parse repository path from request query
	repoPath := r.URL.Query().Get("repo")

	// Stage changes, commit, and push
	if err := commitAndPush(repoPath); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Changes committed and pushed successfully"}`))
}

// Clone the repository
func cloneRepo(repoPath string) error {
	cmd := exec.Command("git", "clone", repoPath)
	return cmd.Run()
}

// Get status of files in repository
func getRepoStatus(repoPath string) (*RepoStatus, error) {
	// TODO: Implement status checking using go-git or other Git library
	// For demonstration purposes, we'll return a mock status
	status := &RepoStatus{
		New:     []string{"new_file.txt"},
		Updated: []string{"updated_file.txt"},
		Deleted: []string{"deleted_file.txt"},
	}
	return status, nil
}

// Stage changes, commit, and push
func commitAndPush(repoPath string) error {
	// TODO: Implement committing and pushing changes using go-git or other Git library
	// For demonstration purposes, we'll just print a message
	log.Println("Changes committed and pushed successfully")
	return nil
}
