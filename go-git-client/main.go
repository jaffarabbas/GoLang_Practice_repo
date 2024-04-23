package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

func main() {
	var dirpath = "J:\\GitHub"
	selectedDirectory := getSelectedDirectory(dirpath)
	fmt.Printf("Selected directory: %s\n", selectedDirectory)
	// Open the repository
	r, err := git.PlainOpen(selectedDirectory)
	if err != nil {
		fmt.Printf("Error opening repository: %v\n", err)
		os.Exit(1)
	}

	// Get the worktree
	w, err := r.Worktree()
	if err != nil {
		fmt.Printf("Error getting worktree: %v\n", err)
		os.Exit(1)
	}

	// Get the status
	status, err := w.Status()
	if err != nil {
		fmt.Printf("Error getting status: %v\n", err)
		os.Exit(1)
	}

	// Print the status
	fmt.Println("Repository status:")
	printStatus(status)
	gitCommitAndPush(selectedDirectory, status, w, r)
}

func gitCommitAndPush(path string, status git.Status, w *git.Worktree, r *git.Repository) {
	changedFiles := 0
	for _, stat := range status {
		if stat.Worktree != git.Unmodified {
			changedFiles++
		}
	}

	if changedFiles == 0 {
		fmt.Println("No changes to commit")
		return
	}
	addChanges(w)
	commitMessage := fmt.Sprintf("Commit at %s", time.Now().Format(time.RFC3339))
	_, err := commitChanges(w, commitMessage)
	if err != nil {
		log.Fatal("Error committing changes: ", err)
	}

	err = pushChangesWithShell(path)
	if err != nil {
		log.Fatal("Error pushing changes: ", err)
	}

	fmt.Printf("Pushed %d files\n", changedFiles)
}

func getDirectories(path string) []string {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var dirs []string
	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, file.Name())
		}
	}

	return dirs
}

func getSelectedDirectory(dirpath string) string {
	dirs := getDirectories(dirpath)
	for i, dir := range dirs {
		fmt.Printf("%d: %s\n", i, dir)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the index of the directory you want: ")
	text, _ := reader.ReadString('\n')
	text = strings.ReplaceAll(text, "\r", "") // remove the carriage return character
	text = strings.TrimSpace(text)            // remove the newline character
	index, err := strconv.Atoi(text)
	if err != nil {
		fmt.Printf("Error converting input to integer: %v\n", err)
		os.Exit(1)
	}

	if index < 0 || index >= len(dirs) {
		fmt.Println("Invalid index")
		os.Exit(1)
	}

	return dirpath + "\\" + dirs[index]
}

func printStatus(status git.Status) {
	// Print untracked files
	fmt.Println("Untracked files:")
	for name, stat := range status {
		if stat.Worktree == git.Untracked {
			fmt.Printf("\t%s\n", name)
		}
	}
	// Print changes
	fmt.Println("Changes to be committed:")
	for name, stat := range status {
		if stat.Staging == git.Modified || stat.Staging == git.Added || stat.Staging == git.Deleted {
			fmt.Printf("\t%s - %s\n", name, formatStatus(stat.Staging))
		}
	}
}

func formatStatus(s git.StatusCode) string {
	switch s {
	case git.Modified:
		return "modified"
	case git.Added:
		return "added"
	case git.Deleted:
		return "deleted"
	default:
		return "unknown"
	}
}

func addChanges(w *git.Worktree) {
	err := w.AddGlob("*")
	if err != nil {
		log.Fatal("Error adding changes: ", err)
	}
}

func commitChanges(w *git.Worktree, message string) (plumbing.Hash, error) {
	return w.Commit(message, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "",
			Email: "",
			When:  time.Now(),
		},
	})
}

func pushChanges(r *git.Repository) error {
	auth := &http.BasicAuth{
		Username: "", // this can be anything except an empty string
		Password: "",
	}
	return r.Push(&git.PushOptions{
		Auth: auth,
	})
}

func pushChangesWithShell(path string) error {
	log.Println("Pushing changes", path)
	cmd := exec.Command("git", "push", "origin", "master")
	cmd.Dir = path // Set the working directory for the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error pushing changes: %v - %s", err, output)
	}
	return nil
}
