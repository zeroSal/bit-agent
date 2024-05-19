package bitwarden

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
)

func Login(email string, password string) (session string, errOut string, success bool) {
	var stderr bytes.Buffer
	var stdout bytes.Buffer

	cmd := exec.Command("bw", "login", email, password, "--raw")
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		return "", stderr.String(), false
	}

	return stdout.String(), "", true
}

func IsAuthenticated() (authenticated bool, errOut string, success bool) {
	var response map[string]interface{}
	var stderr bytes.Buffer
	var stdout bytes.Buffer

	cmd := exec.Command("bw", "status")
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		return false, stderr.String(), false
	}

	err = json.Unmarshal(stdout.Bytes(), &response)
	if err != nil {
		return false, "Failed to parse JSON response.", false
	}

	return response["status"] != "unauthenticated", "", true
}

func Unlock(password string) (session string, errOut string, success bool) {
	var stderr bytes.Buffer
	var stdout bytes.Buffer

	cmd := exec.Command("bw", "unlock", password, "--raw")
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		return "", stderr.String(), false
	}

	return stdout.String(), "", true
}

// Retrieves the notes item content.
// The item parameter can be the item name or its ID.
// A logged in session must be provided.
func GetNotesItem(session string, item string) (content string, errOut string, success bool) {
	cmd := exec.Command("bw", "get", "notes", item)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "BW_SESSION="+session)

	out, err := cmd.Output()

	if err != nil {
		return "", err.Error(), false
	}

	return string(out) + "\n", "", true
}

// Retrieves all the folders.
// A logged in session must be provided.
func ListFolders(session string) (folders []Folder, errOut string, success bool) {
	var response []Folder
	var stderr bytes.Buffer
	var stdout bytes.Buffer

	cmd := exec.Command("bw", "list", "folders")
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "BW_SESSION="+session)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		return nil, stderr.String(), false
	}

	err = json.Unmarshal(stdout.Bytes(), &response)
	if err != nil {
		return nil, "Failed to parse JSON response.", false
	}

	var validFolders []Folder
	for _, folder := range response {
		if len(folder.Id) == 36 {
			validFolders = append(validFolders, folder)
		}
	}

	return validFolders, "", true
}

func ListItemsInFolder(session string, folder Folder) (items []Item, errOut string, success bool) {
	var response []Item
	var stderr bytes.Buffer
	var stdout bytes.Buffer

	cmd := exec.Command("bw", "list", "items", "--folderid", folder.Id)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "BW_SESSION="+session)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		return nil, stderr.String(), false
	}

	err = json.Unmarshal(stdout.Bytes(), &response)
	if err != nil {
		return nil, "Failed to parse JSON response.", false
	}

	return response, "", true
}
