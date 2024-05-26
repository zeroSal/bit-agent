package bitwarden

import (
	"bit-agent/util/cli"
	"os"
	"strconv"
	"strings"
	"time"
)

func Authenticate() (session string) {
	cli.Debug("Checking for authentication...")
	authenticated, errOut, success := isAuthenticated()
	if !success {
		cli.Error("Cannot check for authentication.\n" + errOut)
		os.Exit(1)
	}

	if authenticated {
		return Unlock()
	}

	return Login()
}

func RetrieveKeyByName(session string, name string) (key string) {
	cli.Debug("Retrieving the \"" + name + "\" key...")
	sshKey, errOut, success := getNotesItem(session, name)
	if !success {
		cli.Error("Key retrieving failed.\n" + errOut)
		os.Exit(1)
	}

	if sshKey == "" {
		cli.Warning("The SSH key retrieved seems empty.")
	}

	return sshKey
}

func RetrieveSshFolder(session string) (folder Folder) {
	cli.Debug("Listing all folders...")
	folders, errOut, success := listFolders(session)
	if !success {
		cli.Error("Cannot retrieve folders.\n" + errOut)
		os.Exit(1)
	}

	for _, folder := range folders {
		if folder.Name == "SSH" {
			return folder
		}
	}

	cli.Error("The SSH folder was not found.")
	os.Exit(1)

	return Folder{}
}

func RetrieveSshKeys(session string, folder Folder) (keys []string) {
	cli.Debug("Listing all items in \"" + folder.Name + "\"" + " folder...")
	items, errOut, success := listItemsInFolder(session, folder)
	if !success {
		cli.Error("Cannot retrieve items in folder.\n" + errOut)
		os.Exit(1)
	}

	var keysArray []string
	skipped := 0

	for _, item := range items {
		if strings.Contains(item.Notes, "PRIVATE KEY") {
			keysArray = append(keysArray, item.Notes)
			continue
		}
		skipped++
	}

	if len(keysArray) < 1 {
		cli.Error("No keys loaded.")
		os.Exit(1)
	}

	cli.Notice("Loaded " + strconv.Itoa(len(keysArray)) + " key(s).")

	if skipped > 0 {
		cli.Warning("Skipped " + strconv.Itoa(skipped) + " item(s) as it seems not to be an SSH key.")
	}

	return keysArray
}

func StartSync(session string) {
	cli.Debug("Starting the sync thread...")
	go periodicallySync(session)
}

func periodicallySync(session string) {
	for {
		sync(session)
		time.Sleep(10 * time.Second)
	}
}

func Sync(session string) {
	errOut, success := sync(session)

	if !success {
		cli.Error("Sync failed.\n" + errOut)
		os.Exit(1)
	}
}

func Login() (session string) {
	email := cli.Ask("Bitwarden email: ")
	password := cli.AskHidden("Bitwarden master password: ")

	cli.Debug("Logging into Bitwarden...")
	session, errOut, success := login(email, string(password))
	if !success {
		cli.Error("Login failed.\n" + errOut)
		os.Exit(1)
	}

	return session
}

func Unlock() (session string) {
	password := cli.AskHidden("Bitwarden master password: ")

	cli.Debug("Unlocking the vault...")
	session, errOut, success := unlock(string(password))
	if !success {
		cli.Error("Unlock failed.\n" + errOut)
		os.Exit(1)
	}

	return session
}

func CheckInstallation() {
	cli.Debug("Checking the Bitwarden CLI installation...")
	_, _, success := getVersion()

	if !success {
		cli.Error("The Bitwarden CLI seems not installed or it doesn't work.")
		os.Exit(1)
	}
}

func Version() (version string) {
	cli.Debug("Retrieving the Bitwarden CLI version...")
	v, _, success := getVersion()

	if !success {
		cli.Error("The Bitwarden CLI seems not installed or it doesn't work.")
		os.Exit(1)
	}

	return v
}
