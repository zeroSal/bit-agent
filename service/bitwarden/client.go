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

func GetNotesItem(sessionKey string, itemName string) string {
	cmd := exec.Command("bw", "get", "notes", itemName)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "BW_SESSION="+sessionKey)

	out, err := cmd.Output()

	if err != nil {
		return err.Error()
	}

	return string(out) + "\n"
}
